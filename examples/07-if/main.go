package main

import "fmt"

func main() {
	// Go has if, else if, else
	fmt.Println(".----------------------------------.")
	fmt.Println("| var | value |      if result     |")
	fmt.Println("| -------------------------------- |")
	// You can declare a variable inside an if statement
	if a := 10; a > 10 {
		fmt.Printf("|  a  |   %v  | is greater than 10 |\n", a)
	} else if a == 10 {
		fmt.Printf("|  a  |   %v  |      equals 10     |\n", a)
	} else {
		fmt.Printf("|  a  |   %v  |  is less than 10   |", a)
	}
	fmt.Printf(".----------------------------------.\n\n")

	// Go has if true/false
	var b bool = true
	if b {
		fmt.Println(".--------------------------.")
		fmt.Println("| var | value | if result  |")
		fmt.Println("| ------------------------ |")
		fmt.Printf("|  b  | %v  | b is true  |\n", b)
		fmt.Printf(".--------------------------.\n\n")
	} else {
		fmt.Println(".---------------------------.")
		fmt.Println("| var | value | if result   |")
		fmt.Println("| ------------------------- |")
		fmt.Printf("|  b  | %v | b is false  |\n", b)
		fmt.Printf(".---------------------------.\n\n")
	}
}
