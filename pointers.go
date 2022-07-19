package main

import "fmt"

func main() {
	tryPointer()
}


func tryPointer() {
	var a int = 42 
	var b *int = &a //b is pointing integer to address a 
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a,*b) //asterix untuk derefrence (address ke value)
 *b = 14
	fmt.Println(a, *b)
}
