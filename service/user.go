package service

import (
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
)

// Logic 登录函数通过前端传回的username和password对比和数据库里的是否相同
func Logic(username string, password string) (string, error) {
	// 从数据库里获取user
	user, err := dao.GetUserByUsername(username)
	// 如果返回的err不是nil说明没查到，及用户不名不存在
	if err != nil {
		return "", e.NewError(e.ERROR_USER_NOT_EXIST)
	}
	// 如果成功获取了user,那么对比数据库里的user的密码和前端传回登录者输入的密码是否相同
	// 如果不同返回密码错误的错误码
	if user.Password != password {
		return "", e.NewError(e.ERROR_PASSWORD_WRONG)
	}
	// 如果相同，就调用jwt生成token
	token, err := util.GenerateToken(user.UserId, username)
	// 生成失败返回状态
	if err != nil {
		return "", e.NewError(e.ERROR_TOKEN_WRONG)
	}
	// 成功返回token
	return token, nil
}

// Signup 注册函数username不重复的时候添加用户到数据库
func Signup(username string, password string, email string) error {
	// 判断username是否已经存在
	// 1.存在就返回用户名已经存在
	_, err := dao.GetUserByUsername(username)
	// err为nil说明查到了user应返回用户名已经存在
	if err == nil {
		return e.NewError(e.ERROR_USERNAME_EXIST)
	}
	// 2.用户名不存在就将用户信息存入数据库并返回注册成功
	// 先创建user,将api层传来的数据赋给user
	var user model.User
	user.Username = username
	user.Password = password
	user.Email = email
	// 使用dao层函数添加user
	err = dao.CreateUser(&user)
	// 如果返回err说明添加失败
	if err != nil {
		return err
	}
	// 添加成功返回nil
	return nil
}

// UpdateUser 修改用户信息,如果存在就修改信息，不存在就返回
func UpdateUser(username string, password string, email string) error {
	// 先判断用户名是否存在
	// 查找用户如果存在返回user,nil;  不存在返回nil和err
	user, err := dao.GetUserByUsername(username)
	// 用户不存在
	if err != nil {
		// 返回用户不存在
		return e.NewError(e.ERROR_USER_NOT_EXIST)
	}
	// 用户存在就修改用户信息
	user.Password = password
	user.Email = email
	// 将修改后的用户信息传给dao层覆盖修改
	err = dao.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
