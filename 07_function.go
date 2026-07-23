package main

import "fmt"

func name(name string) string {
	return name
}

func main(){
	x := "Hello"
	intro := name(x)
	fmt.Println(intro)
}