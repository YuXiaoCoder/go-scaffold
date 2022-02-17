package rds

import (
	"errors"
	"go-scaffold/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorUserExist = errors.New("用户已存在")
)

type UserDB struct{}

func GetUserDB() *UserDB {
	return &UserDB{}
}

// CheckExistByEmail 通过邮箱判断用户是否存在
func (*UserDB) CheckExistByEmail(email string) (err error) {
	var count int64
	if err = db.Model(models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// CheckExistByID 通过ID判断用户是否存在
func (*UserDB) CheckExistByID(id int64) (err error) {
	var count int64
	if err = db.Model(models.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// EncryptPassword 对密码进行加密
func (*UserDB) EncryptPassword(password string) (encodePWD string, err error) {
	// 对密码进行加密
	encodePWDByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encodePWDByte), nil
}

// ComparePassword 验证密码的正确性
func (*UserDB) ComparePassword(encodePWD, password string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encodePWD), []byte(password))
}

// Create 创建用户
func (*UserDB) Create(user *models.User) (err error) {
	return db.Create(user).Error
}

// GetByConditions 获取用户列表
func (*UserDB) GetByConditions(params *models.QueryUserInfoListNull) (users []*models.User, count int64, err error) {
	// 初始化
	users = make([]*models.User, 0)

	// 事务
	var tx = db

	// 链式编程
	if params.Limit.Valid {
		tx = tx.Limit(int(params.Limit.Int64))
	}
	if params.Offset.Valid {
		tx = tx.Offset(int(params.Offset.Int64))
	}

	// 执行查询
	err = tx.Find(&users).Limit(-1).Offset(-1).Count(&count).Error
	return users, count, err
}

// GetByID 通过ID查询用户
func (*UserDB) GetByID(id int64) (user *models.User, err error) {
	user = new(models.User)
	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByEmail 通过邮箱查询用户
func (*UserDB) GetByEmail(email string) (user *models.User, err error) {
	user = new(models.User)
	err = db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateByMap 更新用户
func (*UserDB) UpdateByMap(id int64, values map[string]interface{}) error {
	return db.Model(models.User{}).Where("id = ?", id).Updates(values).Error
}

// DeleteByID 删除用户
func (*UserDB) DeleteByID(id int64) error {
	return db.Delete(&models.User{}, id).Error
}
