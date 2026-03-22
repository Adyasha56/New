package main

import "fmt"

func main() {
	fmt.Println("learning about methods today")

	//no inheritance in golang
	//No super or parent, No child concept in golang
	adyashaNanda := User{"Adyasha", "adyasha12@gmail.com", true, 23}
	fmt.Println(adyashaNanda)
	fmt.Printf("the details of adyasha are: %+v\n", adyashaNanda)
	fmt.Printf("Name is : %v and email is: %v\n  ", adyashaNanda.Name, adyashaNanda.Email)

	adyashaNanda.GetStatus()
	adyashaNanda.NewMail()
	
	fmt.Printf("Name is : %v and email is: %v\n  ", adyashaNanda.Name, adyashaNanda.Email)
	//it will show u the old email ,cause the copy is created , we are not updating the email. Thts why pointer concept is there(remember)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is the user active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.gmail"
	fmt.Println("the new email is : ", u.Email)
}
