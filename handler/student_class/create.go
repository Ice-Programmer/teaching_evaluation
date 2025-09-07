package student_class

import (
	"context"
	"fmt"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/utils"
)

func CreateStudentClass(ctx context.Context, req *eva.StudentClassCreateRequest) (*eva.StudentClassCreateResponse, error) {

	if err := checkParam(req.ClassNumber); err != nil {
		return nil, err
	}

	id := utils.GetId()

	if err := db.CreateStudentClass(ctx, db.DB, &db.StudentClass{
		ID:          id,
		ClassNumber: req.ClassNumber,
		CreateAt:    utils.GetNowSecs(),
		IsDelete:    0,
	}); err != nil {
		return nil, err
	}

	return &eva.StudentClassCreateResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func checkParam(classNumber string) error {
	if classNumber == "" {
		return fmt.Errorf("class number is required")
	}

	return nil
}
