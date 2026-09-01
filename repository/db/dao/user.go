package dao

import (
	"gin-mall/pkg/e"
	"gin-mall/repository/db"
	"gin-mall/repository/db/model"
)

// GetUserByUsername 通过用户名获取用户
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	// 从数据库中按username查询到第一个user赋值给&user
	err := db.DB.Where("username = ?", username).First(&user).Error
	// 如果err不是nil,说明没有查找到user,返回err
	if err != nil {
		return nil, err
	}
	// err是nil说明查到了user返回&user
	return &user, nil
}

// CreateUser 在数据库里创建新用户
func CreateUser(user *model.User) error {
	// Gorm语句创建user
	err := db.DB.Create(user).Error
	// 创建失败返回err
	if err != nil {
		return err
	}
	// 成功返回nil
	return nil
}

// UpdateUser 修改用户信息
func UpdateUser(user *model.User) error {
	// 覆盖修改user
	result := db.DB.Save(user)
	// 如果是系统错误就返回错误信息
	if result.Error != nil {
		return result.Error
	}
	// 如果修改了0行数据用户id有错，返回用户不存在的自定义错误信息
	if result.RowsAffected == 0 {
		return e.NewError(e.ERROR_USER_NOT_EXIST)
	}
	// 修改成功
	return nil
}
