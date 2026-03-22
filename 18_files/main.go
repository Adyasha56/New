package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with files")

	content := "hi ! adyasha this side"

	file, err := os.Create("./myFilesInGO.txt")

	if err != nil {
		panic(err) //panic will shutdown the program once it
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err) //panic will shutdown the program once it
	}

	fmt.Println("length is : ", length)
	defer file.Close()

	readFile("./myFilesInGO.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err) //panic will shutdown the program once it
	}

	fmt.Println("text data inside the file is \n", string(databyte))
}
