package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time stuff")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2025, time.April, 13,12,23,02,022,time.UTC)
	fmt.Println(createDate)


	var arr[4] int
	arr[0] = 1;
	arr[1] = 2;
	arr[3] = 5;

	fmt.Println(arr);

	var a = [3]int{-1,3,3}
	fmt.Println(a)



}