package main

import "fmt"

func main() {
	const language = "Go"

	for i := 1; i <= 3; i++ {
		if i == 2 {
			fmt.Println(language)
		}
	}
}
