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

func FindClassById(ctx context.Context, db *gorm.DB, id int64) (*StudentClass, error) {
	if db == nil {
		db = DB
	}
	var studentClass StudentClass
	err := db.Table(StudentClassTableName).WithContext(ctx).
		Where("id = ? and is_delete = 0", id).
		First(&studentClass).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "FindClassById db failed: %v", err)
		return nil, err
	}

	return &studentClass, nil
}

func UpdateClass(ctx context.Context, db *gorm.DB, studentClass *StudentClass) error {
	if db == nil {
		db = DB
	}

	err := db.Table(StudentClassTableName).WithContext(ctx).
		Where("id = ? and is_delete = 0", studentClass.ID).
		Updates(studentClass).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "UpdateClass db failed: %v", err)
		return err
	}
	return nil
}

func FindClassByNumber(ctx context.Context, db *gorm.DB, number string) (*StudentClass, error) {
	if db == nil {
		db = DB
	}
	var studentClass StudentClass
	err := db.Table(StudentClassTableName).WithContext(ctx).
		Where("class_number = ? and is_delete = 0", number).
		First(&studentClass).Error

	if err != nil {
		hlog.CtxErrorf(ctx, "FindClassByNumber db failed: %v", err)
		return nil, err
	}
	return &studentClass, nil
}

func FindClassListByNumberList(ctx context.Context, db *gorm.DB, numberList []string) ([]*StudentClass, error) {
	if db == nil {
		db = DB
	}

	var studentClassList []*StudentClass
	err := db.Table(StudentClassTableName).WithContext(ctx).
		Where("class_number in (?) and is_delete = 0", numberList).
		Find(&studentClassList).Error

	if err != nil {
		hlog.CtxErrorf(ctx, "FindClassListByNumberList db failed: %v", err)
		return nil, err
	}
	return studentClassList, nil
}

func BatchCreateListByNumber(ctx context.Context, db *gorm.DB, studentClassList []*StudentClass) error {
	if db == nil {
		db = DB
	}

	if err := db.Table(StudentClassTableName).WithContext(ctx).Create(studentClassList).Error; err != nil {
		hlog.CtxErrorf(ctx, "BatchCreateListByNumber db failed: %v", err)
		return err
	}
	return nil
}
