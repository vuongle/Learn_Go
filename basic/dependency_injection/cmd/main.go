package main

import (
	"dependency_injection/internal/application"
	"dependency_injection/internal/printer"
)

func main() {

	// instantiate different printer implementations
	consolePrinter := printer.ConsolePrinter{}
	filePrinter := printer.FilePrinter{FilePath: "output.txt"}

	// Inject implementations (dependencies) into the struct
	app1 := application.NewApplication(consolePrinter)
	app2 := application.NewApplication(filePrinter)

	app1.Printer.Print("Hello from App1")
	app2.Printer.Print("Hello from App2")
}
