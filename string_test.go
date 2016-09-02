package graph

import "fmt"

func ExampleString() {
	g := map[string]map[string]struct{}{
		"5":  {"11": zzz},
		"7":  {"11": zzz, "8": zzz},
		"3":  {"8": zzz, "10": zzz},
		"11": {"2": zzz, "9": zzz, "10": zzz},
		"8":  {"9": zzz},
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
