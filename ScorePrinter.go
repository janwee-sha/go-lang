package main

import "fmt"

var (
	greeting string = "Hello,%s!\n"
	name string = "Janwee" //5.因式分解关键字写法
)
func main() {
	var passed bool= true //1.var name = value方式
	var scores=[]float64{78.5,89,90}//2.根据值自行判断变量类型的方式
	courses := []string{"chinese","math","english"} //3.用:=赋值符号替代var name = value的方式
	averageScore,rank := 0.0,23//4.多变量声明
	fmt.Printf(greeting,name)
	for i,course := range courses {
		fmt.Printf("course : %s ,score : %f\n", course, scores[i])
		averageScore += scores[i]
	}
	averageScore/= float64(len(courses))
	fmt.Printf("average score : %f\n", averageScore)
	fmt.Printf("rank : %d\n",rank)
	fmt.Printf("passed : %t\n",passed)
}
