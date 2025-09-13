package login

import (
	"context"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/utils"
)

func GetCurrentUser(ctx context.Context) (*eva.GetCurrentUserResponse, error) {
	userInfo, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return &eva.GetCurrentUserResponse{
		UserInfo: userInfo,
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
