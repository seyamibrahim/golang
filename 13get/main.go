package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello this is tut of get req using golang")
	PerformGetReq()

}

func PerformGetReq() {

	const myurl = "http://localhost:3000/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status Code : ", res.StatusCode)
	fmt.Println("Content lenght is : ", res.ContentLength)

	// var responseString strings.Builder
	// content, _ := ioutil.ReadAll(res.Body)
	// bytecount, _ := responseString.Write(content)
	// fmt.Println(responseString.String())
	// fmt.Println(string(content))
}
