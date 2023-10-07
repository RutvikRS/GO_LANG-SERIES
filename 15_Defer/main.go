package main

import "fmt"

func main() {
	defer fmt.Println("Three ")
	defer fmt.Println(" Two ")
	defer fmt.Println(" One ")
	myDefer()

}

//One,Two

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)

	}
}
