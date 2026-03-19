package main

import "fmt"

func main() {
	fmt.Println("learning loops today")

	days := []string{"sunday", "monday", "wednesday", "friday", "saturday"}
	fmt.Println(days)

	//for loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// i - not only individual value but also index
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for i, day := range days {
		fmt.Printf("index is %v and value is %v\n", i, day)
	}

	//somthing with number
	randomValue := 1
	for randomValue < 10 {
		//goto
		if randomValue == 4 {
			goto adyasha
		}

		if randomValue == 2 {
			randomValue++
			continue
		}

		if randomValue == 7 {
			break
		}

		fmt.Println("value is: ", randomValue)
		randomValue++
	}

	//a label where goto jumps
	adyasha:
		fmt.Println("jumping here")

}
