package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Password string   `json:"-"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json tut")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", "abc123", 199, "LearnCodeOnline.in", []string{"web-dev", "js"}},
		{"MERN Bootcamp", "bcd123", 299, "LearnCodeOnline.in", []string{"full-stack", "js"}},
		{"Go Bootcamp", "sey123", 399, "LearnCodeOnline.in", nil},
	}

	//package this data as Json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
func DecodeJson() {

	JsonDataFromWeb := []byte(`
		 {
                "coursename": "ReactJS Bootcamp",
                "price": 199,
                "website": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	`)

	var lcocourse course

	isValid := json.Valid(JsonDataFromWeb)

	if isValid {
		fmt.Println("Data is valid")
		json.Unmarshal(JsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	}else{
		fmt.Println("Data is not Valid")
	}

	var mydata map[string]interface{}

	json.Unmarshal(JsonDataFromWeb, &mydata)
	fmt.Printf("%#v\n", mydata)

	for key, val := range mydata{
		fmt.Printf("key is %v and value is %v and type is %T\n", key, val, val)
	}
}
