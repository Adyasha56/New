package main

import "fmt"

func main() {
	fmt.Println("learning about ifelse today")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "freq user"
	} else if loginCount > 10 {
		result = "normal user"
	} else {
		result = "excatly 10 count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	//special - used in web request
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")

	} else {
		fmt.Println("num is not less than 10")
	}

}
