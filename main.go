package main

import (
	"fmt"

	"github.com/mehmeddjug/golangintro/datastructs"
	_ "github.com/mehmeddjug/golangintro/dummy"
)

func main() {
	fmt.Println("--------------------------------------------------")
	datastructs.Decleration()
	fmt.Println("--------------------------------------------------")
	arr := [3]string{"Gold", "Silver", "Bronze"}
	datastructs.Array(arr)
	fmt.Println("--------------------------------------------------")
	ptrArr := [3]*string{new(string), new(string), new(string)}
	datastructs.PtrArray(ptrArr)
	fmt.Println("--------------------------------------------------")
	fmt.Println("Hello World!")
}
