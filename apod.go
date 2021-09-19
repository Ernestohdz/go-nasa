package nasa

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

var apodAPI = &apiConfig{
	host: "https://api.nasa.gov",
	path: "/planetary/apod",
}

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

func (a *ApodOptions) params() url.Values {
	q := make(url.Values)
	if a == nil {
		return q
	}
	if a.Date != "" {
		q.Set("date", a.Date)
	}
	if a.StartDate != "" {
		q.Set("start_date", a.StartDate)
	}
	if a.EndDate != "" {
		q.Set("end_date", a.EndDate)
	}
	if a.Thumbs {
		q.Set("thumbs", "true")
	}
	return q
}

/* Returns Apod for today*/
func (c *Client) Apod() ([]ApodResults, error) {
	return c.ApodWOpt(nil)
}

func (c *Client) ApodWOpt(options *ApodOptions) ([]ApodResults, error) {
	var data ApodResults

	// set it as ApodResults then only set array if its startdate - end date
	// then for regular date create of size 0
	if options == nil {
		err := c.getJSON(apodAPI, options, &data)
		if err != nil {
			return nil, err
		}
		d := []ApodResults{data}
		return d, nil
	}
	if options.Date != "" && (options.StartDate != "" || options.EndDate != "") {
		return nil, errors.New("date options cannot be used with StartDate and EndDate")
	}
	if options.Date != "" {
		if _, err := time.Parse(layoutISO, options.Date); err != nil {
			return nil, err
		}
		err := c.getJSON(apodAPI, options, &data)
		if err != nil {
			return nil, err
		}
		d := []ApodResults{data}
		return d, nil
	}
	if (options.StartDate == "" && options.EndDate != "") || (options.StartDate != "" && options.EndDate == "") {
		return nil, errors.New("StartDate/EndDate option missing EndDate/StartDate option")
	}
	if _, err := time.Parse(layoutISO, options.StartDate); err != nil {
		return nil, errors.New("incorrect StartDate format")
	}
	if _, err := time.Parse(layoutISO, options.EndDate); err != nil {
		return nil, errors.New("incorrect EndDate format")
	}
	var arr []ApodResults
	err := c.getJSON(apodAPI, options, &arr)

	if err != nil {
		return nil, err
	}
	return arr, nil
}

type countOptions struct {
	count  int
	thumbs bool
}

func (c *countOptions) params() url.Values {
	q := make(url.Values)
	q.Set("count", fmt.Sprint(c.count))
	if c.thumbs {
		q.Set("thumbs", "true")
	}
	return q
}

/* Randomly chosen images will be returned. Cannot be used with date or start_date and end_date */
func (c *Client) ApodCount(count int) ([]ApodResults, error) {
	return c.countHelper(count, false)
}

/*
	Randomly chosen images will be returned with thumbnails. Cannot be used with date or start_date and end_date.
	If an APOD is not a video, this parameter is ignored.
*/
func (c *Client) ApodCountWThumbs(count int) ([]ApodResults, error) {
	return c.countHelper(count, true)
}
func (c *Client) countHelper(count int, thumbs bool) ([]ApodResults, error) {
	var arr []ApodResults
	options := &countOptions{
		count:  count,
		thumbs: thumbs,
	}
	err := c.getJSON(apodAPI, options, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
