package db

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/utils"
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

func StudentLogin(ctx context.Context, db *gorm.DB, userAccount, userPassword string) (*Student, error) {
	if db == nil {
		db = DB
	}

	var student *Student
	err := db.Table(StudentTableName).WithContext(ctx).
		Where("student_number = ? and password = ? and is_delete = 0", userAccount, userPassword).
		First(&student).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "StudentLogin db failed: %v", err)
		return nil, err
	}

	return student, nil
}

func QueryStudentPage(ctx context.Context, db *gorm.DB, pageSize, pageNum int32, condition *eva.QueryStudentCondition) ([]*Student, int64, error) {
	if db == nil {
		db = DB
	}

	query := buildCondition(db, condition).Table(StudentTableName).WithContext(ctx)

	// 统计总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryStudentPage db failed: %v", err)
		return nil, 0, err
	}

	offset := int((pageNum - 1) * pageSize)
	var studentList []*Student
	err := query.Where("is_delete = 0").
		Limit(int(pageSize)).Offset(offset).
		Find(&studentList).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "QueryStudentPage db failed: %v", err)
		return nil, 0, err
	}

	return studentList, total, nil
}

func buildCondition(db *gorm.DB, condition *eva.QueryStudentCondition) *gorm.DB {
	if condition == nil {
		return db
	}

	if condition.ID != nil {
		return db.Where("id = ?", condition.ID)
	}

	if condition.Major != nil {
		return db.Where("major = ?", condition.Major)
	}

	if condition.Name != nil {
		return db.Where("student_name like ?", utils.WrapLike(*condition.Name))
	}

	if condition.Number != nil {
		return db.Where("student_number = ?", condition.Number)
	}

	if condition.ClassId != nil {
		return db.Where("class_id = ?", condition.ClassId)
	}

	if condition.Grade != nil {
		return db.Where("grade = ?", condition.Grade)
	}

	return db
}
