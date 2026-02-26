package main

import "fmt"

const LoginToken string = "fdsksjk"

func main() {
	// fmt.Println("variables")

	var username string = "adyasha"
	fmt.Println(username)
	fmt.Printf("the username is of type : %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("the isLoggedin is of type : %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("the smallVal is of type : %T \n", smallVal)

	var smallFloatVal float32 = 255.45679797020
	fmt.Println(smallFloatVal)
	fmt.Printf("the smallFloatVal is of type : %T \n", smallFloatVal)

	//default values & alias

	var anotherVariable int
	fmt.Println(anotherVariable) //0
	fmt.Printf("the anotherVariable is of type : %T \n", anotherVariable)

	//implicit way to assign a value
	var name = "adyasha"
	fmt.Println(name)
	fmt.Printf("the name is of type : %T \n", name)

	//declaration without var
	myName := "haadsaa"
	fmt.Println(myName);

}
