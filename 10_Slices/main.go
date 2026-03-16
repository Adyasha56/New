package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("learning about slices today")

	var heroList = []string{"thor", "groot", "hulk"}
	fmt.Printf("the type of herolist is %T\n", heroList)

	heroList = append(heroList, "ironman", "black panther")
	fmt.Println(heroList)

	heroList = append(heroList[1:])
	fmt.Println(heroList) // 1st ele is no longer available

	heroList = append(heroList[1:3])
	fmt.Println(heroList) // 1st and (as its not inclusive) ele is no longer available

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 534
	highscores[3] = 834
	// highscores[4] = 1000  - give u err about range -like size is 4 , no more val can be added

	//but below check this line
	highscores = append(highscores, 764, 231, 769)
	//go behaves a little diff - append method will reallocate the memory so new val can get accomodated

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println("sorted slices are : ", highscores)

	//how to remove a val from a slice based on Index

	var courses = []string{"js", "go", "ts", "ruby", "nodejs"}
	fmt.Println("all courses are: ", courses)

	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}
