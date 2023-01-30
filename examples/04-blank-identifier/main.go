package main

import (
	"fmt"
	"io"
	"net/http"
)

// Here we are making a GET request to google.com and printing the response body
func main() {
	// The blank identifier ( underscore ) is used to ignore the value of a variable
	res, _ := http.Get("http://www.google.com/")
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	// Since body is a byte array, we need to convert it to a string
	// %s is used to print a variable as a string
	fmt.Printf("%s\n", body)
	fmt.Printf("\n.-------------------------------------------------------.\n")
	fmt.Println("| The result above is the response body from google.com |")
	fmt.Println(".-------------------------------------------------------.")
}
