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

func CheckStudent(ctx context.Context, req *eva.CreateStudentRequest) error {
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

	record, err := db.FindStudentByNumber(ctx, db.DB, req.StudentNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if record != nil {
		return fmt.Errorf("[%s] student number is exist", req.StudentNumber)
	}

	return nil
}
