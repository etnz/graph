package graph

import "github.com/etnz/stringset"

// Clone a simple graph
func Clone(graph map[string]map[string]interface{}) (clone map[string]map[string]interface{}) {
	clone = make(map[string]map[string]interface{})
	for k, v := range graph {
		clone[k] = stringset.Clone(v)
	}
	return
}
