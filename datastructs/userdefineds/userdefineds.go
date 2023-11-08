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
		name: "James",
		pass: true,
		age:  25,
	}
	someThree := user{
		name: "John",
		pass: false,
		age:  35,
	}
	fmt.Println(someOne)
	fmt.Println(someTwo)
	fmt.Println(someThree)
}
