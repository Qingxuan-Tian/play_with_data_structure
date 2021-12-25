package student

import "fmt"

type student struct {
	name string
	score int
}

func New(name string, score int)student{
	return student{
		name: name,
		score: score,
	}
}
func (s *student)string()string{
	return fmt.Sprintf("Student(name: %s, score: %d)", s.name, s.score)
}