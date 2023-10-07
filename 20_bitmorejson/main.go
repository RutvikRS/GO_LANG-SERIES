package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"password"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON ")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "RS2123", []string{"web-dev", "js"}},
		{"Mern", 199, "LearnCodeOnline.in", "abc123", []string{"Express", "node.js"}},
		{"Angular", 209, "LearnCodeOnline.in", "AC123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"name": "ReactJS Bootcamp",
		"price": 299,
		"website": "LearnCodeOnline.in",
		"password": "RS2123",
		"tags": ["web-dev","js"]
}
	`)

	var lcoCourse course

	checkjson := json.Valid(jsonDataFromWeb)

	if checkjson {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where are you just want to add data to key value

	var myOnlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlinedata)
	fmt.Printf("%#v\n", myOnlinedata)

	for k, v := range myOnlinedata {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
