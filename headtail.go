package graph

import "github.com/etnz/stringset"

//Tails  return all vertices that are at the tail of an edge
func Tails(graph map[string]map[string]interface{}) (ends map[string]interface{}) {
	ends = make(map[string]interface{})
	for _, targets := range graph {
		stringset.Append(ends, targets)
	}
	return
}

//Heads return all vertices that are head of an edge
func Heads(graph map[string]map[string]interface{}) (heads map[string]interface{}) {
	heads = make(map[string]interface{})
	for h, targets := range graph {
		if len(targets) != 0 {
			heads[h] = nil //add it in that case
		}
	}
	return
}
