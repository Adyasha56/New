package main

import "fmt"

func main() {
	fmt.Println("learning about struct today")

	//no inheritance in golang
	//No super or parent, No child concept in golang
	adyashaNanda := User{"Adyasha", "adyasha12@gmail.com", true, 23}
	fmt.Println(adyashaNanda)
	fmt.Printf("the details of adyasha are: %+v\n", adyashaNanda)
	fmt.Printf("Name is : %v and email is: %v  ", adyashaNanda.Name, adyashaNanda.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
