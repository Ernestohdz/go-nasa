package nasa

import (
	"encoding/json"
	"net/http"
)

type diameter struct {
	Min float64 `json:"estimated_diameter_min"`
	Max float64 `json:"estimated_diameter_max"`
}

type closeApproachData struct {
	CloseApproachDate      string `json:"close_approach_date"`
	CloseApproachDateFull  string `json:"close_approach_date_full"`
	EpochDateCloseApproach int64  `json:"epoch_date_close_approach"`
	RelativeVelocity       struct {
		KilometersPerSecond string `json:"kilometers_per_second"`
		KilometersPerHour   string `json:"kilometers_per_hour"`
		MilesPerHour        string `json:"miles_per_hour"`
	} `json:"relative_velocity"`
	MissDistance struct {
		Astronomical string `json:"astronomical"`
		Lunar        string `json:"lunar"`
		Kilometers   string `json:"kilometers"`
		Miles        string `json:"miles"`
	} `json:"miss_distance"`
	OrbitingBody string `json:"orbiting_body"`
}

type Asteroid struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ID                 string  `json:"id"`
	NeoReferenceID     string  `json:"neo_reference_id"`
	Name               string  `json:"name"`
	NasaJplURL         string  `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH float64 `json:"absolute_magnitude_h"`
	EstimatedDiameter  struct {
		Kilometers diameter `json:"kilometers"`
		Meters     diameter `json:"meters"`
		Miles      diameter `json:"miles"`
		Feet       diameter `json:"feet"`
	} `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool                `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData              []closeApproachData `json:"close_approach_data"`
	IsSentryObject                 bool                `json:"is_sentry_object"`
}
type NeoWResult struct {
	Links struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
		Self string `json:"self"`
	} `json:"links"`
	ElementCount     int                   `json:"element_count"`
	NearEarthObjects map[string][]Asteroid `json:"near_earth_objects"`
}

func (c *Client) NeoW() (*NeoWResult, error) {

	url := "https://api.nasa.gov/neo/rest/v1/feed?start_date=2021-09-18&end_date=2021-09-19&api_key=DEMO_KEY"

	request, _ := http.NewRequest("GET", url, nil)

	resp, err := c.httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	var body NeoWResult
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}
