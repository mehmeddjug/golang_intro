package userdefineds

import "fmt"

type user struct {
	name string
	pass bool
	age  int16
}

func Decleration() {
	var someOne user
	fmt.Println(someOne)
}
