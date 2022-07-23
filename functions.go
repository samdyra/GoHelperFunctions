package main

import "fmt"

func main() {
	tryFunctions("hello", 2)
	s,error := sum()
	if error != nil {
		fmt.Println(error)
		return 
	} 
	fmt.Println(s)
}

func tryFunctions(msg string, idx int) {
	fmt.Println(msg, idx)
}

func sum(values ...int) (int, error) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	if len(values) == 0 {
		return 0, fmt.Errorf("array nil")
	}
	return result, nil
}