package login

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/consts"
	"teaching_evaluation_backend/handler"
	"teaching_evaluation_backend/middle"
	"teaching_evaluation_backend/model/db"
	"time"
)

func UserLogin(ctx context.Context, userAccount, userPassword string) (*eva.UserLoginResponse, error) {

	if err := checkLoginParam(userAccount, userPassword); err != nil {
		return nil, err
	}

	userInfo, err := findUserInfo(ctx, userAccount, userPassword)
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(3 * time.Hour)
	token, err := middle.GenerateToken(expirationTime, userInfo)
	if err != nil {
		return nil, err
	}

	return &eva.UserLoginResponse{
		BaseResp: handler.ConstructSuccessResp(),
		Token:    token,
		ExpireAt: expirationTime.Unix(),
		UserInfo: userInfo,
	}, nil
}

func checkLoginParam(userAccount, userPassword string) error {
	if userAccount == "" {
		return fmt.Errorf("userAccount is empty")
	}
	if userPassword == "" {
		return fmt.Errorf("userPassword is empty")
	}

	return nil
}

func findUserInfo(ctx context.Context, userAccount, userPassword string) (*eva.UserInfo, error) {
	// 1. check is student
	studentInfo, err := db.StudentLogin(ctx, db.DB, userAccount, userPassword)
	if err != nil {
		return nil, err
	}

	if studentInfo != nil {
		if studentInfo.Status == consts.BanStatus {
			hlog.CtxErrorf(ctx, "student banned: %v", studentInfo)
			return nil, fmt.Errorf("student banned: %v", studentInfo)
		}

		return &eva.UserInfo{
			ID:       studentInfo.ID,
			Name:     studentInfo.StudentName,
			Role:     eva.UserRole_Student,
			CreateAt: studentInfo.CreateAt,
		}, nil
	}

	// 2. check is admin
	admin, err := db.AdminLogin(ctx, db.DB, userAccount, userPassword)
	if err != nil {
		return nil, err
	}

	if admin == nil {
		return nil, fmt.Errorf("username or password error")
	}

	return &eva.UserInfo{
		ID:       admin.ID,
		Name:     admin.UserName,
		Role:     eva.UserRole_Admin,
		CreateAt: admin.CreateAt,
	}, nil
}
