package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers ")
	// var ptr *int
	// fmt.Println("value of pointer is :", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("The value of pointer is : ", ptr)
	fmt.Println("The value of pointer is : ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("The value of pointer is :", *ptr)
}
