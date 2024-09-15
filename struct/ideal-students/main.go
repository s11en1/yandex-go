package main

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (student Student) IsExcellentStudent() bool {
	return (float64(student.solvedProblems) * student.scoreForOneTask) >= student.passingScore
}
