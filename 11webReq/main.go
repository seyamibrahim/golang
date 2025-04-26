package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://docs.chaicode.com"

func main() {
	fmt.Println("Web Request tut")
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response of type : %T\n", res)

	defer res.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}
