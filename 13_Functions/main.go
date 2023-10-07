package main

import "fmt"

func main() {
	fmt.Println("Welcomes to functions")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is:", result)

	proRes := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Namastey to golang")

}
