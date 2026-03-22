package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("u 2nd")
	defer fmt.Println("u 1st")

	//defer follows last in first out concept
	fmt.Println("hello")
	myDefer()

}

func myDefer() {
	for i := 0; i < 4; i++ {
		fmt.Print(i)
	}
}

// output will be
// hello , 3210, u 1st , u 2nd, world
