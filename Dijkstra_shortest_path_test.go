package graph

import (
	"fmt"
	"testing"
)

func TestDijkstraShortestPathTree(t *testing.T) {
	// according to wikipedia:
	graph := map[string]map[string]interface{}{
		"1": {"2": 7.0, "3": 9.0, "6": 14.0},
		"2": {"3": 10.0, "4": 15.0},
		"3": {"4": 11.0, "6": 2.0},
		"4": {"5": 6.0},
		"5": {},
		"6": {"5": 9.0},
	}
	// trick we are using the graph edge value to store a float: the wieght
	dists, preds := DijkstraShortestPathTree(graph, "1", func(a, b string) float64 { return graph[a][b].(float64) })

	xdists := map[string]float64{
		"1": 0.0,
		"2": 7.0,
		"3": 9.0,
		"4": 20.0,
		"5": 20.0,
		"6": 11.0,
	}

	xpreds := map[string]string{
		"1": "",
		"2": "1",
		"3": "1",
		"4": "3",
		"5": "6",
		"6": "3",
	}

	for k := range xdists {
		if dists[k] != dists[k] {
			t.Errorf("invalid distance, got %v expecting %v", dists[k], xdists[k])
		}
	}
	for k := range xpreds {
		if xpreds[k] != preds[k] {
			t.Errorf("invalid preds, got %v expecting %v", preds[k], xpreds[k])
		}
	}

}

func ExampleDijkstraShortestPathTree() {
	// according to wikipedia: https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	graph := map[string]map[string]interface{}{
		"1": {"2": 7.0, "3": 9.0, "6": 14.0},
		"2": {"3": 10.0, "4": 15.0},
		"3": {"4": 11.0, "6": 2.0},
		"4": {"5": 6.0},
		"5": {},
		"6": {"5": 9.0},
	}

	// trick we are using the graph edge value to store a float: the weight
	weight := func(a, b string) float64 { return graph[a][b].(float64) }

	_, preds := DijkstraShortestPathTree(graph, "1", weight)
	path := Path(preds, "5")
	fmt.Println(path)
	// Output: [1 3 6 5]

}

func ExamplePath() {
	//a reversed map (gradient map)
	reversed := map[string]string{
		"1": "",
		"2": "1",
		"3": "1",
		"4": "3",
		"5": "6",
		"6": "3",
	}
	p := Path(reversed, "1")
	fmt.Println(p)
	//Output :  [1 3 6 5]

}
