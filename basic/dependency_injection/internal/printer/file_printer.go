package printer

import (
	"fmt"
)

// Define a struct that implements the "Printer" interface
type FilePrinter struct {
	FilePath string
}

func (fp FilePrinter) Print(message string) {
	fmt.Println("File Printer:", message)
}
