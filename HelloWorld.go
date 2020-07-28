package main

import "fmt"

func main() {
	var greeting string="Hello,%s!\n"
	var name string="Janwee"
	var score float32=78.5
	var rank int=23
	var passed bool=true
	fmt.Printf(greeting,name)
	fmt.Printf("score : %f\n",score)
	fmt.Printf("rank : %d\n",rank)
	fmt.Printf("passed : %t\n",passed)
}
