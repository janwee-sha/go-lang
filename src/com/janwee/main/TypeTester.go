package main

import (
	"../common"
	"fmt"
)

func main() {
	s1,err := common.NewStudent("Janwee",map[string]float64{common.Chinese: 101,common.Math:75,common.English:89},
	23)
	if err != nil{
		fmt.Println("error occurred while creating student")
	}else{
		common.PrintScore(s1)
	}

}
