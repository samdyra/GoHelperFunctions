package main

import "fmt"

func main() {
	trySlice()
}

func trySlice() {
	a := []int{1,2,3,4,5,6} //init slice
	b := a[:] //all
	c := a[2:] //3rd to end
	d := a[:3] //first 3 element
	e := a[3:6] // 456
	f := make([]int, 3, 100) //len, capacity
 fmt.Println(a)
 fmt.Println(b)
 fmt.Println(c)
 fmt.Println(d)
 fmt.Println(e)
 fmt.Println(f)
}