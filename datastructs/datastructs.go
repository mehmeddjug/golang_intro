package datastructs

func ArrayDecleration() {
	Decleration()
}

func ArrayAssignPtr() {
	ptrArr := [3]*string{new(string), new(string), new(string)}
	PtrArray(ptrArr)
}

func ArrayAssign() {
	arr := [3]string{"Gold", "Silver", "Bronze"}
	Array(arr)
}
