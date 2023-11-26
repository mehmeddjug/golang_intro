package main

import (
	"fmt"

	"github.com/mehmeddjug/golangintro/concurrency"
	"github.com/mehmeddjug/golangintro/datastructs"
	_ "github.com/mehmeddjug/golangintro/dummy"
)

func main() {
	fmt.Println("Hello Arrays!")
	fmt.Println("--------------------------------------------------")
	datastructs.ArrayDecleration()
	fmt.Println("--------------------------------------------------")
	datastructs.ArrayAssign()
	fmt.Println("--------------------------------------------------")
	datastructs.ArrayAssignPtr()
	fmt.Println("--------------------------------------------------")
	fmt.Println("Hello Slices!")
	fmt.Println("--------------------------------------------------")
	datastructs.SliceDecleration()
	fmt.Println("--------------------------------------------------")
	datastructs.SliceAssign()
	fmt.Println("--------------------------------------------------")
	datastructs.MapDecleration()
	fmt.Println("--------------------------------------------------")
	datastructs.MapAssign()
	fmt.Println("--------------------------------------------------")
	datastructs.UserdefinedDecleration()
	fmt.Println("--------------------------------------------------")
	datastructs.InterfacedsDecleration()
	fmt.Println("--------------------------------------------------")
	fmt.Println("Concurrency examples:")
	fmt.Println("--------------------------------------------------")
	concurrency.ConcurrencyExample()
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")
	concurrency.LockSharedResourcesExample()
	fmt.Println("--------------------------------------------------")
}
