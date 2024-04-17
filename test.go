package main

import (
	"fmt"
	"math"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	defer fmt.Println("THIS LINE PRINTED LAST ONCE the program exited")

	// ------------------------------ variables and constants ------------------------------
	var i int = 12
	f := float32(i) // this is called "type conversion"
	isGo := true    //this is called "type inference": tu dong xac dinh kieu du lieu
	const name string = "golang"
	const description = "a robust language for backend development"

	// ------------------------------  if-else ------------------------------
	// if v := math.Pow(2, 3); v < 10 ==> this way is same as 2 following lines
	// v := math.Pow(2, 3)
	// if v < 10
	if v := math.Pow(2, 3); v < 10 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	// ------------------------------  switch-case ------------------------------
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorrow")
	default:
		fmt.Println("Too far away")
	}

	// ------------------------------ for, for...range ------------------------------
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// for-range: like for...each
	//1.define a slice
	pow := []int{1, 2, 4, 8, 16, 32, 64}
	// i,v := range means that: want to get both index and value for each loop
	// i: index
	// v: value
	for i, v := range pow {
		fmt.Printf("index in range: %d -> value in range: %d\n", i, v)
	}
	// get only index
	for i := range pow {
		fmt.Printf("index in range: %d \n", i)
	}
	// get only value
	for _, v := range pow {
		fmt.Printf("value in range: %d\n", v)
	}

	// ------------------------------  struct: same as object in other languages ------------------------------
	// 1. define a struct
	type Vertex struct {
		X int
		Y int
	}
	//2.create a var and assign a struct
	v := Vertex{4, 7}
	fmt.Println(v.X)
	//3.change value of created struct
	v.X = 33
	fmt.Println(v.X)

	// ------------------------------  arrays(fix size) ------------------------------
	var names [2]string
	names[0] = "one"

	primes := [7]int{1, 2, 3, 4, 9, 23, 45}

	// ------------------------------ slices(dynamic size) ------------------------------
	// create slice from array
	s := primes[3:]
	fmt.Println("slice\"s\": ", s)
	// create slice from beginning(not from another array)
	s2 := []int{5, 3, 9, 10}
	fmt.Printf("slice \"s2\" len: %d, cap: %d \n", len(s2), cap(s2))
	s3 := append(s2, 111)
	fmt.Println("Append slice: ", s3)

	fmt.Printf("Call a func: %d", add(i, 34))
	fmt.Println(isGo, f, name, today, names, primes)

	// ------------------------------ map ------------------------------
	type Location struct {
		Lat, Lon float64
	}

	// key is string
	// value is Location
	var m = map[string]Location{
		"Bell_Labs": {
			40.78989, -79.000876,
		},
		"Google": {
			3.8978765, 12.0905654,
		},
	}
	fmt.Println("map: ", m)
	fmt.Println("get value from map: ", m["Google"])
	m["HCM"] = Location{7.096440000, -17.09097001} // insert or update map
	delete(m, "Google")                            // delete map by key

	// ------------------------------ function values ------------------------------
	// func(x, y float64) float64: this is a function. it is used as assignment value
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("function value as return values: ", hypot(5, 12))
	fmt.Println("function value as function arguments: ", compute(hypot))
}

// ------------------------------ function values ------------------------------
// func(float64, float64) float64: this is a func, receive 2 prams and return float64
// it is passed as function arguments of the compute()
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4) // call the function that has been passed into compute()
}
