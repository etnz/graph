package graph

import "github.com/etnz/stringset"

// expanding is the action to convert a layered graph so that each edge goes from one layer to the next one
// edges shall not cross layers.

// Expand a graph by adding vertex so that if we call L the function that associate a vertex 'u' to the index of the layer it belongs to: L(u) =1+ L(v)
//
// If an edge (uv) "crosses" layers L(u)>1+L(v) then the function 'vertexgen' is called to generate a new vertex 'w' so that
// and (uv) is replaced by (uw) and (w,v), by construction now, L(u)=1+L(w).
func Expand(graph map[string]map[string]struct{}, vertexgen func() string) (expanded map[string]map[string]struct{}) {
	//alg:
	// compute the current layers
	// then for each vertex in the top layer, create if needed a vertex in the layer below.

	expanded = Clone(graph)
	layers := Layer(expanded) //compute the layers

	for i := len(layers) - 1; i >= 0; i-- {
		// i starting from  top layer down
		for vertex := range layers[i] {
			// let's test if we need to augment this vertex or not
			for tail := range expanded[vertex] {
				// the deal here is that we have an edge from vertex -> tail
				// vertex belongs to layer i, if tail does not belong to i-1, we need to create one
				if !stringset.Contains(layers[i-1], tail) {
					//here we are  tail is not in the right layer.
					// just create a new vertex, and add it
					ivertex := vertexgen()
					// operation consist in removing tail from the tails of "vertex"
					// and replacing it by ivertex
					// then creating anew edge from ivertex to tail
					delete(expanded[vertex], tail)  //remove tail from tails of vertex
					expanded[vertex][ivertex] = zzz // replace by an edge to ivertex
					//and now create the new ivertex, and its edge
					expanded[ivertex] = stringset.New(tail)
					// as we want to check ivertex in the next layer, we need to append it:
					layers[i-1][ivertex] = zzz // also append inode, the lower layer

				}
			}
		}
	}
	return

}
