package main

import "fmt"

func main() {
	fmt.Println("learning about array today")

	var heroList [4]string //u have to add a val explicitly - compulsion to it
	heroList[0] = "IronMan"
	heroList[1] = "Thor"
	// heroList[2] = "hulk"
	heroList[3] = "Groot"

	fmt.Println("the heros are : ", heroList)
	fmt.Println("the length of the arr is : ", len(heroList))

	var travelList = [3]string{"Nainital", "Mussorie", "Flowervalley"}
	fmt.Println("my travel plan is: ", travelList)

}
