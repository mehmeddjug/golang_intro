package datastructs

import (
	"github.com/mehmeddjug/golangintro/datastructs/arrayds"
	"github.com/mehmeddjug/golangintro/datastructs/mapds"
	"github.com/mehmeddjug/golangintro/datastructs/sliceds"
	"github.com/mehmeddjug/golangintro/datastructs/userdefineds"
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

func MapDecleration() {
	mapds.Decleration()
}

func MapAssign() {
	mapDict := map[string]string{"One": "One-Two", "Three": "Three-Four"}
	mapds.Map(mapDict, "Five", "Five-Six")
}

func UserdefinedDecleration() {
	userdefineds.Decleration()
}
