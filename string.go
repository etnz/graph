package graph

import (
	"bytes"
	"fmt"

	"github.com/etnz/stringset"
)

// String returns a dot-formatter representation of this graph. It does not print the graph definition (`graph name {}`) but just the content
func String(g map[string]map[string]struct{}) string {

	var buf bytes.Buffer

	//sort keys alphabetically
	keys := stringset.Sort(Vertices(g))
	for _, k := range keys {
		if len(g[k]) > 0 {
			targets := stringset.Sort(g[k])
			for _, v := range targets {
				fmt.Fprintf(&buf, "%q -> %q;\n", k, v)
			}
		} else {
			fmt.Fprintf(&buf, "%q;\n", k)
		}
	}
	return buf.String()

}
