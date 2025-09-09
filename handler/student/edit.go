package student

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/consts"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/service/student"
	"teaching_evaluation_backend/utils"
)

func EditStudent(ctx context.Context, req *eva.EditStudentRequest) (*eva.EditStudentResponse, error) {
	if err := checkParam(req); err != nil {
		return nil, err
	}

	studentInfo, err := db.FindStudentByID(ctx, db.DB, req.ID)
	if err != nil {
		return nil, err
	}

	if req.StudentNumber != studentInfo.StudentNumber {
		if err := student.ValidateStudentExist(ctx, req.StudentNumber); err != nil {
			return nil, err
		}
	}

	class, err := db.FindClassByNumber(ctx, db.DB, req.ClassNumber)
	if err != nil {
		hlog.CtxErrorf(ctx, "FindClassByNumber error: %s", err.Error())
		return nil, fmt.Errorf("class number not exist")
	}

	updateStudent := &db.Student{
		ID:            req.ID,
		StudentNumber: req.StudentNumber,
		Password:      req.StudentNumber,
		StudentName:   req.StudentName,
		Gender:        int8(req.Gender),
		ClassID:       class.ID,
		Major:         int8(req.Major),
		Grade:         int(req.Grade),
		Status:        int8(req.Status),
	}

	if err := db.UpdateStudent(ctx, db.DB, updateStudent); err != nil {
		hlog.CtxErrorf(ctx, "UpdateStudent error: %s", err.Error())
		return nil, err
	}

	return &eva.EditStudentResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func checkParam(req *eva.EditStudentRequest) error {
	if req.GetID() <= 0 {
		return fmt.Errorf("student id is required")
	}

	if !utils.Contains(consts.StatusList, int(req.Status)) {
		return fmt.Errorf("status must in %d", consts.StatusList)
	}

	return student.CheckStudent(&eva.CreateStudentRequest{
		StudentNumber: req.StudentNumber,
		StudentName:   req.StudentName,
		Gender:        req.Gender,
		ClassNumber:   req.ClassNumber,
		Major:         req.Major,
		Grade:         req.Grade,
	})
}
