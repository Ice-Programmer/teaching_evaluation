package student_class

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/utils"
)

func BatchCreateStudentClass(ctx context.Context, req *eva.BatchCreateStudentRequest) (*eva.BatchCreateStudentResponse, error) {
	if err := CheckBatchClassList(ctx, req.ClassNumberList); err != nil {
		hlog.CtxErrorf(ctx, "CheckBatchClassList Precheck error: %s", err.Error())
		return nil, err
	}

	numberClassList := make([]*db.StudentClass, 0, len(req.ClassNumberList))
	for _, number := range req.ClassNumberList {
		numberClassList = append(numberClassList, &db.StudentClass{
			ID:          utils.GetId(),
			ClassNumber: number,
			CreateAt:    utils.GetNowSecs(),
			IsDelete:    0,
		})
	}

	if err := db.BatchCreateListByNumber(ctx, db.DB, numberClassList); err != nil {
		hlog.CtxErrorf(ctx, "BatchCreateByNumber Precheck error: %s", err.Error())
		return nil, err
	}

	return &eva.BatchCreateStudentResponse{
		Num:      int32(len(numberClassList)),
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func CheckBatchClassList(ctx context.Context, classNumberList []string) error {
	if len(classNumberList) == 0 {
		return fmt.Errorf("classNumberList is empty")
	}

	seen := make(map[string]struct{}, len(classNumberList))
	for _, classNumber := range classNumberList {
		if _, exists := seen[classNumber]; exists {
			return fmt.Errorf("classNumberList contains duplicate value: %s", classNumber)
		}
		seen[classNumber] = struct{}{}

		// 参数校验
		if err := CheckClassParam(classNumber); err != nil {
			return fmt.Errorf("classNumber is invalid: %s", err.Error())
		}
	}

	studentClasses, err := db.FindClassListByNumberList(ctx, db.DB, classNumberList)
	if err != nil {
		return err
	}

	if len(studentClasses) != 0 {
		duplicateNumberList := utils.MapStructList(studentClasses, func(c *db.StudentClass) string {
			return c.ClassNumber
		})
		return fmt.Errorf("class %s has already been created", duplicateNumberList)
	}

	return nil
}
