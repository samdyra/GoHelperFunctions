package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	defering()
}

func defering() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //run setelah defering() selesai
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
