package nasa

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	fmt.Println("Hello World")
	c := NewClient("DEMO_KEY")
	fmt.Println(c.BaseURL)
}
