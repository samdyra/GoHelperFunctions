package main

import "fmt"

func main() {
	tryArray()
}

func tryArray() {
	var num [5]string //array of string length 5
	num[1] = "Sam"
	fmt.Printf("my grade: %v, %T\n", num, num )
	fmt.Printf("%v\n", len(num))

}