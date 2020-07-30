package main
/*
a score printer
 */
import (
	"fmt"
)

const (
	greeting = "hello,%s!\n" //5.因式分解关键字写法
)
var(
	name = "Janwee"
)
func main() {
	var passed bool= true //1.var name = value方式
	var scores=[]float64{78.5,89,90}//2.根据值自行判断变量类型的方式
	courses := []string{"Chinese","Math","English"} //3.用:=赋值符号替代var name = value的方式
	averageScore,rank := 0.0,23//4.多变量声明方式

	fmt.Printf(greeting,name)
	printPassed(passed)
	for i, course := range courses {
		fmt.Printf(getFormatAndCourseAndScore(scores,course,i))
		averageScore += scores[i]
		passed = passed && scores[i]>=60
	}

	averageScore = calcAverageScore(averageScore, len(courses))

	fmt.Printf("average score : %f\n", averageScore)
	fmt.Printf("rank : %d\n", rank)
}

func calcAverageScore(score float64, count int) float64 {
	score /= float64(count)
	return score
}

func getFormatAndCourseAndScore(scores []float64,course string,index int) (string,string,float64) {
	return "course : %s ,score : %f\n", course, scores[index];
}

func printPassed(passed bool) {
	if passed {
		fmt.Println("congratulations,you passed the exam!")
	} else {
		fmt.Println("sorry,you failed to pass the exam!")
	}
}
