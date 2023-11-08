package userdefineds

import "fmt"

type user struct {
	name string
	pass bool
	age  int16
}

func Decleration() {
	var someOne user
	someTwo := user{
		name: "Jon",
		pass: true,
		age:  25,
	}
	fmt.Println(someOne)
	fmt.Println(sometwo)
}
