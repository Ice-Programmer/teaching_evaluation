package db

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

var (
	StudentClassTableName = "student_class"
)

type StudentClass struct {
	ID          int64  `gorm:"primary_key"`
	ClassNumber string `gorm:"class_number"`
	CreateAt    int64  `gorm:"create_at"`
	IsDelete    int64  `gorm:"is_delete"`
}

func (s StudentClass) TableName() string {
	return StudentClassTableName
}

func CreateStudentClass(ctx context.Context, db *gorm.DB, studentClass *StudentClass) error {
	if db == nil {
		db = DB
	}
	if err := db.Table(StudentClassTableName).WithContext(ctx).Create(&studentClass).Error; err != nil {
		hlog.CtxErrorf(ctx, "AddStudentClass db failed: %v", err)
		return err
	}

	return nil
}
