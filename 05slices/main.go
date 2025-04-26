package main

import "fmt"

func main() {
	fmt.Println("This is giving the idea of slices")

	// var fl = []string{}
	// fl = append(fl, "apple", "mango", "banana", "grapes")
	// fmt.Println(fl)
	
 	
	// scores := make([]int,0)

	// scores = append(scores, 12,54,11,98)
	// fmt.Println(scores)

	// sort.Ints(scores)
	// fmt.Println(scores)

	// fmt.Println(sort.IntsAreSorted(scores))

	// // how to remove a value from slices based on indexconst
	// var courses = []string{"reactjs","javascript", "python", "java", "MongoDB"}

	// fmt.Println(courses)

	// ind := 2;

	// courses = append(courses[:ind], courses[ind+1:]...)
	// fmt.Println(courses)

	// langs := make(map[string]string)

	// langs["js"] = "jvas"
	// langs["go"] = "GO"
	// langs["ru"] = "Ruby"
	// langs["cot"] = "Cotlin"
	// langs["py"] = "Python"
	// fmt.Println(langs)
	// delete(langs, "ru")
	// fmt.Println(langs)
	// for key, val := range langs{
	// 	fmt.Printf("for key %v , value is %v\n", key, val)
	// }

	// seyam := User{"seyam", "seyam.ai", 20}

	// fmt.Println(seyam)

	// fmt.Printf("seyam details are : %+v\n", seyam)



}

type User struct{
	Name string
	Email string
	age int
}