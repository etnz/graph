package graph

import "testing"

func TestClone(t *testing.T) {

	g1 := map[string]map[string]interface{}{
		"a": {"b": nil, "c": nil},
		"b": {"a": nil, "c": nil},
	}

	g2 := Clone(g1)

	if !Equals(g1, g2) {
		t.Errorf("g2 is not the clone of g1: \n%v\n===\\===\n%v\n", String(g2), String(g1))
	}

}
