package student_class

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/model/db"
)

func DeleteStudentClass(ctx context.Context, req *eva.DeleteStudentClassRequest) (*eva.DeleteStudentClassResponse, error) {

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		hlog.CtxErrorf(ctx, "strconv.ParseInt(%s): %v", req.ID, err)
		return nil, err
	}

	if err := db.DeleteStudentClass(ctx, db.DB, id); err != nil {
		hlog.CtxErrorf(ctx, "db.DeleteStudentClass(%d): %v", id, err)
		return nil, err
	}

	return &eva.DeleteStudentClassResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
