package main

import "fmt"

func main() {

	//
	// The following code is a simple example of type assertion, without panic
	//
	var i interface{} = "hello"
	s := i.(string) // type assertion
	fmt.Println(s)

	//
	// The following code causes a panic because not check the assertion
	//
	// number := i.(int)
	// fmt.Println(number)

	//
	// To resolve the panic, perform a safe assertion
	//
	number2, ok := i.(int)
	if ok {
		fmt.Println("i is integer", number2)
	} else {
		fmt.Println("i is not integer")
	}
}
