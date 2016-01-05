package graph

import "github.com/etnz/stringset"

//CyclicGraphError is an error return while attempting to toposort a cyclic graph.
//
// it's an error, but it is also a graph, of the remaining "headless" subgraph.
//
type CyclicGraphError map[string]map[string]interface{}

func (e CyclicGraphError) Error() string { return String(e) }

//Toposort computes the Toposort for a graph
//
// A toposort of a graph is a linear ordering of its vertices such that for every edge from vertex 'u' to vertex 'v', 'u' comes before 'v' in the ordering.
func Toposort(graph map[string]map[string]interface{}) (tsort []string, err error) {

	// get all vertices from the graph
	g := Clone(graph) // clone the graph to destroy it

	// if g defines vertices implicitely (they only appear as tails) then they are beeing discarded when the edge pointing to them is discarded.
	// It's not a big deal, because they are not involved in the sort, but, hey, they still part of the solution.
	// so we normalize the graph
	g = explicit(g)

	vertices := Vertices(g) // g all vertices

	//vertices : the set of vertices with no incoming edges
	// to get started its a clone of all vertices
	// then we visit every node in the graph, and remove them if they are on the right hand side
	length := len(vertices) // the length of all vertices before removing tails

	stringset.Sub(vertices, Tails(g))

	// the goal is just to sort vertices
	tsort = make([]string, 0, length)

	// s is ready
	for len(vertices) > 0 { // there are still some vertices to deal with
		ns := stringset.Sort(vertices)
		tsort = append(tsort, ns...)
		// and remove all outgoing edges
		for _, n := range ns {
			delete(g, n) // remove n from graph (including all outgoing edges right)
		}
		// all vertices have been processed, rebuild s
		vertices = Vertices(g) // all vertices, again (this time without the heads)
		stringset.Sub(vertices, Tails(g))
	}

	// done: no more "vertices" without incoming edges (i.e heads), a headless graph must be empty, unless it has cycles
	if len(g) > 0 {
		// the graph wasn't acyclic
		return tsort, CyclicGraphError(g)
	}
	return tsort, nil

}

// Reverse a given slice of string
func Reverse(sort []string) {
	ns := len(sort) - 1
	for i := 0; i < (ns+1)/2; i++ {
		sort[i], sort[ns-i] = sort[ns-i], sort[i]
	}
}
