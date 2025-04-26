package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to post request tutorial in golang")
	PostJsonReq()

}

func PostJsonReq() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{

	 	"name" : "seyamibrahim and sjhflsadfkl;as",
		"age" : 15,
		"email" :          "xyx@exmaple.com"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
