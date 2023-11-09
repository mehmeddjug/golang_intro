package userdefineds

import "fmt"

type user struct {
	name string
	pass bool
	age  int16
}

type timeout int32

func Decleration() {
	var someOne user

	someTwo := user{
		name: "James",
		pass: true,
		age:  25,
	}
	someThree := user{"John", false, 35}

	fmt.Println(someOne)
	fmt.Println(someTwo)
	fmt.Println(someThree)

	delay := timeout(30)
	fmt.Println("Delay:", delay)

	someTwo.hello()
}

func (user user) hello() {
	fmt.Printf("Hello %s, I'm %d\n", user.name, user.age)
}
