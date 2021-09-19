package nasa

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func mockServer(code int, body string) *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprintln(w, body)
	}))
	return s
}

func TestApod(t *testing.T) {

	mockData := `{
		"copyright": "Damian Peach",
		"date": "2021-09-19",
		"explanation": "On Saturn, the rings tell you the season. On Earth, Wednesday marks an equinox, the time when the  Earth's equator tilts directly toward the Sun. Since Saturn's grand rings orbit along the planet's equator, these rings appear most prominent -- from the direction of the Sun -- when the spin axis of Saturn points toward the Sun. Conversely, when Saturn's spin axis points to the side, an equinox occurs and the edge-on rings are hard to see from not only the Sun -- but Earth. In the featured montage, images of Saturn between the years of 2004 and 2015 have been superposed to show the giant planet passing from southern summer toward northern summer. Saturn was as close as it can get to planet Earth last month, and this month the ringed giant is still bright and visible throughout much of the night",
		"hdurl": "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_2504.jpg",
		"media_type": "image",
		"service_version": "v1",
		"title": "Rings and Seasons of Saturn",
		"url": "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_960.jpg"
		}`
	server := mockServer(200, mockData)
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))

	resp, err := client.Apod()

	if err != nil {
		t.Error(err)
		return
	}

	if resp[0].Date != "2021-09-19" {
		t.Errorf("date not matching %v %v", resp[0].Date, "2021-09-19")
	}
	correctResponse := &ApodResults{
		Copyright:      "Damian Peach",
		Date:           "2021-09-19",
		Explanation:    "On Saturn, the rings tell you the season. On Earth, Wednesday marks an equinox, the time when the  Earth's equator tilts directly toward the Sun. Since Saturn's grand rings orbit along the planet's equator, these rings appear most prominent -- from the direction of the Sun -- when the spin axis of Saturn points toward the Sun. Conversely, when Saturn's spin axis points to the side, an equinox occurs and the edge-on rings are hard to see from not only the Sun -- but Earth. In the featured montage, images of Saturn between the years of 2004 and 2015 have been superposed to show the giant planet passing from southern summer toward northern summer. Saturn was as close as it can get to planet Earth last month, and this month the ringed giant is still bright and visible throughout much of the night",
		Hdurl:          "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_2504.jpg",
		MediaType:      "image",
		ServiceVersion: "v1",
		Title:          "Rings and Seasons of Saturn",
		URL:            "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_960.jpg",
	}
	if !reflect.DeepEqual(*correctResponse, resp[0]) {
		t.Errorf("incorrect response expected: %+v, was \n%+v", *correctResponse, resp[0])
	}
}

func TestApodWDate(t *testing.T) {
	mockData := `{
		"date": "2021-09-14",
		"explanation": "Which way up Mount Sharp? In early September, the robotic rover Curiosity continued its ascent up the central peak of Gale Crater, searching for more clues about ancient water and further evidence that Mars could once have been capable of supporting life.  On this recent Martian morning, before exploratory drilling, the rolling rover took this 360-degree panorama, in part to help Curiosity's human team back on Earth access the landscape and chart possible future routes.  In the horizontally-compressed featured image, an amazing vista across Mars was captured, complete with layered hills, red rocky ground, gray drifting sand, and a dusty atmosphere. The hill just left of center has been dubbed Maria Gordon Notch in honor of a famous Scottish geologist.  The current plan is to direct Curiosity to approach, study, and pass just to the right of Gordon Notch on its exploratory trek.",
		"hdurl": "https://apod.nasa.gov/apod/image/2109/MarsPan360_Curiosity_6144.jpg",
		"media_type": "image",
		"service_version": "v1",
		"title": "Mars Panorama 360 from Curiosity",
		"url": "https://apod.nasa.gov/apod/image/2109/MarsPanCompressed_Curiosity_1080.jpg"
		}
	`

	server := mockServer(200, mockData)
	defer server.Close()
	client := NewClient(WithBaseURL(server.URL))

	params := &ApodOptions{
		Date: "2021-09-14",
	}

	res, err := client.ApodWOpt(params)
	if err != nil {
		t.Error(err)
		return
	}

	if res[0].Date != "2021-09-14" {
		t.Errorf("incorrect date")
	}
	correctResponse := &ApodResults{
		Date:           "2021-09-14",
		Explanation:    "Which way up Mount Sharp? In early September, the robotic rover Curiosity continued its ascent up the central peak of Gale Crater, searching for more clues about ancient water and further evidence that Mars could once have been capable of supporting life.  On this recent Martian morning, before exploratory drilling, the rolling rover took this 360-degree panorama, in part to help Curiosity's human team back on Earth access the landscape and chart possible future routes.  In the horizontally-compressed featured image, an amazing vista across Mars was captured, complete with layered hills, red rocky ground, gray drifting sand, and a dusty atmosphere. The hill just left of center has been dubbed Maria Gordon Notch in honor of a famous Scottish geologist.  The current plan is to direct Curiosity to approach, study, and pass just to the right of Gordon Notch on its exploratory trek.",
		Hdurl:          "https://apod.nasa.gov/apod/image/2109/MarsPan360_Curiosity_6144.jpg",
		MediaType:      "image",
		ServiceVersion: "v1",
		Title:          "Mars Panorama 360 from Curiosity",
		URL:            "https://apod.nasa.gov/apod/image/2109/MarsPanCompressed_Curiosity_1080.jpg",
	}
	if !reflect.DeepEqual(*correctResponse, res[0]) {
		t.Errorf("incorrect response expected: %+v\nreceived: %+v", *correctResponse, res[0])
	}
}

func TestApodCount(t *testing.T) {
	mockData := `[{
		"copyright": "Damian Peach",
		"date": "2021-09-19",
		"explanation": "On Saturn, the rings tell you the season. On Earth, Wednesday marks an equinox, the time when the  Earth's equator tilts directly toward the Sun. Since Saturn's grand rings orbit along the planet's equator, these rings appear most prominent -- from the direction of the Sun -- when the spin axis of Saturn points toward the Sun. Conversely, when Saturn's spin axis points to the side, an equinox occurs and the edge-on rings are hard to see from not only the Sun -- but Earth. In the featured montage, images of Saturn between the years of 2004 and 2015 have been superposed to show the giant planet passing from southern summer toward northern summer. Saturn was as close as it can get to planet Earth last month, and this month the ringed giant is still bright and visible throughout much of the night",
		"hdurl": "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_2504.jpg",
		"media_type": "image",
		"service_version": "v1",
		"title": "Rings and Seasons of Saturn",
		"url": "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_960.jpg"
		},{
			"date": "2021-09-14",
			"explanation": "Which way up Mount Sharp? In early September, the robotic rover Curiosity continued its ascent up the central peak of Gale Crater, searching for more clues about ancient water and further evidence that Mars could once have been capable of supporting life.  On this recent Martian morning, before exploratory drilling, the rolling rover took this 360-degree panorama, in part to help Curiosity's human team back on Earth access the landscape and chart possible future routes.  In the horizontally-compressed featured image, an amazing vista across Mars was captured, complete with layered hills, red rocky ground, gray drifting sand, and a dusty atmosphere. The hill just left of center has been dubbed Maria Gordon Notch in honor of a famous Scottish geologist.  The current plan is to direct Curiosity to approach, study, and pass just to the right of Gordon Notch on its exploratory trek.",
			"hdurl": "https://apod.nasa.gov/apod/image/2109/MarsPan360_Curiosity_6144.jpg",
			"media_type": "image",
			"service_version": "v1",
			"title": "Mars Panorama 360 from Curiosity",
			"url": "https://apod.nasa.gov/apod/image/2109/MarsPanCompressed_Curiosity_1080.jpg"
		}]
	`
	server := mockServer(200, mockData)
	defer server.Close()
	client := NewClient(WithBaseURL(server.URL))

	res, err := client.ApodCount(2)

	if err != nil {
		t.Error(err)
		return
	}

	if len(res) != 2 {
		t.Errorf("returns incorrect number of elements")
		return
	}
	correctResponse := []ApodResults{
		{
			Copyright:      "Damian Peach",
			Date:           "2021-09-19",
			Explanation:    "On Saturn, the rings tell you the season. On Earth, Wednesday marks an equinox, the time when the  Earth's equator tilts directly toward the Sun. Since Saturn's grand rings orbit along the planet's equator, these rings appear most prominent -- from the direction of the Sun -- when the spin axis of Saturn points toward the Sun. Conversely, when Saturn's spin axis points to the side, an equinox occurs and the edge-on rings are hard to see from not only the Sun -- but Earth. In the featured montage, images of Saturn between the years of 2004 and 2015 have been superposed to show the giant planet passing from southern summer toward northern summer. Saturn was as close as it can get to planet Earth last month, and this month the ringed giant is still bright and visible throughout much of the night",
			Hdurl:          "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_2504.jpg",
			MediaType:      "image",
			ServiceVersion: "v1",
			Title:          "Rings and Seasons of Saturn",
			URL:            "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_960.jpg",
		},
		{
			Date:           "2021-09-14",
			Explanation:    "Which way up Mount Sharp? In early September, the robotic rover Curiosity continued its ascent up the central peak of Gale Crater, searching for more clues about ancient water and further evidence that Mars could once have been capable of supporting life.  On this recent Martian morning, before exploratory drilling, the rolling rover took this 360-degree panorama, in part to help Curiosity's human team back on Earth access the landscape and chart possible future routes.  In the horizontally-compressed featured image, an amazing vista across Mars was captured, complete with layered hills, red rocky ground, gray drifting sand, and a dusty atmosphere. The hill just left of center has been dubbed Maria Gordon Notch in honor of a famous Scottish geologist.  The current plan is to direct Curiosity to approach, study, and pass just to the right of Gordon Notch on its exploratory trek.",
			Hdurl:          "https://apod.nasa.gov/apod/image/2109/MarsPan360_Curiosity_6144.jpg",
			MediaType:      "image",
			ServiceVersion: "v1",
			Title:          "Mars Panorama 360 from Curiosity",
			URL:            "https://apod.nasa.gov/apod/image/2109/MarsPanCompressed_Curiosity_1080.jpg",
		},
	}

	if !reflect.DeepEqual(correctResponse, res) {
		t.Errorf("incorrect response expected: %+v\nreceived: %+v\n", correctResponse, res)
	}
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

func TestInvalidOptions(t *testing.T) {
	client := NewClient()

	date := "2021-09-10"
	startDate := "2021-05-1"
	endDate := "2021-06-1"

	options := &ApodOptions{
		Date:      date,
		StartDate: startDate,
		EndDate:   endDate,
	}
	_, err := client.ApodWOpt(options)
	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestDateRange(t *testing.T) {
	mockData := `[
		{
		"copyright": "T. Humbert, S. Barré, A. Desmougin & D. WalliangSociété Lorraine d'AstronomieAstroqueyras",
		"date": "2021-09-17",
		"explanation": "There has been a flash on Jupiter. A few days ago, several groups monitoring our Solar System's largest planet noticed a two-second long burst of light. Such flashes have been seen before, with the most famous being a series of impactor strikes in 1994. Then, fragments of Comet Shoemaker-Levy 9 struck  Jupiter leaving dark patches that lasted for months. Since then, at least seven impacts have been recorded on Jupiter -- usually discovered by amateur astronomers. In the featured video, variations in the Earth's atmosphere cause Jupiter's image to shimmer when, suddenly, a bright flash appears just left of center.  Io and its shadow are visible on the right. What hit Jupiter will likely never be known, but considering what we do know of the nearby Solar System, it was likely a piece of rock and ice -- perhaps the size of a bus -- that broke off long-ago from a passing comet or asteroid.",
		"media_type": "video",
		"service_version": "v1",
		"title": "Video: Flash on Jupiter",
		"url": "https://www.youtube.com/embed/ImVl_TfTFEY?rel=0"
		},
		{
		"date": "2021-09-18",
		"explanation": "In this Hubble Space Telescope image the bright, spiky stars lie in the foreground toward the heroic northern constellation Perseus and well within our own Milky Way galaxy. In sharp focus beyond is UGC 2885, a giant spiral galaxy about 232 million light-years distant. Some 800,000 light-years across compared to the Milky Way's diameter of 100,000 light-years or so, it has around 1 trillion stars. That's about 10 times as many stars as the Milky Way. Part of an investigation to understand how galaxies can grow to such enormous sizes, UGC 2885 was also part of An Interesting Voyage and astronomer Vera Rubin's pioneering study of the rotation of spiral galaxies. Her work was the first to convincingly demonstrate the dominating presence of dark matter in our universe.",
		"hdurl": "https://apod.nasa.gov/apod/image/2109/RubinsGalaxy_hst2000.jpg",
		"media_type": "image",
		"service_version": "v1",
		"title": "Rubin's Galaxy",
		"url": "https://apod.nasa.gov/apod/image/2109/RubinsGalaxy_hst1024.jpg"
		},
		{
		"copyright": "Damian Peach",
		"date": "2021-09-19",
		"explanation": "On Saturn, the rings tell you the season. On Earth, Wednesday marks an equinox, the time when the  Earth's equator tilts directly toward the Sun. Since Saturn's grand rings orbit along the planet's equator, these rings appear most prominent -- from the direction of the Sun -- when the spin axis of Saturn points toward the Sun. Conversely, when Saturn's spin axis points to the side, an equinox occurs and the edge-on rings are hard to see from not only the Sun -- but Earth. In the featured montage, images of Saturn between the years of 2004 and 2015 have been superposed to show the giant planet passing from southern summer toward northern summer. Saturn was as close as it can get to planet Earth last month, and this month the ringed giant is still bright and visible throughout much of the night",
		"hdurl": "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_2504.jpg",
		"media_type": "image",
		"service_version": "v1",
		"title": "Rings and Seasons of Saturn",
		"url": "https://apod.nasa.gov/apod/image/2109/saturn2004to2015_peach_960.jpg"
		}
		]`
	server := mockServer(200, mockData)
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))
	options := &ApodOptions{
		StartDate: "2021-09-17",
		EndDate:   "2021-09-19",
	}

	resp, err := client.ApodWOpt(options)

	if err != nil {
		t.Errorf("expected no error received %s", err)
	}
	if len(resp) != 3 {
		t.Errorf("expected an array of size 3, received %d", len(resp))
	}
}

func TestRequestError(t *testing.T) {
	client := NewClient(WithBaseURL("dsadas"))

	_, err := client.Apod()

	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestApodOptionError(t *testing.T) {
	client := NewClient(WithBaseURL("dsadsa"))

	options := &ApodOptions{
		Date:   "2020-12-01",
		Thumbs: true,
	}
	_, err := client.ApodWOpt(options)

	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestMissingEndDate(t *testing.T) {
	client := NewClient()

	options := &ApodOptions{
		StartDate: "2021-09-01",
	}
	_, err := client.ApodWOpt(options)
	if err == nil {
		t.Errorf("expected error received nil")
	}
}
func TestMissingStartDate(t *testing.T) {
	client := NewClient()

	options := &ApodOptions{
		EndDate: "2021-09-01",
	}
	_, err := client.ApodWOpt(options)
	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestIncorrectStartFormat(t *testing.T) {
	client := NewClient()

	options := &ApodOptions{
		StartDate: "Jan. 26, 2021",
		EndDate:   "2021-09-01",
	}
	_, err := client.ApodWOpt(options)
	if err == nil {
		t.Errorf("expected error received nil")
	}
}
func TestIncorrectEndFormat(t *testing.T) {
	client := NewClient()

	options := &ApodOptions{
		EndDate:   "Jan. 26, 2021",
		StartDate: "2021-09-01",
	}
	_, err := client.ApodWOpt(options)
	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestBadBaseURL(t *testing.T) {
	client := NewClient(WithBaseURL("random"))

	options := &ApodOptions{
		EndDate:   "2021-08-01",
		StartDate: "2021-09-01",
	}
	_, err := client.ApodWOpt(options)
	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestCountWThumbsError(t *testing.T) {
	client := NewClient(WithBaseURL("random"))

	_, err := client.ApodCountWThumbs(2)

	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestNewRequestError(t *testing.T) {
	client := NewClient(WithBaseURL("%"))

	_, err := client.ApodCountWThumbs(2)

	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestRandomServerError(t *testing.T) {
	server := mockServer(500, "Internal Server Error")
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))
	_, err := client.Apod()
	if err == nil {
		t.Errorf("expected error received nil")
	}
}
