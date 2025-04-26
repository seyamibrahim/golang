package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")
	mydefer()

	
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}