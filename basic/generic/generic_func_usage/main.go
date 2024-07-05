package main

import "fmt"

type PrintInterface interface {
	print()
}

// Define a specific type that implements the PrintInterface interface
type PrintNum int

func (p PrintNum) print() {
	fmt.Println("Num:", p)
}

// Define another specific type that implements the PrintInterface interface
type PrintText string

func (p PrintText) print() {
	fmt.Println("Text:", p)
}

// Define a generic func that accepts any slices that implement the PrintInterface interface
// This is a kind of "Type Constraints"
func PrintSlice[T PrintInterface](t []T) {
	for _, value := range t {
		value.print()
	}
}

func main() {
	// Example with PrintSlice
	stringSlice := []PrintText{"apple", "banana", "orange"}
	intSlice := []PrintNum{5, 2, 9, 1, 7}

	fmt.Println("String slice:")
	PrintSlice(stringSlice)

	fmt.Println("\nInteger slice:")
	PrintSlice(intSlice)
}
