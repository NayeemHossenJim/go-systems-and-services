package main

import "fmt"

const siteName = "GeeksforGeeks"

type Student struct {
	Name string
}

func display() {
	fmt.Println(siteName)
}

func main() {
	var age = 20
	display()
	fmt.Println(age)
	_, b := 10, 20
	fmt.Println(b)
}

/*
The underscore (_) is known as the blank identifier. It is used when a value is required syntactically but does not need to be stored.
Identifiers that begin with an uppercase letter are called exported identifiers. They can be accessed from other packages.
*/