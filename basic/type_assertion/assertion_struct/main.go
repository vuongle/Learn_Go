package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Animal struct {
	name string
	age  int
}

func main() {
	var p interface{}
	p = Person{name: "Alice", age: 25}

	// Type assertion to access the name, age field
	fmt.Println(p.(Person).name)
	fmt.Println(p.(Person).age)

	// an example of FAIL
	// panic: interface conversion: interface {} is main.Animal, not main.Person
	p = Animal{name: "Dog", age: 5}
	fmt.Println(p.(Person).name)
}
