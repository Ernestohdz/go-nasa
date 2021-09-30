package nasa

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// /rovers/%s/photos
// %s - different types of rovers:
// curiosity, opportunity, spirit
var curiosityAPI = &apiConfig{
	host: "https://api.nasa.gov",
	path: "/mars-photos/api/v1/rovers/curiosity/photos",
}
var opportunityAPI = &apiConfig{
	host: "https://api.nasa.gov",
	path: "/mars-photos/api/v1/rovers/opportunity/photos",
}
var spiritAPI = &apiConfig{
	host: "https://api.nasa.gov",
	path: "/mars-photos/api/v1/rovers/spirit/photos",
}

type RoverResult struct {
	Photos []struct {
		ID     int `json:"id"`
		Sol    int `json:"sol"`
		Camera struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			RoverID  int    `json:"rover_id"`
			FullName string `json:"full_name"`
		} `json:"camera"`
		ImgSrc    string `json:"img_src"`
		EarthDate string `json:"earth_date"`
		Rover     struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			LandingDate string `json:"landing_date"`
			LaunchDate  string `json:"launch_date"`
			Status      string `json:"status"`
		} `json:"rover"`
	} `json:"photos"`
}

type RoverOptions struct {
	Sol       int `default:"-1"`
	EarthDate string
	Camera    string
	Page      int
}

func (r *RoverOptions) params() url.Values {
	q := make(url.Values)
	if r == nil {
		return q
	}
	if r.Sol != -1 {
		q.Set("sol", fmt.Sprint(r.Sol))
	}
	if r.EarthDate != "" {
		q.Set("earth_date", r.EarthDate)
	}
	if r.Camera != "" {
		q.Set("camera", r.Camera)
	}
	if r.Page != 0 {
		q.Set("page", fmt.Sprint(r.Page))
	}
	return q
}

func (c *Client) SpiritRover() (*RoverResult, error) {
	return c.helper(spiritAPI, nil)
}
func (c *Client) CuriosityRover() (*RoverResult, error) {
	return c.helper(curiosityAPI, nil)
}
func (c *Client) OpportunityRover() (*RoverResult, error) {
	return c.helper(opportunityAPI, nil)
}

func (c *Client) helper(rover *apiConfig, options *RoverOptions) (*RoverResult, error) {
	var result RoverResult
	if options == nil {
		err := c.getJSON(rover, options, result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	if options.Sol != -1 && options.EarthDate != "" {
		return nil, errors.New("cannot provide Sol and EarthDate option")
	}
	if _, err := time.Parse(layoutISO, options.EarthDate); options.EarthDate != "" && err != nil {
		return nil, err
	}
	err := c.getJSON(rover, options, result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
