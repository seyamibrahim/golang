package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "os"
	// "strconv"
	// "strings"
)

func main() {

	welcome := "Welcome to Go programming"
	fmt.Println(welcome)
	fmt.Println("Give rating for our pizza")

	// take input with white spaces
	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')
	// fmt.Println("Rating given by user is : ", input)
	

	// newinput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Added 1 to your rating : ", newinput+1)
	// }
	

	// var x int 
	// var y string
	// _ , err :=  fmt.Scan(&x, &y)
	
	// fmt.Println("You entered x : " , x , " and y : ", y, " Error : ", err)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("enter rating is : ", input)
	fmt.Printf("input is of type %T", input)
	fmt.Println()
	a, err  := strconv.ParseInt(strings.TrimSpace(input), 10, 32)


	if err != nil {
		fmt.Println("Error in taking input: ", err)
	} else{
		fmt.Println("Added 3 to input : ", a + 3);
	}
}