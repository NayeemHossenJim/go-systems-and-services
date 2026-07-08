package main

import "fmt"

func main() {
	val := "Jim"
	for index, char := range val {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
