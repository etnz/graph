package graph

import (
	"fmt"
	"testing"

	"github.com/etnz/stringset"
)

func TestLayer(t *testing.T) {

	g := map[string]map[string]interface{}{
		"5":  {"11": nil},
		"7":  {"11": nil, "8": nil},
		"3":  {"8": nil, "10": nil},
		"11": {"2": nil, "9": nil, "10": nil},
		"8":  {"9": nil},
	}
	layers := Layer(g)
	x := []map[string]interface{}{
		stringset.New("2", "9", "10"),
		stringset.New("11", "8"),
		stringset.New("5", "7", "3"),
	}

	if len(layers) != len(x) {
		t.Errorf("invalid layer length:got %v expected %v", len(layers), len(x))
	}
	for i := 0; i < 3; i++ {

		if !stringset.Equals(x[i], layers[i]) {
			t.Errorf("invalid layer[%v]:got %v expected %v", i, layers[i], x[i])
		}
	}

}

func ExampleLayer() {
	// from a given graph
	g := map[string]map[string]interface{}{
		"5":  {"11": nil},
		"7":  {"11": nil, "8": nil},
		"3":  {"8": nil, "10": nil},
		"11": {"2": nil, "9": nil, "10": nil},
		"8":  {"9": nil},
	}
	layers := Layer(g)
	for i, l := range layers {
		fmt.Printf("%d: %v\n", i, stringset.Sort(l)) //layer is sorted for repeatibility
	}
	//Output:
	// 0: [10 2 9]
	// 1: [11 8]
	// 2: [3 5 7]
}
