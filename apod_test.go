package nasa

import (
	"fmt"
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

	date := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	params := &ApodOptions{
		Date: date,
	}

	res, err := client.ApodWOpt(params)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Date != date.Format(layoutISO) {
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

	fmt.Println(client.rateLimit)
}
