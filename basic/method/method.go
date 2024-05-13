package main

import "fmt"

// ------------------------------ Method ------------------------------
// step1: define a struct
type Position struct {
	X, Y float64
}

// step2: define a method(func + receiver becomes a method)
// (p Position): is called a "value receiver" of a method
// after this step: GetCurrentPostion belongs to Position struct
func (p Position) GetCurrentPostion() float64 {
	return p.X + p.Y
}

// define a method with "pointer receiver"
func (p *Position) Scale(f float64) {
	// Change value of "p"
	p.X = p.X * f
	p.Y = p.Y * f
}

// define a method with "value receiver"
// implementation inside "Scale()" and "Scale2()" is same
// but Scale() does change value of struct
// 	   Scale2() does not change
func (p Position) Scale2(f float64) {
	// Change value of "p"
	p.X = p.X * f
	p.Y = p.Y * f
}

type Employee struct {
	name string
	age  int
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName

	// e: this struct is copied then assign newName
	// therefore, it is different from "e" in main()
	fmt.Println("Employee INSIDE change name: ", e)
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
	// e: this struct is NOT copied
	// therefore, it is same as "e" in main()
}

func main() {
	p := Position{8, 12}
	fmt.Println("use method: ", p.GetCurrentPostion())
	fmt.Println("before scale: ", p)
	//p.Scale(3) // change value of "p" because "p" is a poiter receiver
	p.Scale2(3) // not change of "p" because "p" is a value receiver
	fmt.Println("after scale: ", p)

	//
	e := Employee{
		name: "Mark",
		age:  50,
	}
	fmt.Println("Employee before change name/age: ", e)
	e.changeName("Andrew")
	e.changeAge(51)

	fmt.Println("\nEmployee after change name/age: ", e)
}
