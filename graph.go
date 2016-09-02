// Package graph implements a simple graph definitions and some algorithms
//
// A graph is usually defined through it's simplest basic type:
//
//    map[string]map[string]struct{}
//
// Which is basically a map from a vertex to a set of vertices,representing edges.
//
// each vertex is represented by a 'string' usually its id.
package graph

import "github.com/etnz/stringset"

var zzz = struct{}{}

/*
# Design

The first goal here is to build a [Layered graph](https://en.wikipedia.org/wiki/Layered_graph_drawing),
and to build all the required "bits" as independent algorithm.
so far we built:
- stringset in its own repo (a lot of operations on nodesets)
- permute in its own repo (to scan all possible crossings)






*/

//Vertices return all the vertices of this graph
func Vertices(g map[string]map[string]struct{}) (vertices map[string]struct{}) {
	vertices = make(map[string]struct{})
	for k, targets := range g {
		vertices[k] = zzz
		stringset.Append(vertices, targets)
	}
	return
}
