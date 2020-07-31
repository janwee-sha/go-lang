package main

import "../common"

func main() {
	var students1 []*common.Student=make([]*common.Student,1)
	s := common.NewStudent("Sherlock", map[string]float64{"English":90},1)
	students1 = append(students1, s)
	for i,_ := range students1 {
		common.PrintScore(students1[i])
	}

}