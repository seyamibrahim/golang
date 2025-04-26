package main

import "fmt"

func main() {
	seyam := User{"seyam", "seyam@.in", 15, true}

	seyam.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)

}