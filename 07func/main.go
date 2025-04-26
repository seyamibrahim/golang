package main

import "fmt"

func adder(a int, b int) int {
	return a+b
}

func proadder(values ...int) int{
	total := 0
	for _,  val := range values {
		total += val
	}

	return total
}


func main(){
	fmt.Println("Welcome to func in golang")
	var a, b int
	fmt.Println("Enter numbers which you want to add : ")
	_, err := fmt.Scan(&a, &b)
	if err != nil{
		fmt.Println("there is something wrong...")
	}
	fmt.Println("Entered numbers is: ", a,"and", b)

	fmt.Printf("type of entered numbers is %T %T\n", a,b)



	res := adder(a,b)
	fmt.Println("Summation is : ", res)
	prores := proadder(1,2,3,4,5)
	fmt.Println("Proresult is : ", prores)
	
}