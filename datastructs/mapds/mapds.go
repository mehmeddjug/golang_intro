package mapds

import "fmt"

// Function with starting capital letter is public function/method
func Decleration() {
	// Create a map with a key/value of type string/int.
	mapDict := make(map[string]int)
	fmt.Println(mapDict)

	// Create and initialize the map with 2 key/value of type string pairs.
	mapDictTwo := map[string]string{"One": "One-Two", "Three": "Three-Four"}
	fmt.Println(mapDictTwo)

	// Create a map using a slice of strings as the value.
	mapDictThree := map[int][]string{}
	fmt.Println(mapDictThree)

	// Create an empty map.
	emptyMap := map[string]string{}

	// Add Some numbers to emptyMap.
	emptyMap["Some"] = "numbers"
	fmt.Println(emptyMap)
}
