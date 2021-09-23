package nasa

import (
	"errors"
	"net/url"
	"time"
)

var neoAPI = &apiConfig{
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

// Near Earth Objects(NEO) struct representation for NeoWResults
type NEO struct {
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

// NeoFeedResult is the struct representation of NASA's Asteroid NeoWs response
type NeoFeedResult struct {
	// structure contains the links of prev, next, and self feeds.
	Links struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
		Self string `json:"self"`
	} `json:"links"`
	// total number of near earth objects
	ElementCount int `json:"element_count"`
	// Contains dates mapping to an array of Asteroid objects
	NearEarthObjects map[string][]NEO `json:"near_earth_objects"`
}

// NeoFeedOptions is the optins struct for NeoW
type NeoFeedOptions struct {
	// search from specific start date
	StartDate string
	// end search at specific end date
	EndDate string
}

func (n *NeoFeedOptions) params() url.Values {
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

// NeoFeed function returns a 7 day feed, starting from current date, of near earth objects
func (c *Client) NeoFeed() (*NeoFeedResult, error) {
	return c.NeoFeedWOpt(nil)
}

// NeoFeedWOpt returns a custom feed with given options for a range of dates of near
// earth objects
func (c *Client) NeoFeedWOpt(options *NeoFeedOptions) (*NeoFeedResult, error) {
	var result NeoFeedResult

	if options == nil || (options.StartDate == "" && options.EndDate == "") {
		err := c.getJSON(neoAPI, options, &result)
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

	err := c.getJSON(neoAPI, options, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
