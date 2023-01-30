package main

import "fmt"

func main() {

	// ** Go has switch statements **
	a := "Vini"
	switch a {
	case "Maria":
		fmt.Println("Hey Maria!")
	case "Ricardo":
		fmt.Println("Hey Ricardo!")
	case "Guilherme":
		fmt.Println("Hey Guilherme!")
	default:
		fmt.Println("I did not find your name!")
	}
}
