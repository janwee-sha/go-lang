package common

import "fmt"

func PrintGreeting(name string) {
	fmt.Printf("hello,%s\n",name)
}

func PrintPassed(passed bool) {
	if passed {
		fmt.Println("congratulations,you passed the exam!")
	} else {
		fmt.Println("sorry,you failed to pass the exam!")
	}
}

func PrintAverageScore(averageScore float64) {
	fmt.Printf("average score : %f\n", averageScore)
}

func PrintRank(rank int) {
	fmt.Printf("rank : %d\n", rank)
}

func PrintScore(s *Student){
	PrintGreeting(s.GetName())
	PrintPassed(s.passed)
	for course := range s.scoreMap {
		score := s.scoreMap[course]
		fmt.Printf("course : %s ,score : %f\n",course,score)
	}
	PrintAverageScore(s.averageScore)
	PrintRank(s.rank)
}
