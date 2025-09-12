package student

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/utils"
)

func QueryStudent(ctx context.Context, req *eva.QueryStudentRequest) (*eva.QueryStudentResponse, error) {

	pageSize, pageNum := utils.SetPageDefault(req.PageSize, req.PageNum)
	studentList, total, err := db.QueryStudentPage(ctx, db.DB, pageSize, pageNum, req.QueryStudentCondition)
	if err != nil {
		hlog.CtxErrorf(ctx, "QueryStudentPage db failed: %v", err)
		return nil, err
	}

	studentInfoList, err := WrappedStudentInfo(ctx, studentList)
	if err != nil {
		hlog.CtxErrorf(ctx, "WrappedStudentInfo db failed: %v", err)
		return nil, err
	}

	return &eva.QueryStudentResponse{
		StudentInfoList: studentInfoList,
		Total:           total,
		BaseResp:        handler.ConstructSuccessResp(),
	}, nil
}

func WrappedStudentInfo(ctx context.Context, studentList []*db.Student) ([]*eva.StudentInfo, error) {

	// get class list
	classIDs := utils.MapStructList(studentList, func(student *db.Student) int64 {
		return student.ClassID
	})

	classIDSet := utils.DistinctIntArray(classIDs)

	classList, err := db.FindClassListByIds(ctx, db.DB, classIDSet)
	if err != nil {
		hlog.CtxErrorf(ctx, "FindClassListByIds db failed: %v", err)
		return nil, err
	}

	// class ID -> class number
	classMap := utils.ToMap(classList,
		func(class *db.StudentClass) int64 { return class.ID },
		func(class *db.StudentClass) string { return class.ClassNumber },
	)

	studentInfoList := make([]*eva.StudentInfo, 0, len(classList))
	for _, student := range studentList {
		studentInfoList = append(studentInfoList, &eva.StudentInfo{
			ID:            &student.ID,
			StudentNumber: student.StudentNumber,
			StudentName:   student.StudentName,
			ClassNumber:   classMap[student.ClassID],
			Grade:         int8(student.Grade),
			Gender:        eva.Gender(student.Gender),
		})
	}

	return studentInfoList, nil
}
