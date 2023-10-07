package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Welcome to ArthaVeda"
	fmt.Println(Welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(("Enter the rating of our Pizza:"))

	//comma ok ||err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input)
}
