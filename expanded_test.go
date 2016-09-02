package graph

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	g := map[string]map[string]struct{}{
		"5":  {"11": zzz},
		"7":  {"11": zzz, "8": zzz},
		"3":  {"8": zzz, "10": zzz},
		"11": {"2": zzz, "9": zzz, "10": zzz},
		"8":  {"9": zzz},
	}
	// once layered and expanded, it has to create a node between 3 and 10
	i := 0
	expanded := Expand(g, func() string {
		i++
		return fmt.Sprintf("i%v", i)
	})

	x := map[string]map[string]struct{}{
		"5":  {"11": zzz},
		"7":  {"11": zzz, "8": zzz},
		"3":  {"8": zzz, "i1": zzz},
		"11": {"2": zzz, "9": zzz, "10": zzz},
		"8":  {"9": zzz},
		"i1": {"10": zzz},
	}

	if !Equals(expanded, x) {
		t.Errorf("Invalid graph expansion: got \n%v, expecting \n%v", String(expanded), String(x))

	}

}
