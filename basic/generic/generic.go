package main

import "fmt"

// define a func with generic
// T: a generic type
// comparable: built-in interface of Go. allow to use ==, !=, ... on "T"
// s: a slice having type T
func Index[T comparable](s []T, x T) int {
	// loop each item in slice "s" by using for...range
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

// define a normal func(not use generic)
// mapInt(): has 2 params, return a new slice of int
// "arr": a slice
// "cb": a function
func mapInt(arr []int, cb func(int) int) []int {
	// create a new slice having same length with param "arr", using make()
	result := make([]int, len(arr))

	for i, v := range arr {
		// logic of cb: is define at calling point(in main())
		result[i] = cb(v) // pass "v" to the function "cb" then assign returned value from "cb" to result
	}

	return result
}

// re-write mapInt() in generic type
// K, V: any kind of type
// K: type of param passed to mapAny()
// V: type of returned value from mapAny()
func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))

	for i, v := range arr {
		// logic of f: is define at calling point(in main())
		result[i] = f(v) // pass "v" to the function "f" then assign returned value from "cb" to result
	}

	return result
}

func main() {
	// use generic of int
	si := []int{10, 20, 15, -10}
	foundIndex := Index(si, 15)
	println("generic of int: ", foundIndex)

	// use generic of string
	ss := []string{"foo", "bar", "baz"}
	foundIndex = Index(ss, "fb")
	println("generic of string: ", foundIndex)

	// use mapInt() function
	arr := []int{1, 2, 3, 4, 5}
	newArr := mapInt(arr, func(v int) int {
		return v * 2
	})
	fmt.Println("mapInt", newArr)

	// use mapAny() function
	newArrAny := mapAny(arr, func(v int) int {
		return v * 2
	})
	fmt.Println("mapAny", newArrAny)

}
