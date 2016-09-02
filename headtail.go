package graph

import "github.com/etnz/stringset"

//Tails  return all vertices that are at the tail of an edge
func Tails(graph map[string]map[string]struct{}) (ends map[string]struct{}) {
	ends = make(map[string]struct{})
	for _, targets := range graph {
		stringset.Append(ends, targets)
	}
	return
}

//Heads return all vertices that are head of an edge
func Heads(graph map[string]map[string]struct{}) (heads map[string]struct{}) {
	heads = make(map[string]struct{})
	for h, targets := range graph {
		if len(targets) != 0 {
			heads[h] = zzz //add it in that case
		}
	}
	return
}
