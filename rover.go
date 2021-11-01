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

var manifestAPI = &apiConfig{
	host: "https://api.nasa.gov",
	path: "/mars-photos/api/v1/manifests",
}

type ManifestResult struct {
	PhotoManifest struct {
		Name        string `json:"name"`
		LandingDate string `json:"landing_date"`
		LaunchDate  string `json:"launch_date"`
		Status      string `json:"status"`
		MaxSol      int    `json:"max_sol"`
		MaxDate     string `json:"max_date"`
		TotalPhotos int    `json:"total_photos"`
		Photos      []struct {
			Sol         int      `json:"sol"`
			EarthDate   string   `json:"earth_date"`
			TotalPhotos int      `json:"total_photos"`
			Cameras     []string `json:"cameras"`
		} `json:"photos"`
	} `json:"photo_manifest"`
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
	return c.roverHelper(spiritAPI, nil)
}
func (c *Client) SpiritRoverWOpt(options *RoverOptions) (*RoverResult, error) {
	return c.roverHelper(spiritAPI, options)
}
func (c *Client) SpiritManifest() (*ManifestResult, error) {
	manifest := manifestAPI
	manifest.path = manifest.path + "/spirit"
	return c.manifestHelper(manifest)
}

func (c *Client) CuriosityRover() (*RoverResult, error) {
	return c.roverHelper(curiosityAPI, nil)
}
func (c *Client) CuriosityRoverWOpt(options *RoverOptions) (*RoverResult, error) {
	return c.roverHelper(curiosityAPI, options)
}
func (c *Client) CuriosuityManifest() (*ManifestResult, error) {
	manifest := manifestAPI
	manifest.path = manifest.path + "/curiosity"
	return c.manifestHelper(manifest)
}

func (c *Client) OpportunityRover() (*RoverResult, error) {
	return c.roverHelper(opportunityAPI, nil)
}
func (c *Client) OpportunityRoverWOpt(options *RoverOptions) (*RoverResult, error) {
	return c.roverHelper(opportunityAPI, options)
}
func (c *Client) OpportunityManifest() (*ManifestResult, error) {
	manifest := manifestAPI
	(*manifest).path = (*manifest).path + "/opportunity"
	return c.manifestHelper(manifest)
}

func (c *Client) manifestHelper(path *apiConfig) (*ManifestResult, error) {
	var result ManifestResult
	err := c.getJSON(path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) roverHelper(rover *apiConfig, options *RoverOptions) (*RoverResult, error) {
	var result RoverResult
	if options == nil {
		err := c.getJSON(rover, options, &result)
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
	err := c.getJSON(rover, options, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
