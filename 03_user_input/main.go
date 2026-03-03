package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("learning how to take user input")
	//bufio/os is a lib or package
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name")

	//comma ok syntax || err ok syntax

	input, _ := reader.ReadString('\n') //why \n = when do u stop reading - at new line 
	//the _ represents the err part, as 
	fmt.Println("well got your name: ", input)
	fmt.Printf("the type of input is %T",input)

}
