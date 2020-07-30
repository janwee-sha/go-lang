package common

type Student struct {
	name string
	scoreMap map[string]float64
	rank int
	totalScore float64
	averageScore float64
	passed bool

}

func (s *Student) GetName() string {
	return s.name
}

func NewStudent (name string,scoreMap map[string]float64,rank int) *Student {
	var totalScore float64
	var averageScore float64
	var passed bool=true
	for course := range scoreMap {
		score := scoreMap[course]
		totalScore+=score
		passed = passed && score>=60
	}
	averageScore=totalScore/float64(len(scoreMap))
	return &Student{
		name: name,
		scoreMap: scoreMap,
		rank: rank,
		totalScore: totalScore,
		averageScore: averageScore,
		passed: passed,
	}
}
