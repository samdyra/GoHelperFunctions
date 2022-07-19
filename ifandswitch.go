package main

import "fmt"

func main() {
	ifState()
}

func switchState() {
	i:= 2
	switch {
	case i < 5:
		fmt.Println("lol")
	default: fmt.Println("lolzzz")
	}
}

func ifState() {
	if true {
		fmt.Println("this is true")
	}
}