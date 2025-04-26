package main

import "fmt"

func main() {

	// fmt.Println("Welcome to loops tut")
	// days := []string{"Sunday", "monday", "tuesday", "wednesday", "thursday","friday"}

	// for i:= 0; i < 6; i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days{
	// 	fmt.Printf("index is %v and value is %v\n", i, day)
	// }

	count := 1
	for count < 10 {
		if count == 2 {
			goto land
		}
		if count == 3 {
			count++
			continue
		}
		fmt.Println("count is : ", count)
		count++
	}

land:
	fmt.Println("goto statement landed on this line")
}
