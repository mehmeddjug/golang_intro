package arrayds

import "fmt"

// Function with starting capital letter is public function/method
func Decleration() {
	// Declare an array of ten integers.
	var arrayOne [10]int
	fmt.Println(arrayOne)

	// Initialize each element with a specific value.
	arrayTwo := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrayTwo)

	// Capacity is determined based on the number of values initialized.
	arrayThree := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arrayThree)

	// Initialize elements with index 1 and 2.
	// The rest of the elements contain their zero value.
	arrayFour := [10]int{1: 10, 2: 9}
	fmt.Println(arrayFour)

	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	array := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*array[0] = 3
	*array[1] = 7
	fmt.Println(array)
	fmt.Println(*array[0], *array[1])

	// Declare two dimensional arrays of integers.
	var arraymd1 [2][2]int
	var arraymd2 [2][2]int
	// Add values to elements.
	arraymd2[0][0] = 11
	arraymd2[0][1] = 12
	arraymd2[1][0] = 21
	arraymd2[1][1] = 22
	fmt.Println(arraymd2)
	// Copy arraymd2 to arraymd1.
	arraymd1 = arraymd2
	fmt.Println(arraymd1)
}

func Array(array1 [3]string) {
	array2 := [3]string{"Red", "Blue", "Yellow"}
	fmt.Println(array2)
	array2 = array1
	fmt.Println(array2)
}

func PtrArray(array [3]*string) {
	*array[0] = "Red"
	*array[1] = "Blue"
	*array[2] = "Green"
	fmt.Println(array)
	fmt.Println(*array[0], *array[1], *array[2])
}
