package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("learning conversion today")
	fmt.Println("rate my learning rate btn 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating", input)
	// fmt.Printf("the type of input is %T", input)

	// numRating,err := strconv.ParseFloat(input,64)
	// if err != nil {
	// 	fmt.Println(err)
	// }else {
	// 	fmt.Println("added 1 to your rating:: ", numRating+1)
	// }

	//this will give err = stringstrconv.ParseFloat: parsing "2\r\n": invalid syntax

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating:: ", numRating+1)
	}

}
