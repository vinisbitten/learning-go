package main

import "fmt"

func main() {
	// declaring variables
	a := 10
	b := "hello"
	c := 4.17
	d := true
	// declaring string variable with backticks (can be multiline)
	e := `hello`

	// printing variables %v prints variable value, %T prints variable type
	// (var x: value | type: type)
	fmt.Println(".------------------------------.")
	fmt.Printf("| var a: %v    | type: %T     |\n", a, a)
	fmt.Printf("| var b: %v | type: %T  |\n", b, b)
	fmt.Printf("| var c: %v  | type: %T |\n", c, c)
	fmt.Printf("| var d: %v  | type: %T    |\n", d, d)
	fmt.Printf("| var e: %v | type: %T  |\n", e, e)
	fmt.Println(".------------------------------.")
}
