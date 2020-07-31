package main

import (
	"../common"
	"fmt"
)

func main() {
	s1,err := common.NewStudent("Sherlock",map[string]float64{common.Chinese: 99,common.Math:92,common.English:95},
	1)
	handleError(err, s1)

	s2,err := common.NewStudent("John",map[string]float64{common.Chinese: 101,common.Math:75,common.English:89},
		23)
	handleError(err, s2)

}

func handleError(err error, s2 *common.Student) {
	if err != nil {
		fmt.Println("error occurred while creating student")
	} else {
		common.PrintScore(s2)
	}
}
