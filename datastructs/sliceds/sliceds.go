package sliceds

import "fmt"

// Function with starting capital letter is public function/method
func Decleration() {
	// Create a slice of strings with length and capacity of 5 elements.
	sliceOne := make([]string, 5)
	fmt.Println(sliceOne)

	// Create a slice of integers with length of 3 and has a capacity of 5 elements.
	sliceTwo := make([]int, 3, 5)
	fmt.Println(sliceTwo)

	// Create a slice with slice literals.
	sliceThree := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	sliceFour := []int{1, 2, 3}
	fmt.Println(sliceThree, sliceFour)

	// Initialize the 5th element with an empty string.
	sliceFive := []string{4: ""}
	fmt.Println(sliceFive)

	// Create an array of three integers.
	arrayDS := [3]int{1, 2, 3}
	// Create a slice of integers with a length and capacity of three.
	sliceDS := []int{1, 2, 3}
	fmt.Println(arrayDS, sliceDS)

	// Create a nil slice of integers.
	var sliceNil []int
	fmt.Println(sliceNil)
	// Use make to create an empty slice of integers.
	sliceEmpty := make([]int, 0)
	// Use a slice literal to create an empty slice of integers.
	sliceEmptyTwo := []int{}
	fmt.Println(sliceEmpty, sliceEmptyTwo)
}

func Slice(slice1 [3]string) {
	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice := []int{1, 2, 3, 4, 5}

	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	newSlice := slice[1:3]
	fmt.Println(slice)
	fmt.Println(newSlice)
	// For slice[i:j] with an underlying array of capacity k
	// Length:   j-i
	// Capacity: k-i
	// For slice[1:3] with an underlying array of capacity 5
	// Length:   3-1=2
	// Capacity: 5-1=4
}

func PtrSlice(slice [3]*string) {
	*slice[0] = "Red"
	*slice[1] = "Blue"
	*slice[2] = "Green"
	fmt.Println(slice)
	fmt.Println(*slice[0], *slice[1], *slice[2])
}
