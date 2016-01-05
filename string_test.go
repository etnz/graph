package graph

import "fmt"

func ExampleString() {
	g := map[string]map[string]interface{}{
		"5":  {"11": nil},
		"7":  {"11": nil, "8": nil},
		"3":  {"8": nil, "10": nil},
		"11": {"2": nil, "9": nil, "10": nil},
		"8":  {"9": nil},
	}
	fmt.Println(String(g))
	//Output:
	// "10";
	// "11" -> "10";
	// "11" -> "2";
	// "11" -> "9";
	// "2";
	// "3" -> "10";
	// "3" -> "8";
	// "5" -> "11";
	// "7" -> "11";
	// "7" -> "8";
	// "8" -> "9";
	// "9";
}
