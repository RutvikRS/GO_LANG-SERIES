package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Web Requests")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of response is :%T/n", response)

	defer response.Body.Close() //Caller responsibilty is to close the connectivity

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}
