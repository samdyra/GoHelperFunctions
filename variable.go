package main

import (
	"fmt"
	"strconv"
)

func main() {
	printMyName()
}

func printMyName() {
	var myAge int = 98 // this is the same as myAge := 98
	var herAge string
	herAge = strconv.Itoa(myAge)
	fmt.Println("my age is and my age type is")
	fmt.Printf("%T, %v\n", herAge, herAge)
}

