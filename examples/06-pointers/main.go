package main

import "fmt"

func main() {

	// ** Adressing memory **

	x := 10
	// The & operator returns the memory address of the variable
	fmt.Println(".------------------------------.")
	fmt.Println("| Memory addres of variable x: |")
	fmt.Printf("|         %v         |\n", &x)
	fmt.Printf(".------------------------------.\n\n")

	// We are declaring a variable y that points to the memory address of x
	y := &x
	// The * operator returns the value of the memory address
	fmt.Println(".-----------------------------------------.")
	fmt.Println("| Derreferencing y to get the value of x: |")
	fmt.Printf("|                    %v                   |\n", *y)
	fmt.Printf(".-----------------------------------------.\n\n")

	// changing value of the memory address that y points to
	*y = 20
	// x is also changed
	fmt.Println(".---------------------------------------.")
	fmt.Println("| Changing value of *y then printing x: |")
	fmt.Printf("|                   %v                  |\n", x)
	fmt.Printf(".---------------------------------------.\n\n")

	// changing value of x
	x = 30
	// y value is also changed
	fmt.Println(".---------------------------------------.")
	fmt.Println("| Changing value of x then printing *y: |")
	fmt.Printf("|                   %v                  |\n", *y)
	fmt.Printf(".---------------------------------------.\n\n")

	// We can also declare a pointer variable
	// asserting that the variable will point to an integer
	var z *int = &x
	fmt.Println(".--------------------------------.")
	fmt.Println("| Printing a pointer variable z: |")
	fmt.Printf("|           %v         |\n", z)
	fmt.Printf(".--------------------------------.\n\n")

	// ** Pointers and functions **

	// no pointer function
	a := 10
	fmt.Println(".------------------------------------.")
	fmt.Println("|         No pointer function        |")
	fmt.Println("| ---------------------------------- |")
	fmt.Println("| var | Func Result | var after func |")
	fmt.Println("| ---------------------------------- |")
	fmt.Printf("| %v  |      %v     |       %v       |\n", a, xpto(a), a)
	fmt.Printf(".------------------------------------.\n\n")

	// pointer function
	b := 10
	fmt.Println(".------------------------------------.")
	fmt.Println("|          Pointer function          |")
	fmt.Println("| ---------------------------------- |")
	fmt.Println("| Var | Func Result | var after func |")
	fmt.Println("| ---------------------------------- |")
	fmt.Printf("| %v  |     %v      |       %v       |\n", b, xpto2(&b), b)
	fmt.Println(".------------------------------------.")
}

// no pointer function
func xpto(a int) int {
	a = a + 1
	return a
}

// pointer function
func xpto2(a *int) int {
	*a = *a + 1
	return *a
}
