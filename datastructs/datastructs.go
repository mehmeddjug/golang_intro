package datastructs

import "fmt"

func decleration() {
	// Declare an array of ten integers.
	var arrayOne [10]int
	fmt.Println(arrayOne)
	// Initialize each element with a specific value.
	arrayTwo := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrayTwo)
	// Capacity is determined based on the number of values initialized.
	arrayThree := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arrayThree)
	// Declare an array of ten integers.
	// Initialize specific values to elements with index 1 and 2.
	// The rest of the elements contain their zero value.
	arrayFour := [10]int{1: 10, 2: 9}
	fmt.Println(arrayFour)
}
