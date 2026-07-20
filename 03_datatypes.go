package main

import "fmt"

func main() {
	a := 20.45
	b := 34.89

	c := b - a
	fmt.Println("Result is:", c)

	fmt.Println("The type of c is :", c)

	str1 := "GeeksforGeeks"
	str2 := "geeksForgeeks"
	str3 := "GeeksforGeeks"
	result1 := str1 == str2
	result2 := str1 == str3

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("The type of result1 is %T and "+ "the type of result2 is %T", result1, result2)

	fmt.Println("New string : ", str1+str2)
}