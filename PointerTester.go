package main

import "fmt"

/*
a pointer tester
 */
func main() {
	var i int =10
	var ptr *int=&i
	var pptr **int=&ptr

	fmt.Println(i)
	fmt.Println(*ptr)
	fmt.Println(ptr)
	fmt.Println(*pptr)
	fmt.Println(pptr)
}
