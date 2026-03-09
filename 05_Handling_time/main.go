package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("learning time handling in go")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	//why 01-02-2026
	//cause its mentioned in the official documentation, to put the date in this format

	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	//also same for the day, cause its the syntax

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
	// same for the time as well

	fmt.Println("---------------------")

	//want to create a date by ownn
	createDate := time.Date(2026, time.March, 10, 12, 22, 12, 0, time.UTC)
	fmt.Println(createDate) // looks destructured, need a little bit formatting

	fmt.Println(createDate.Format("01-02-2006 Monday 15:04:05"))

}
