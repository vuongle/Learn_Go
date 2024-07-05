package printer

import "fmt"

// Define a struct that implements the "Printer" interface
type ConsolePrinter struct{}

func (cp ConsolePrinter) Print(message string) {
	fmt.Println("Console Printer:", message)
}
