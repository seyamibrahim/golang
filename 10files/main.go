package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files tut in golang")

	content := "this needs to go in file -gagroups.in"
	file, err := os.Create("./hellofiles.txt")

	checkNilErr(err)

	len, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is : ", len)
	defer file.Close()
	readFile("./hellofiles.txt")
}


func readFile(filename string){

	databyte , err := ioutil.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("Text data inside the file is : \n", string(databyte))

}

func checkNilErr(err error)  {
	if(err != nil){
		panic(err)
	}
}