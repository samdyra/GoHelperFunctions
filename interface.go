package main

import "fmt"

func main() {
	var sam Person
	sam.Name	 = "sam"
	sayHello(sam)
}

type HasName interface {
	GetName() string
}

func (person Person) GetName() string {
	return person.Name
}
type Person struct {
	Name string
}

func sayHello(hasName HasName) {
	fmt.Println(hasName.GetName())
}




