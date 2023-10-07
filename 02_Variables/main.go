package main

import "fmt"

func main() {
	var username string = "Rutvik Sakpal"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var Smallval uint8 = 204
	fmt.Println(Smallval)
	fmt.Printf("Variable is of type: %T \n", Smallval)

	var Smallfloat float32 = 204.555555555
	fmt.Println(Smallfloat)
	fmt.Printf("Variable is of type: %T \n", Smallfloat)

	//default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	//implicit type
	var website = "Welcome to Arthaveda"
	fmt.Println(website)

	//no var style
	noofuser := 2000
	fmt.Println(noofuser)
}
