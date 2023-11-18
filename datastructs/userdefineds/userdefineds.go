package userdefineds

import "fmt"

type userOne struct {
	name string
	pass bool
	age  int16
}

type userTwo struct {
	name string
	pass bool
	age  int16
}

type timeout int32

type helloer interface {
	hello()
}

func sayHello(h helloer) {
	h.hello()
}

func Decleration() {
	// user defined type user (struct)
	var someOne userOne

	someTwo := userOne{
		name: "James",
		pass: true,
		age:  25,
	}
	someThree := userOne{"John", false, 35}

	fmt.Println(someOne)
	fmt.Println(someTwo)
	fmt.Println(someThree)

	// user defined type timeout (int32)
	delay := timeout(30)
	fmt.Println("Delay:", delay)

	someThree.hello()
	someTwo.hello()
	someThree.updateName("Michael")
	someThree.hello()
}

// Method decleration -> func (type) name
// Call by value receiver
func (v userOne) hello() {
	fmt.Printf("Hello %s, I'm %d\n", v.name, v.age)
}

// Method decleration -> func (type) name
// Call by pointer receiver
func (v *userOne) updateName(name string) {
	v.name = name
}

// Method decleration -> func (type) name
// Call by value receiver
func (v userTwo) hello() {
	fmt.Printf("Hello %s\n", v.name)
}

// Method decleration -> func (type) name
// Call by pointer receiver
func (v *userTwo) updateName(name string) {
	v.name = name
}

func DeclerationWithInterface() {
	someOne := userOne{"John", false, 35}

	someTwo := userTwo{
		name: "James",
		pass: true,
		age:  25,
	}
	fmt.Printf("Interface:\n")
	sayHello(someOne)
	sayHello(someTwo)
	fmt.Printf("End of functions.\n")
}
