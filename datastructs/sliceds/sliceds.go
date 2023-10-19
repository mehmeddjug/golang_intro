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

	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice := []int{1, 2, 3, 4, 5}

	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	newSlice := slice[1:3]

	// By changing index 1 of newSlice, index 2 of the original slice is also changed.
	newSlice[1] = 7
	// newSlice[3] = 13 runtime error
	// Allocate a new element from capacity and assign 8 to the new element.
	newSlice = append(newSlice, 8)
	newSliceTwo := append(slice, 9)
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(newSliceTwo)
	// For slice[i:j] with an underlying array of capacity k
	// Length:   j-i
	// Capacity: k-i
	// For slice[i:j:k]
	// For slice[1:3] with an underlying array of capacity 5
	// Length:   3-1=2
	// Capacity: 5-1=4
	fmt.Println("Range with index and value:")
	for index, value := range newSliceTwo {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	fmt.Println("Range with blank identifier and value:")
	for _, value := range newSliceTwo {
		fmt.Printf("Value: %d\n", value)
	}
	fmt.Println("For loop:")
	for index := 2; index < len(newSliceTwo); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, newSliceTwo[index])
	}
	// Create a slice of a slice of integers.
	slicemd := [][]int{{3}, {15, 19}}
	fmt.Println("Multidimensional slice:")
	fmt.Println(slicemd)
	// Append the value of 20 to the first slice of integers.
	slicemd[0] = append(slicemd[0], 23)
	fmt.Println("Append 23 to first element:")
	fmt.Println(slicemd)
}

// Slice has pointer, length and capacity so the array will not be copied to stack
func Slice(sliceInput []string) {
	// Create a slice of strings with length and capacity of 5 elements.
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// Slice the third element and restrict the capacity.
	// Contains a length of 1 element and capacity of 2 elements.
	slice := source[2:3:4]
	// slice := source[2:3:13] runtime error

	// Append a new string to the slice.
	slice = append(slice, "Kiwi")
	fmt.Println(source)
	fmt.Println(slice)

	// Create a slice of strings with length and capacity of 5 elements.
	sourceTwo := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// Slice the third element and restrict the capacity.
	// Contains a length and capacity of 1 element.
	sliceTwo := sourceTwo[2:3:3]
	// Append a new string to the slice.
	sliceTwo = append(sliceTwo, "Kiwi")
	fmt.Println(sourceTwo)
	fmt.Println(sliceTwo)

	// Append the two slices together and display the results.
	sliceThree := append(source, sourceTwo...)
	fmt.Printf("%v\n", sliceThree)

	// Iterate over each element and display each value.
	for index, value := range sliceThree {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}
	fmt.Println("Print function input:")
	for index, value := range sliceInput {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}
}
