package db

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

var (
	StudentTableName = "student"
)

type Student struct {
	ID            int64  `gorm:"primary_key;column:id"`
	StudentNumber string `gorm:"column:student_number"`
	Password      string `gorm:"column:password"`
	StudentName   string `gorm:"column:student_name"`
	Gender        int8   `gorm:"column:gender"` // 0 - 女, 1 - 男
	ClassID       int64  `gorm:"column:class_id"`
	Major         int8   `gorm:"column:major"` // 0 - 计算机, 1 - 自动化
	Grade         int    `gorm:"column:grade"`
	Status        int8   `gorm:"column:status"` // 0 - 正常使用, 1 - 拒绝访问
	CreateAt      int64  `gorm:"column:create_at"`
	IsDelete      int8   `gorm:"column:is_delete"`
}

func (s Student) TableName() string {
	return StudentTableName
}

func AddStudent(ctx context.Context, db *gorm.DB, student *Student) error {
	if db == nil {
		db = DB
	}

	if err := db.Table(StudentTableName).WithContext(ctx).Create(student).Error; err != nil {
		hlog.CtxErrorf(ctx, "AddStudent db failed: %v", err)
		return err
	}

	return nil
}

func FindStudentByNumber(ctx context.Context, db *gorm.DB, studentNumber string) (*Student, error) {
	if db == nil {
		db = DB
	}

	var student Student
	err := db.Table(StudentTableName).
		Where("student_number = ? and is_delete = 0", studentNumber).
		First(&student).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "FindStudentByNumber db failed: %v", err)
		return nil, err
	}
	return &student, err
}

func FindStudentListByNumberList(ctx context.Context, db *gorm.DB, studentNumberList []*string) ([]*Student, error) {
	if db == nil {
		db = DB
	}

	var studentList []*Student
	err := db.Table(StudentTableName).WithContext(ctx).
		Where("student_number in (?) and is_delete = 0", studentNumberList).
		Find(&studentList).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "FindStudentListByNumber db failed: %v", err)
		return nil, err
	}

	return studentList, nil
}

func BatchCreateStudents(ctx context.Context, db *gorm.DB, students []*Student) error {
	if db == nil {
		db = DB
	}

	if err := db.Table(StudentTableName).WithContext(ctx).Create(students).Error; err != nil {
		hlog.CtxErrorf(ctx, "AddStudent db failed: %v", err)
		return err
	}

	return nil
}

func FindStudentByID(ctx context.Context, db *gorm.DB, id int64) (*Student, error) {
	if db == nil {
		db = DB
	}

	var student *Student
	err := db.Table(StudentTableName).
		Where("id = ? and is_delete = 0", id).First(&student).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "FindStudentByID db failed: %v", err)
		return nil, err
	}

	return student, nil
}

func UpdateStudent(ctx context.Context, db *gorm.DB, student *Student) error {
	if db == nil {
		db = DB
	}

	err := db.Table(StudentTableName).WithContext(ctx).
		Where("id = ? and is_delete = 0", student.ID).
		Updates(student).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "UpdateStudent db failed: %v", err)
		return err
	}
	return nil
}
