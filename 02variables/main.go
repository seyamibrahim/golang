package main

import "fmt"


const Logintoken string = " slfjlajf;as" // public

func main() {

	var user string = "seyam"
	fmt.Println("variables")
	fmt.Println(user)
	fmt.Printf("varible if of type : %T \n", user)

	var isloggedIn bool = false
	fmt.Println(isloggedIn)
	fmt.Printf("varible if of type : %T \n", isloggedIn)

	var smallVal int = 256;
	fmt.Println(smallVal);
	fmt.Printf("var if of type : %T \n", smallVal)

	// default values and aliases
	var another int
	fmt.Println(another);
	fmt.Printf("var if of type : %T \n", another)

	var anothers string
	fmt.Println(anothers);
	fmt.Printf("var if of type : %T \n", anothers)

	name := "seyam"
	fmt.Println(name)
	fmt.Printf("var if of type : %T \n", name)


	fmt.Println(Logintoken)
	fmt.Printf("var if of type : %T \n", Logintoken)
}
