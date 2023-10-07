package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("My Files")
	content := "This needs to go in a file - LearncodeOnline.com"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.text")

}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n ", string(databyte))

}
