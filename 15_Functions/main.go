package main

import "fmt"

func main() {
	fmt.Println("welcome to func")
	wish()

	// func wishTwo(){
	// 	fmt.Println("can i write the func def here ? Ans -NO, u can only call here")
	// }
	wishTwo()

	result := adder(7, 19)
	fmt.Println("the result is : ", result)

	//when we dont know how many val can come in
	sum := adderPro(1, 4, 8, 9)
	fmt.Println("sum of pro adder is: ", sum)
}

func wish() {
	fmt.Println("hi adyasha! hope u land a good job by this year end")
}
func wishTwo() {
	fmt.Println("can i write the func def here(inside main) ? Ans -NO, u can only call here")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

//the int outside the bracket tells - which type of return content it is

func adderPro(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
