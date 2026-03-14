package main

import "fmt"

func main() {
	fmt.Println("learning about pointers")

	//creating a pointer
	var ptr *int
	fmt.Println("value of pointer is ",ptr) //<nil>

	//referencing a pointer - where & come in 
	myNumber := 23
	var ptr1 = &myNumber

	fmt.Println("value of  pointer is ",ptr1)
	fmt.Println("the value of actual pointer is: ",*ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("new value is ",myNumber)
}
