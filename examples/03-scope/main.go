package main

import (
	"fmt"
	// Importing other packages
	goodbye "scope/other-package"
)

// global variables
var (
	a int
	b bool
)

func main() {
	// since global variables are declared, we can use them anywhere
	a = 1
	b = true
	printVars()

	/*
		Trynig to print c and d will give us an error
		because they are declared inside the printVars function
		therefore they are outside of the scope of the main function
	*/
	// fmt.Printf("| var c: %v   | type: %T    |\n", c, c)

	// We can also access other files variables and functions if they are in the same package
	fmt.Printf("\nAccessing func from other file (inside package)...\n\n")
	fmt.Println(".-----------------------.")
	fmt.Printf("| %v |\n", printText())
	fmt.Println(".-----------------------.")

	// We can also access functions from other packages if we import them
	// using package goodbye from other-package
	fmt.Printf("\nAccessing func from other file (outside package)...\n\n")
	fmt.Println(".-------------------------.")
	fmt.Printf("| %v |\n", goodbye.Say())
	fmt.Println(".-------------------------.")

	/*
		Trying to access functions from other packages that are not declared with a capital letter will give us an error
	*/
	// fmt.Printf("| %v |\n", goodbye.scream())
}

func printVars() {
	// local variables

	//We can declare varibles this way
	var (
		c int
		d bool
	)
	//And then assign values to them
	c = 2
	d = false

	// or we can declare and assign values to them in one line
	//c, d := 2, false || c := 2; d := false

	// we can even use global variables inside this function
	fmt.Printf("\nPrinting global variables...\n\n")
	fmt.Println(".--------------------------.")
	fmt.Printf("| var a: %v    | type: %T  |\n", a, a)
	fmt.Printf("| var b: %v | type: %T |\n", b, b)
	fmt.Println(".--------------------------.")
	// printing c and d declared inside this function
	fmt.Printf("\nPrinting local variables...\n\n")
	fmt.Println(".---------------------------.")
	fmt.Printf("| var c: %v   | type: %T    |\n", c, c)
	fmt.Printf("| var d: %v | type: %T |\n", d, d)
	fmt.Println(".---------------------------.")
}
