package main

import (
	"../common"
)

func main() {
	s1 := common.NewStudent("Janwee",map[string]float64{common.Chinese: 65.5,common.Math:75,common.English:89},
	23)
	common.PrintScore(s1)
}
