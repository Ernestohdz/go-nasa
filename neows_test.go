package nasa

import (
	"fmt"
	"testing"
)

func TestNeo(t *testing.T) {
	c := NewClient()

	resp, err := c.NeoW()

	if err != nil {
		t.Errorf("%s", err)
		return
	}

	fmt.Printf("%+v\n", *resp)
}
