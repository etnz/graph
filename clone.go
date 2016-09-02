package graph

import "github.com/etnz/stringset"

// Clone a simple graph
func Clone(graph map[string]map[string]struct{}) (clone map[string]map[string]struct{}) {
	clone = make(map[string]map[string]struct{})
	for k, v := range graph {
		clone[k] = stringset.Clone(v)
	}
	return
}
