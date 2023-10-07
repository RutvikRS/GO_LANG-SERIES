package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in go lang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Chiku"

	fmt.Println("The Fruitlist is :", fruitList)

	fmt.Println("The Fruitlist is :", len(fruitList))

	var veglist [4]string

	veglist[0] = "Onion"
	veglist[0] = "Potato"
	veglist[0] = "Cucumber"
	fmt.Println("The Veglist is :", len(veglist))
}
