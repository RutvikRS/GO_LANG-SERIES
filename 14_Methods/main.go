package main

import "fmt"

func main() {
	fmt.Println("Methods in go langs")
	// no inheritance in go langs: No Super and no parents

	Rutvik := user{"Rutvik", "rutviksakpal@gmail.com", true, 22}
	fmt.Println(Rutvik)
	fmt.Printf("Rutvik details are: %+v\n", Rutvik)
	fmt.Printf("Name is %v and Email is %v", Rutvik.Name, Rutvik.Email)
	Rutvik.GetStatus()
	Rutvik.NewMail()
}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u user) GetStatus() {
	fmt.Println(" : Is user active: ", u.Status)
}

func (u user) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
