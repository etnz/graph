package graph

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	g := map[string]map[string]interface{}{
		"5":  {"11": nil},
		"7":  {"11": nil, "8": nil},
		"3":  {"8": nil, "10": nil},
		"11": {"2": nil, "9": nil, "10": nil},
		"8":  {"9": nil},
	}
	// once layered and expanded, it has to create a node between 3 and 10
	i := 0
	expanded := Expand(g, func() string {
		i++
		return fmt.Sprintf("i%v", i)
	})

	x := map[string]map[string]interface{}{
		"5":  {"11": nil},
		"7":  {"11": nil, "8": nil},
		"3":  {"8": nil, "i1": nil},
		"11": {"2": nil, "9": nil, "10": nil},
		"8":  {"9": nil},
		"i1": {"10": nil},
	}

	if !Equals(expanded, x) {
		t.Errorf("Invalid graph expansion: got \n%v, expecting \n%v", String(expanded), String(x))

	}

}
