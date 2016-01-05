package graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToposort(t *testing.T) {

	g := map[string]map[string]interface{}{
		"5":  {"11": nil},
		"7":  {"11": nil, "8": nil},
		"3":  {"8": nil, "10": nil},
		"11": {"2": nil, "9": nil, "10": nil},
		"8":  {"9": nil},
	}
	tsort, err := Toposort(g)
	if err != nil {
		t.Errorf("toposort err: %v", err)
	}
	if len(tsort) != 8 || tsort[0] != "3" || tsort[1] != "5" || tsort[2] != "7" || tsort[3] != "11" || tsort[4] != "8" || tsort[5] != "10" || tsort[6] != "2" || tsort[7] != "9" {

		t.Errorf("tsort=%v expecting [3 5 7 11 8 10 2 9]", tsort)
	}

}

func ExampleToposort() {
	//for a given graph
	g := map[string]map[string]interface{}{
		"5":  {"11": nil},
		"7":  {"11": nil, "8": nil},
		"3":  {"8": nil, "10": nil},
		"11": {"2": nil, "9": nil, "10": nil},
		"8":  {"9": nil},
	}
	sorted, _ := Toposort(g)
	//vertices in topological order
	fmt.Println(sorted)
	//Output: [3 5 7 11 8 10 2 9]
}

func TestReverse(t *testing.T) {
	ps := [][]string{
		{"a"},
		{"a", "b"},
		{"a", "b", "c"},
		{"a", "b", "c", "d"},
	}
	xs := [][]string{
		{"a"},
		{"b", "a"},
		{"c", "b", "a"},
		{"d", "c", "b", "a"},
	}
	for i := range ps {
		Reverse(ps[i])
		if !reflect.DeepEqual(ps[i], xs[i]) {
			t.Errorf("Invalid reverse ! OMG ! got %v expecting %v", ps[i], xs[i])
		}
	}

}
