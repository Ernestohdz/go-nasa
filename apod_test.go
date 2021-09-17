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
	}

	d := time.Now().Format(layoutISO)

	if res.Date != d {
		t.Errorf("date not matching")
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
	}

	if res.Date != date.Format(layoutISO) {
		t.Errorf("incorrect date")
	}
}
