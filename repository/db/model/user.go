package model

// User 定义用户结构体
type User struct {
	UserId   int64  `gorm:"column:user_id;primaryKey;autoIncrement;comment:'用户id'"`
	Username string `gorm:"column:username;comment:'用户名'"`
	Password string `gorm:"column:password;comment:'用户密码'"`
	Email    string `gorm:"column:email;comment:'用户邮箱'"`
}
