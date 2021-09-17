package nasa

import (
	"testing"
	"time"
)

func TestApod(t *testing.T) {
	client := NewClient()

	res, err := client.Apod()

	if err != nil {
		t.Error(err)
		return
	}

	d := time.Now().Format(layoutISO)

	if res.Date != d {
		t.Errorf("date not matching %v %v", res.Date, d)
	}
}

func TestApodWDate(t *testing.T) {
	client := NewClient()

	params := &ApodOptions{
		Date: "2021-09-14",
	}

	res, err := client.ApodWOpt(params)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Date != "2021-09-14" {
		t.Errorf("incorrect date")
	}
}

func TestApodCount(t *testing.T) {
	client := NewClient()

	res, err := client.ApodCount(2)

	if err != nil {
		t.Error(err)
		return
	}

	if len(*res) != 2 {
		t.Errorf("returns incorrect number of elements")
	}
}
func TestRateLimit(t *testing.T) {
	client := NewClient()

	client.ApodCount(10)

}

func TestInvalidDate(t *testing.T) {
	client := NewClient()

	date := "January 22, 2021"

	options := &ApodOptions{
		Date: date,
	}

	_, err := client.ApodWOpt(options)

	if err == nil {
		t.Errorf("invalid date not checked")
	}
}
