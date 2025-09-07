package student_class

import (
	"context"
	"fmt"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
)

func EditStudentClass(ctx context.Context, req *eva.StudentClassEditRequest) (*eva.StudentClassEditResponse, error) {
	if err := CheckClassParam(req.ClassNumber); err != nil {
		return nil, err
	}

	if _, err := db.FindClassById(ctx, db.DB, req.GetID()); err != nil {
		return nil, err
	}

	updateClass := &db.StudentClass{
		ID:          req.ID,
		ClassNumber: req.ClassNumber,
	}

	if err := db.UpdateClass(ctx, db.DB, updateClass); err != nil {
		return nil, fmt.Errorf("update student class failed: %v", err)
	}

	return &eva.StudentClassEditResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
