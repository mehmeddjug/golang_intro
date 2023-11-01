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

	// Declering a nil map.
	var emptyMapNil map[string]string
	fmt.Println(emptyMapNil)

	// Create an empty map.
	emptyMap := map[string]string{}

	// Add Some numbers to emptyMap.
	emptyMap["Some"] = "numbers"
	fmt.Println(emptyMap)

	// Read value of the key "Some".
	value1, exists := emptyMap["Some"]

	// Check if the key exists?
	if exists {
		fmt.Println(value1)
	}

	// Read value of the key "Some".
	value2 := emptyMap["Some"]

	// Check if the key exists?
	if value2 != "" {
		fmt.Println(value2)
	}

	// Delete the key/value pair from the map.
	delete(emptyMap, "Some")
	fmt.Println(emptyMap)
}

func Map(mapDict map[string]string, key string, value string) {
	mapDict[key] = value
	fmt.Println(mapDict)
}
