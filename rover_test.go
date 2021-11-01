package nasa

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCuriosityRoverEmptyQuery(t *testing.T) {
	mockData := `{
		"photos": []
		}`
	server := mockServer(200, mockData)
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))

	resp, err := client.CuriosityRover()

	if err != nil {
		t.Error(err)
		return
	}
	correctResponse := RoverResult{}
	err = json.Unmarshal([]byte(mockData), &correctResponse)
	if err != nil {
		t.Errorf("error when parsing mock data: %s\n", err)
		return
	}
	if !reflect.DeepEqual(correctResponse, *resp) {
		t.Errorf("incorrect response from client: %+v\n", *resp)
	}
}

func TestOpportunityManifest(t *testing.T) {
	mockData := `{
		"photo_manifest": {
			"name": "Opportunity",
			"landing_date": "2004-01-25",
			"launch_date": "2003-07-07",
			"status": "complete",
			"max_sol": 5111,
			"max_date": "2018-06-11",
			"total_photos": 198439,
			"photos": [{
				"sol": 1,
				"earth_date": "2004-01-26",
				"total_photos": 95,
				"cameras": [
				"ENTRY",
				"FHAZ",
				"NAVCAM",
				"PANCAM",
				"RHAZ"
				]
			}]
		}
	}
	`
	server := mockServer(200, mockData)
	defer server.Close()

	client := NewClient(WithBaseURL(server.URL))

	resp, err := client.OpportunityManifest()
	if err != nil {
		t.Error(err)
		return
	}
	correctResponse := ManifestResult{}
	err = json.Unmarshal([]byte(mockData), &correctResponse)

	if err != nil {
		t.Errorf("test unmarshal error")
		return
	}
	if !reflect.DeepEqual(correctResponse, *resp) {
		t.Errorf("incorrect response from client: %+v\n", *resp)
	}
}
