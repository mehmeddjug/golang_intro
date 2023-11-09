package userdefineds

import "fmt"

type user struct {
	name string
	pass bool
	age  int16
}

type timeout int32

func Decleration() {
	// user defined type user (struct)
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

	// user defined type timeout (int32)
	delay := timeout(30)
	fmt.Println("Delay:", delay)

	someTwo.hello()
	someThree.updateName("Michael")
	someThree.hello()
}

// Method decleration -> func (type) name
// Call by value receiver
func (v user) hello() {
	fmt.Printf("Hello %s, I'm %d\n", v.name, v.age)
}

// Method decleration -> func (type) name
// Call by pointer receiver
func (v *user) updateName(name string) {
	v.name = name
}
