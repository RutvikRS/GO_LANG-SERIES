package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitlist = []string{"Apple", "Banana", "Mango"}
	fmt.Printf("Types of Fruitlist is %T/n", fruitlist)

	fruitlist = append(fruitlist, "Peach", "Orange")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highScore := make([]int, 4)

	highScore[0] = 800
	highScore[1] = 200
	highScore[2] = 900
	highScore[3] = 400

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "Javascript", "Swift", "Python", "Java"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
