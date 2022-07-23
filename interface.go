package main

import "fmt"

func main() {
	var sam Person
	sam.Name	 = "sam"
	sayHello(sam)

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
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



type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
