package graph

//Equals returns true iif g1 and g2 are identical, (same edges, and same vertices)
//
// It does not recognize graphs that could be equivalent.
func Equals(g1, g2 map[string]map[string]interface{}) bool {

	//build the union of all keys, to be able to compare them
	all := make(map[string]interface{})
	for k := range g1 {
		all[k] = nil
	}
	for k := range g2 {
		all[k] = nil
	}
	//all now contains all keys, use them to compare
	for k := range all {
		v1, v2 := g1[k], g2[k]
		if len(v1) != len(v2) {
			return false
		}
		//ok smae length
		for i := range v1 {
			if v1[i] != v2[i] {
				return false
			}
		}
	}
	//everything was ok !
	return true

}
