package nasa

import (
	"errors"
	"net/url"
	"time"
)

var neowAPI = &apiConfig{
	host: "https://api.nasa.gov",
	path: "/neo/rest/v1/feed",
}

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

type NeoOptions struct {
	StartDate string
	EndDate   string
}

func (n *NeoOptions) params() url.Values {
	q := make(url.Values)
	if n == nil {
		return q
	}
	if n.StartDate != "" {
		q.Set("start_date", n.StartDate)
	}
	if n.EndDate != "" {
		q.Set("end_date", n.EndDate)
	}
	return q
}

func (c *Client) NeoW() (*NeoWResult, error) {
	return c.NeoWOpt(nil)
}

func (c *Client) NeoWOpt(options *NeoOptions) (*NeoWResult, error) {
	var result NeoWResult

	if options == nil || (options.StartDate == "" && options.EndDate == "") {
		err := c.getJSON(neowAPI, options, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	if options.StartDate == "" && options.EndDate != "" {
		return nil, errors.New("no StartDate provided")
	}

	if _, err := time.Parse(layoutISO, options.StartDate); err != nil {
		return nil, err
	}
	if _, err := time.Parse(layoutISO, options.EndDate); err != nil {
		return nil, err
	}

	err := c.getJSON(neowAPI, options, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
