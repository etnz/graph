package graph

import "testing"

func TestEquals(t *testing.T) {

	g1 := map[string]map[string]struct{}{
		"a": {"b": zzz, "c": zzz},
		"d": {"b": zzz},
		"c": {},
	}
	g2 := map[string]map[string]struct{}{
		"a": {"b": zzz, "c": zzz},
		"d": {"b": zzz},
	}

	if !Equals(g1, g2) {
		t.Errorf("g1 should be equal to g2: \n%v\n%v\n", String(g1), String(g2))
	}

}
