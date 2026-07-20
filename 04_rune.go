package main

import "fmt"

func main() {
	// Declaring a rune variable
	var r rune = 'A'
	fmt.Println("Rune value is:", r)
	fmt.Printf("The type of r is: %T\n", r)

	// Converting a string to a slice of runes
	str := "Hello, 世界"
	runes := []rune(str)
	fmt.Println("Slice of runes:", runes)
}
