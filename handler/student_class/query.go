package student_class

import (
	"context"
	"strconv"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/utils"
)

func QueryStudentClass(ctx context.Context, req *eva.QueryStudentClassRequest) (*eva.QueryStudentClassResponse, error) {

	pageSize, pageNum := utils.SetPageDefault(req.PageSize, req.PageNum)

	classList, total, err := db.QueryClassPage(ctx, db.DB, pageSize, pageNum, req.Condition)
	if err != nil {
		return nil, err
	}

	return &eva.QueryStudentClassResponse{
		Total:     total,
		ClassList: WrappedClassInfo(classList),
		BaseResp:  handler.ConstructSuccessResp(),
	}, nil
}

func WrappedClassInfo(classList []*db.StudentClass) []*eva.ClassInfo {
	classInfoList := make([]*eva.ClassInfo, 0, len(classList))
	for _, class := range classList {
		classInfoList = append(classInfoList, &eva.ClassInfo{
			ID:          strconv.FormatInt(class.ID, 10),
			ClassNumber: class.ClassNumber,
			CreateAt:    class.CreateAt,
		})
	}
	return classInfoList
}
