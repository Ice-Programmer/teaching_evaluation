package utils

import "fmt"

const (
	MaxGrade = 8
	MinGrade = 1
)

func CheckGradeValue(grade int8) error {
	if grade < MinGrade || grade > MaxGrade {
		return fmt.Errorf("grade must be between %v and %v", MinGrade, MaxGrade)
	}
	return nil
}
