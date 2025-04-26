package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&payment=idbjh12jd"

func main() {
	fmt.Println("Welcome to handling urls in golang")

	fmt.Println(myurl)

	// result, _ := url.Parse(myurl)
	// // fmt.Println(result.Scheme)
	// // fmt.Println(result.Host)
	// // fmt.Println(result.Path)
	// // fmt.Println(result.RawQuery)
	// // fmt.Println(result.RawPath)
	// // fmt.Println(result.Port())
	// // fmt.Println(result.Hostname())
	// // fmt.Println(result.Query())

	// qparams := result.Query()
	// fmt.Printf("The type of query params are : %T\n", qparams)

	// for key, val := range qparams{
	// 	fmt.Println("Key is ", key, "value is ", val)

	// }

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "docs.chaicode.com",
		Path:   "/youtube/chai-aur-git",
		// RawQuery: "user=seyam",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
