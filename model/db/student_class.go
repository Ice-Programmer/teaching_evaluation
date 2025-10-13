package db

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/utils"
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

func FindClassListByIds(ctx context.Context, db *gorm.DB, ids []int64) ([]*StudentClass, error) {
	if db == nil {
		db = DB
	}

	var studentClass []*StudentClass
	err := db.Table(StudentClassTableName).WithContext(ctx).
		Where("id in (?)", ids).
		Find(&studentClass).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "FindClassListByIds db failed: %v", err)
		return nil, err
	}
	return studentClass, nil
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

func QueryClassPage(ctx context.Context, db *gorm.DB, pageSize int32, pageNum int32, condition *eva.QueryClassCondition) ([]*StudentClass, int64, error) {
	if db == nil {
		db = DB
	}

	query := buildClassCondition(db, condition).Table(StudentClassTableName).WithContext(ctx)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryClassPage db count failed: %v", err)
		return nil, 0, err
	}

	var studentClassList []*StudentClass
	offset := int((pageNum - 1) * pageSize)
	if err := query.Limit(int(pageSize)).Offset(offset).Find(&studentClassList).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryClassPage db failed: %v", err)
		return nil, 0, err
	}

	return studentClassList, total, nil
}

func buildClassCondition(db *gorm.DB, condition *eva.QueryClassCondition) *gorm.DB {
	db = db.Where("is_delete = 0")
	if condition == nil {
		return db
	}

	if condition.ID != nil {
		db = db.Where("id = ?", condition.ID)
	}

	if condition.ClassNumber != nil {
		db = db.Where("class_number like ?", utils.WrapLike(*condition.ClassNumber))
	}

	if condition.Ids != nil {
		db = db.Where("id in (?)", condition.Ids)
	}

	return db
}

func DeleteStudentClass(ctx context.Context, db *gorm.DB, id int64) error {
	if db == nil {
		db = DB
	}

	err := db.Table(StudentClassTableName).WithContext(ctx).
		Where("id = ?", id).
		Updates(utils.GenerateDeleteMap()).
		Error
	if err != nil {
		hlog.CtxErrorf(ctx, "BatchCreateListByNumber db failed: %v", err)
		return err
	}
	return nil
}
