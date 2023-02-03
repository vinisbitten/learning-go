package main

import "fmt"

/* function syntax:
func [func name] ([parameter list]) [return types] {
	[function body]
}
*/

// Simple example of a function
func Hey(a string) string {

	result := "Simple func:\n\n   " + "Hey " + a + "!\n"
	return result
}

// Multiple return values
func Hey2(a string) (string, string, string) {
	desc := "Multiple return values:\n"
	result := "Hey " + a + "!\n"
	return desc, result, a
}

// Variadic functions
func Hey3(a ...string) string {
	result := "Variadic func:\n\n   " + "Hey "
	for _, v := range a {
		result += v + ", "
	}
	result = result[:len(result)-2] + "!\n"
	return result
}

// Func inside a func
func Hey4() func() (string, string){
	Hey := "Hey"
	return func() (string, string) {
		return "Func inside func:\n", "   " + Hey + " Vini!"
	}
}

// Closure fibonacci example
func fibonacci() func() int {

	a := 0
	b := 1

	return func() int {

		 a, b = b, a+b
		 return b-a
	}
}

// main package has a special function called main
// that is the entry point of the program
func main() {
	fmt.Println(Hey("Vini")) // Hey Vini!

	// Multiple return values
	desc, greetings, user := Hey2("Vini")
	fmt.Println(desc)
	fmt.Println("   User:", user, "Result:", greetings)

	// Variadic functions
	users := Hey3("Vini", "Maria", "Ricardo", "Guilherme")
	fmt.Println(users)

	// Func inside a func
	hey := Hey4()
	desc, greetings = hey()
	fmt.Println(desc)
	fmt.Println(greetings)

	/* closure
	 A closure is a function value that references variables from outside its body.
	 The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	*/

	f := fibonacci() // fibonacci series using a closure

	fmt.Print("\nFibonacci series:\n\n   ")
	for i := 0; i < 10; i++ {
		 fmt.Print(f(), " ")
	}
	fmt.Print("\n\n")
	// Anonimous functions
	// A function without a name is called an anonymous function.
	rectangleArea := func(l int, b int) int {
		return l * b
	}

	fmt.Print("Anonimous function:\n\n")
	fmt.Println("   Area of a rectangle:", rectangleArea(10, 20))
}
