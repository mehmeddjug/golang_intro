package datastructs

import (
	"github.com/mehmeddjug/golangintro/datastructs/arrayds"
	"github.com/mehmeddjug/golangintro/datastructs/sliceds"
)

func ArrayDecleration() {
	arrayds.Decleration()
}

func ArrayAssignPtr() {
	ptrArr := [3]*string{new(string), new(string), new(string)}
	arrayds.PtrArray(ptrArr)
}

func ArrayAssign() {
	arr := [3]string{"Gold", "Silver", "Bronze"}
	arrayds.Array(arr)
}

func SliceDecleration() {
	sliceds.Decleration()
}

func SliceAssign() {
	arr := []string{"Gold", "Silver", "Bronze"}
	sliceds.Slice(arr)
}
