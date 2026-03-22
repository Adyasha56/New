package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://ankit31.me"

func main() {
	fmt.Println("web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf(" response ids of type %T\n", response)
	defer response.Body.Close() // its imp to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}
