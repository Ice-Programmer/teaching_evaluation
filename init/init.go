package init

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"teaching_evaluation_backend/consts"
	"teaching_evaluation_backend/model/db"
	"teaching_evaluation_backend/utils"
)

func Init(ctx context.Context) error {
	if err := InitDBGorm(); err != nil {
		hlog.CtxErrorf(ctx, "InitDBGorm err: %v", err)
		panic(err)
	}
	hlog.CtxInfof(ctx, "InitDBGorm success")

	if err := utils.InitIdGeneratorClient(); err != nil {
		hlog.CtxErrorf(ctx, "InitIdGeneratorClient err: %v", err)
		panic(err)
	}

	return nil
}

func InitDBGorm() error {
	gormDB, err := newInit()
	if err != nil {
		return err
	}
	db.DB = gormDB
	return nil
}

func newInit() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		consts.DBUser, consts.DBPassword,
		consts.DBHost, consts.DBPort, consts.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
