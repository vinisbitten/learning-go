package main

import (
	"fmt"
)

func main() {

	// Infinite loop

	sum0 := 0
	for {
		sum0++ // repeated forever
		if sum0 > 10 {
			break // break out of the loop
		}
	}
	fmt.Println(".------------.")
	fmt.Printf("| sum 0 | %v |\n", sum0) // 11
	fmt.Printf(".------------.\n\n")

	// Classic for loop

	sum1 := 0
	for i := 1; i < 5; i++ {
		sum1 += i
	}
	fmt.Println(".------------.")
	fmt.Printf("| sum 1 | %v |\n", sum1) // 10 (1+2+3+4)
	fmt.Printf(".------------.\n\n")

	// While loop

	n := 3
	for n < 5 {
		n *= 2
	}

	fmt.Println(".----------------.")
	fmt.Printf("| while loop | %v |\n", n) // 8 (1*2*2*2)
	fmt.Printf(".----------------.\n\n")

	// For-each range loop

	strings := []string{"hello", "world"}
	fmt.Println(".---------------------------.")
	for i, s := range strings {
		fmt.Printf("| position: %v |value: %v |\n", i, s)
	}
	fmt.Printf(".---------------------------.\n\n")

	// Jumping to the next iteration

	sum3 := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum3 += i
	}
	fmt.Println(".-----------.")
	fmt.Printf("| sum 3 | %v |\n", sum3) // 10 (1+2+3+4)
	fmt.Printf(".-----------.\n\n")

}
