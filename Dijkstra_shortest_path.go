package graph

import (
	"math"

	"github.com/etnz/stringset"
)

//IsoWeight is a weight function that always return 1
//
// dedicated to be used as weight function in DijkstraShortestPathTree as default value
func IsoWeight(vertex1, vertex2 string) float64 { return 1.0 }

// DijkstraShortestPathTree Computes the shortest path tree for a given directed acyclic graph.
//
// The shortest path tree consists of:
//
// - a distance value for each vertex: this value is the shortest distance to 'start'
//
// - 'predecessors' map: from vertex 'u' to 'v' such that in the shortest path from 'start' to 'v', 'u' is the last element :
//
//      'start',... 'u', 'v'
//
// This is the easier way to build a shortest path from 'start'.
//
// if 'weight' function is nil, then the IsoWeight function is used instead.
func DijkstraShortestPathTree(g map[string]map[string]interface{}, start string, weight func(v1, v2 string) float64) (distances map[string]float64, predecessors map[string]string) {

	if weight == nil {
		weight = IsoWeight
	}
	distances = make(map[string]float64)
	predecessors = make(map[string]string)

	//all vertices: will be reduced until all vertices have been scanned
	outside := Vertices(g)

	//first init all stuff
	for vertex := range outside {
		distances[vertex] = math.Inf(1)
	}
	//init the dist from start
	distances[start] = 0

	for len(outside) > 0 { // all vertices have not been processed, yet

		// pick one element in the frontier: *
		// one element in the "outside" that is the closest to start
		var frontierElement string // one of the element in the frontier: i.e in
		minDist := math.Inf(1)
		for v := range outside {
			d := distances[v] // computes its distance
			if d < minDist {  // this one beats all the previous frontierElement, keep it
				minDist = d
				frontierElement = v
			}
		}

		//frontierElement contains  one vertex that is closer possible to start
		// as it will be processed, remove them from verts
		delete(outside, frontierElement)

		// now scan all the neighbors of frontier in verts
		// it's the intersection of next element in the graph, and all elements outside
		// f is an element in the frontier.
		// all neighbours are in the original graph
		neighbours := stringset.Inter(g[frontierElement], outside)

		for neighbour := range neighbours {
			// compute the temp distance from start to neighbour
			d := distances[frontierElement] + weight(frontierElement, neighbour)
			if d < distances[neighbour] { // this new path to n, beats all the previous ones
				distances[neighbour] = d                  // store it as the new champion
				predecessors[neighbour] = frontierElement // and of course, the frontier element that makes it possible
			}
		}
	}
	return

}

// Path builds a path out of a reversed graph, to "ends"
func Path(reversed map[string]string, ends string) []string {
	path := make([]string, 0, len(reversed))
	prev, exists := ends, true
	for exists {
		path = append(path, prev)
		prev, exists = reversed[prev]
	}
	//As the graph is reversed
	Reverse(path)
	return path

}
