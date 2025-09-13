package utils

import (
	"context"
	"fmt"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
)

const (
	UserInfoKey = "UserInfoKey"
)

func SetCurrentUserInfo(ctx context.Context, userInfo interface{}) context.Context {
	return context.WithValue(ctx, UserInfoKey, userInfo)
}

func GetUserInfoFromContext(ctx context.Context) (*eva.UserInfo, error) {
	val := ContextGetKeyValue(ctx, UserInfoKey)
	if user, ok := val.(eva.UserInfo); ok {
		return &user, nil
	}
	return nil, fmt.Errorf("user not found in context")
}
