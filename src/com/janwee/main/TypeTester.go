package main

import (
	"../common"
)

func main() {
	s1 := common.NewStudent("Janwee",map[string]float64{"Chinese": 65.5,"Math":75,"English":89},23)
	common.PrintScore(s1)
}
