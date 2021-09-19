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

// ApodResults is the struct representation of NASA's APOD response
type ApodResults struct {
	// Copyright is returned if the image is not public domain
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

// ApodOptions is the funtional options struct for Apod
type ApodOptions struct {
	// The date of the APOD image to retrieve
	Date string
	// The start of a date range, when requesting date for a range of dates.
	// Cannot be used with date.
	StartDate string
	// The end of the date range, when used with start_date.
	EndDate string
	// Return the URL of video thumbnail. If an APOD is not a video, this parameter is ignored.
	Thumbs bool
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

// Apod sends an Apod request and retrieves response
func (c *Client) Apod() ([]ApodResults, error) {
	return c.ApodWOpt(nil)
}

// ApodWOpt sends an Apod request with options provided and retrieves response
func (c *Client) ApodWOpt(options *ApodOptions) ([]ApodResults, error) {
	var data ApodResults

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

// Randomly chosen images will be returned. Cannot be used with date or start_date and end_date
func (c *Client) ApodCount(count int) ([]ApodResults, error) {
	return c.countHelper(count, false)
}

// Randomly chosen images will be returned with thumbnails. Cannot be used with date or start_date and end_date.
// If an APOD is not a video, this parameter is ignored.
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
