package application

import "dependency_injection/internal/printer"

// Define a stuct that has a dependency on the "Printer" interface
type Application struct {
	Printer printer.Printer
}

// Inject the dependency into the struct
func NewApplication(printer printer.Printer) *Application {
	return &Application{
		Printer: printer,
	}
}
