package student

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/consts"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/service/student"
	"teaching_evaluation_backend/utils"
)

func CreateStudent(ctx context.Context, req *eva.CreateStudentRequest) (*eva.CreateStudentResponse, error) {
	if err := student.CheckStudent(req); err != nil {
		hlog.CtxErrorf(ctx, "checkStudent error: %s", err.Error())
		return nil, err
	}

	if err := student.ValidateStudentExist(ctx, req.StudentNumber); err != nil {
		return nil, err
	}

	studentClass, err := db.FindClassByNumber(ctx, db.DB, req.ClassNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("[%s] class not exist", req.ClassNumber)
		}
		return nil, err
	}

	studentInfo := &db.Student{
		ID:            utils.GetId(),
		StudentNumber: req.StudentNumber,
		Password:      req.StudentNumber,
		StudentName:   req.StudentName,
		Gender:        int8(req.Gender),
		ClassID:       studentClass.ID,
		Major:         int8(req.Major),
		Grade:         int(req.Grade),
		Status:        consts.NormalStatus,
		CreateAt:      utils.GetNowSecs(),
	}

	if err = db.AddStudent(ctx, db.DB, studentInfo); err != nil {
		return nil, err
	}

	return &eva.CreateStudentResponse{
		ID:       studentInfo.ID,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
