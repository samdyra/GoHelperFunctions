package main

import (
	"fmt"
	"log"
)

func main() {
	panicing()
}


func panicing() {
	fmt.Println(("start"))
	defer func() {
		if err := recover(); err != nil { //recover function return nil if app tidak panic
			log.Println("Error:", err)
		}
	}()
	panic("error") //error sintesis
	fmt.Println("end") //start setelah defer
}
