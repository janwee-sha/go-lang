package main

import (
	"fmt"
)

type Student struct {
	name string
	scoreMap map[string]float64
	rank int
}

func main() {
	var averageScore float64
	var passed bool = true
	s1 := Student{"Janwee",map[string]float64{"Chinese": 65.5,"Math":75,"English":89},23}
	fmt.Println(s1.scoreMap["Chinese"])
	for course := range s1.scoreMap {
		score := s1.scoreMap[course]
		averageScore+=score
		passed = passed && score>=60
	}
	averageScore/= float64(len(s1.scoreMap))

	fmt.Printf(greeting, s1.name)
	fmt.Printf("average score : %f\n", averageScore)
	fmt.Printf("rank : %d\n", s1.rank)
}
