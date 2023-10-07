package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=abcdef"

func main() {
	fmt.Println("Welcome to Handling url")
	fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of Query Parameter is : %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)

		partsofUrl := &url.URL{
			Scheme:  "https",
			Host:    "lco.dev",
			Path:    "/tutcss",
			RawPath: "user=rutvik",
		}

		anotherURL := partsofUrl.String()

		fmt.Println(anotherURL)

	}

}
