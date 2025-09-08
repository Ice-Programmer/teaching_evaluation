package student

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strings"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/consts"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/service/student"
	"teaching_evaluation_backend/utils"
)

func BatchCreateStudent(ctx context.Context, req *eva.BatchCreateStudentRequest) (*eva.BatchCreateStudentResponse, error) {
	studentInfoList := req.GetStudentList()
	if err := student.CheckBatchStudentParam(studentInfoList); err != nil {
		hlog.CtxErrorf(ctx, "checkBatchStudentParam error: %s", err.Error())
		return nil, err
	}

	if err := student.ValidateStudentListExist(ctx, studentInfoList); err != nil {
		hlog.CtxErrorf(ctx, "validateStudentListExist error: %s", err.Error())
		return nil, err
	}

	// classNumber -> classId
	classMap, err := GetBatchStudentClassMap(ctx, req)
	if err != nil {
		return nil, err
	}

	studentList := make([]*db.Student, 0, len(studentInfoList))
	for _, studentInfo := range studentInfoList {
		studentList = append(studentList, &db.Student{
			ID:            utils.GetId(),
			StudentNumber: studentInfo.StudentNumber,
			Password:      studentInfo.StudentNumber,
			StudentName:   studentInfo.StudentName,
			Gender:        int8(studentInfo.Gender),
			ClassID:       classMap[studentInfo.ClassNumber],
			Major:         int8(studentInfo.Major),
			Grade:         int(studentInfo.Grade),
			Status:        consts.NormalStatus,
			CreateAt:      utils.GetNowSecs(),
		})
	}

	if err = db.BatchCreateStudents(ctx, db.DB, studentList); err != nil {
		return nil, err
	}

	return &eva.BatchCreateStudentResponse{
		Num:      int32(len(studentList)),
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func GetBatchStudentClassMap(ctx context.Context, req *eva.BatchCreateStudentRequest) (map[string]int64, error) {
	studentList := req.GetStudentList()

	classNumberList := utils.MapStructList(studentList, func(studentInfo *eva.StudentInfo) string {
		return studentInfo.ClassNumber
	})

	classNumberSet := utils.DistinctStringArray(classNumberList)

	classList, err := db.FindClassListByNumberList(ctx, db.DB, classNumberSet)
	if err != nil {
		hlog.CtxErrorf(ctx, "FindClassListByNumberList error: %s", err.Error())
		return nil, err
	}

	if len(classList) < len(classNumberSet) {
		recordList := utils.MapStructList(classList, func(class *db.StudentClass) string {
			return class.ClassNumber
		})

		diff := utils.Diff(classNumberList, recordList)
		return nil, fmt.Errorf("class [%s] not exist, please add classes info first", strings.Join(diff, ","))
	}

	classMap := utils.ToMap(classList,
		func(class *db.StudentClass) string { return class.ClassNumber },
		func(class *db.StudentClass) int64 { return class.ID },
	)

	return classMap, nil
}
