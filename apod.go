package nasa

import (
	"fmt"
	"net/http"
	"time"
)

const (
	apodEndpoint = "https://api.nasa.gov/planetary/apod"
	layoutISO    = "2006-01-02"
)

type ApodResults struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

type ApodOptions struct {
	Date      string
	StartDate string
	EndDate   string
	Thumbs    bool
}

func (c *Client) Apod() (*ApodResults, error) {
	return c.ApodWOpt(nil)
}

func (c *Client) ApodWOpt(o *ApodOptions) (*ApodResults, error) {
	req, err := http.NewRequest("GET", apodEndpoint, nil)

	if err != nil {
		return nil, err
	}

	var data ApodResults
	if o == nil {
		err = c.send(req, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}

	if o.Date != "" && (o.StartDate != "" || o.EndDate != "") {
		return nil, fmt.Errorf("date option cannot be used with StartDate or EndDate")
	}
	q := req.URL.Query()

	if o.Thumbs {
		q.Add("thumbs", "true")
	}

	if o.Date != "" {
		p, err := time.Parse(layoutISO, o.Date)
		if err != nil {
			return nil, err
		}
		q.Add("date", p.Format(layoutISO))
		req.URL.RawQuery = q.Encode()
		err = c.send(req, &data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}
	if (o.StartDate == "" && o.EndDate != "") || (o.StartDate != "" && o.EndDate == "") {
		return nil, fmt.Errorf("missing option StartDate or EndDate")
	}

	if _, err := time.Parse(layoutISO, o.StartDate); err != nil {
		return nil, err
	}
	if _, err := time.Parse(layoutISO, o.EndDate); err != nil {
		return nil, err
	}
	q.Add("start_date", o.StartDate)
	q.Add("end_date", o.EndDate)

	req.URL.RawQuery = q.Encode()
	err = c.send(req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) ApodCount(count int) (*[]ApodResults, error) {
	return c.countHelper(count, false)
}
func (c *Client) ApodCountWThumbs(count int) (*[]ApodResults, error) {
	return c.countHelper(count, true)
}
func (c *Client) countHelper(count int, thumbs bool) (*[]ApodResults, error) {
	var arr []ApodResults
	req, _ := http.NewRequest("GET", apodEndpoint, nil)
	q := req.URL.Query()
	if thumbs {
		q.Add("thumbs", "true")
	}
	q.Add("count", fmt.Sprint(count))
	req.URL.RawQuery = q.Encode()
	err := c.send(req, &arr)
	if err != nil {
		return nil, err
	}
	return &arr, nil
}
