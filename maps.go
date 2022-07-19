package main

import "fmt"

func main() {
	println("test")
}

func maps() {
	statePop := make(map[string]string) // make key value pair (objects) with string as both key and value types
	statePop = map[string]string{
		"city" : "bandung",
		"name" : "sam",
	}
	statePop["gf"] = "lol imagine"
	delete(statePop, "city")
	fmt.Println(statePop)
}
