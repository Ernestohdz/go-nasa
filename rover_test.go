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
