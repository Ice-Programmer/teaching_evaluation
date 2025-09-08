package student

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/consts"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/utils"
)

func CheckStudent(req *eva.CreateStudentRequest) error {
	if req.StudentNumber == "" {
		return fmt.Errorf("student number is required")
	}

	if req.ClassNumber == "" {
		return fmt.Errorf("class number is required")
	}

	if req.StudentNumber == "" {
		return fmt.Errorf("student number is required")
	}

	if err := utils.CheckGradeValue(req.Grade); err != nil {
		return err
	}

	if !utils.Contains(consts.MajorList, int8(req.Major)) {
		return fmt.Errorf("major number must in %d", consts.MajorList)
	}

	if !utils.Contains(consts.GenderList, int8(req.Gender)) {
		return fmt.Errorf("gender number must in %d", consts.GenderList)
	}

	return nil
}

func ValidateStudentExist(ctx context.Context, studentNumber string) error {
	record, err := db.FindStudentByNumber(ctx, db.DB, studentNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if record != nil {
		return fmt.Errorf("[%s] student number is exist", studentNumber)
	}

	return nil
}

func CheckBatchStudentParam(studentList []*eva.StudentInfo) error {
	if len(studentList) == 0 {
		return fmt.Errorf("studentList is empty")
	}

	seen := make(map[string]struct{}, len(studentList))

	for index, studentInfo := range studentList {
		if _, ok := seen[studentInfo.StudentNumber]; ok {
			return fmt.Errorf("student at position %d: duplicate student number [%s]", index+1, studentInfo.StudentNumber)
		}
		seen[studentInfo.StudentNumber] = struct{}{}

		if err := CheckStudent(&eva.CreateStudentRequest{
			StudentNumber: studentInfo.StudentNumber,
			StudentName:   studentInfo.StudentName,
			ClassNumber:   studentInfo.ClassNumber,
			Grade:         studentInfo.Grade,
			Major:         studentInfo.Major,
			Gender:        studentInfo.Gender,
		}); err != nil {
			return fmt.Errorf("student at position %d: %s", index+1, err.Error())
		}
	}

	return nil
}

func ValidateStudentListExist(ctx context.Context, studentList []*eva.StudentInfo) error {
	studentNumberList := utils.MapStructList(studentList, func(studentInfo *eva.StudentInfo) *string {
		return &studentInfo.StudentNumber
	})

	existStudentList, err := db.FindStudentListByNumberList(ctx, db.DB, studentNumberList)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if len(existStudentList) == 0 {
		return nil
	}

	existNumberList := utils.MapStructList(existStudentList, func(studentInfo *db.Student) string {
		return studentInfo.StudentNumber
	})

	return fmt.Errorf("%s student number is exist", existNumberList)
}
