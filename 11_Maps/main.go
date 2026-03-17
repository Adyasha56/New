package main

import (
	"fmt"
)

func main() {
	fmt.Println("learning about maps today")

	languages := make(map[string]string)

	languages["js"] = "used more in web dev"
	languages["py"] = "used more in AI-ML"
	languages["c"] = "used more in embedded systems"

	fmt.Println("the value inside the maps are: ", languages)
	fmt.Println("the js inside the maps is: ", languages["js"])
	fmt.Println("the py inside the maps is: ", languages["py"])
	fmt.Println("the c inside the maps is: ", languages["c"])

	delete(languages, "c")
	fmt.Println(languages)

	//loops are pretty intersting in Golang
	for key, value := range languages {
		fmt.Printf("for key %v, value is : %v\n",key, value)
	}

}
