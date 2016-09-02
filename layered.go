package graph

import "github.com/etnz/stringset"

//Layer partition a graph's vertices into layers.
// such that for every edge from vertex 'u' to vertex 'v', the u's layer has a greater index than v's layer.
//
// If we call L the function that associate a vertex 'u' to the index of the layer it belongs to: L(u) > L(v)
//
func Layer(graph map[string]map[string]struct{}) (layers []map[string]struct{}) {

	// the idea is to find the tails "only" elements and put them into a first layer
	// and to repeat it
	g := Clone(graph) //we clone the graph because we are going to destroy the vertices once assigned a layer
	g = explicit(g)   // force all vertices to exist in the graph (even those define only as end of an edge)

	//compute the tails, (among which are the leafs elements)
	for len(g) > 0 {

		leaves := leaves(g)
		// layer is the current layer ! great
		layers = append(layers, leaves)
		//remove all vertices that are now in the  layer
		for t := range leaves {
			delete(g, t)
		}
		// and also all references
		for n, vertices := range g {

			if inter := stringset.Inter(vertices, leaves); len(inter) > 0 {
				// there are references inside, clean them
				newvertex := stringset.Clone(vertices)
				stringset.Sub(newvertex, inter)
				g[n] = newvertex
			}
		}
		// oki, the base layer has been fully removed
	}
	return
}

//explicit the graph: every vertex at the tail of an edge, is also in the map
func explicit(graph map[string]map[string]struct{}) (g map[string]map[string]struct{}) {
	for v := range Tails(graph) {
		if _, exists := graph[v]; !exists {
			//this is an implicity defined node, make it explicit
			graph[v] = nil
		}
	}
	return graph
}

//leaves return all the leave vertices
//
// the graph must NOT define vertices implicitely i.e node referenced but never defined
func leaves(graph map[string]map[string]struct{}) (leaves map[string]struct{}) {
	leaves = make(map[string]struct{})
	for t, target := range graph {
		if len(target) == 0 { //this is a leaf
			leaves[t] = zzz //append it
		}
	}
	return
}
