package main

import "fmt"

func main() {
	greetFunc := greeter{
		name: "Sam",
		call: "hi",
	}
	greetFunc.tryMethodsinFunction()
}

type greeter struct {
	name string
	call string
}


func (greetFunc greeter) tryMethodsinFunction() {
	fmt.Println(greetFunc.call, greetFunc.name)

}
