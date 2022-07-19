package main

import "fmt"

func main() {
	looping()
}


func looping() {
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} 
		fmt.Println("hi")
	}
}

func mapping() {
	s := []int{1, 2, 3}
	for k, v := range s { //mapped array s, with k as key and v as value
		fmt.Println(k,v)
	}
}