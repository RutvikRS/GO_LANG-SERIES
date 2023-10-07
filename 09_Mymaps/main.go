package main

import "fmt"

func main() {

	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages are :", languages)
	fmt.Println("JS SHorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages are :", languages)

	// loops are interetsing in go langs

	for key, value := range languages {
		fmt.Printf("For key %v,value is %v\n", key, value)
	}
}
