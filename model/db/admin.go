package db

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

var (
	AdminTableName = "admin"
)

type Admin struct {
	ID       int64  `gorm:"primary_key;column:id"`
	UserName string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	CreateAt int64  `gorm:"column:create_at"`
	IsDelete bool   `gorm:"column:is_delete"`
}

func (Admin) TableName() string {
	return AdminTableName
}

func AdminLogin(ctx context.Context, db *gorm.DB, account string, password string) (*Admin, error) {
	if db == nil {
		db = DB
	}

	var admin *Admin
	err := db.Where("username = ? and password = ? and is_delete = 0", account, password).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "AdminLogin err: %v", err)
		return nil, err
	}
	return admin, err
}
