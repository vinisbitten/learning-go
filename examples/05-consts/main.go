// Consts are not considered a

package main

import "fmt"

// A const can't be defined without a value
// You can declare a const whithout declering the type
// An untyped consts has a hidden type that is inferred from the value
const a int = 10

// You can also declare a const with its type
const b string = "Hello"

// you can also declare multiple consts at once
const (
	c = 20
	d = "-"
	e = "World"
)

func main() {
	// You can add an untyped const to a typed const
	fmt.Println(".--------------------------------.")
	fmt.Println("| Values bellow are all constans |")
	fmt.Println("|--------------------------------|")
	fmt.Println("| a + c:      |      ", a+c, "        |")
	fmt.Println("| b + e + d:  |  ", b+d+e, "   |")
	fmt.Println(".--------------------------------.")
}
