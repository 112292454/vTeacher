package logic

import (
	"log"
	"vTeacher/dao/mysql"
	"vTeacher/entity"
	"vTeacher/pkg/jwt"
	"vTeacher/pkg/snowflake"
)

// SignUp 注册业务逻辑
func SignUp(p *entity.RegisterForm) (error error) {
	// 1、判断用户存不存在
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		// 数据库查询出错
		return err
	}

	// 2、生成UID
	userId, err := snowflake.GetID()
	if err != nil {
		return mysql.ErrorGenIDFailed
	}
	// 构造一个User实例
	u := entity.User{
		UserID:   userId,
		UserName: p.UserName,
		Password: p.Password,
		Email:    p.Email,
	}
	// 3、保存进数据库
	return mysql.InsertUser(u)
}

// Login 登录业务逻辑代码
func Login(p *entity.LoginForm) (user *entity.User, error error) {
	user = &entity.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	// return jwt.GenToken(user.UserID,user.UserName)
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}

func GetUser(uid int) (user *entity.User, error error) {
	user, err := mysql.GetUserByID(uint64(uid))
	if err != nil {
		return nil, err
	}
	// 生成JWT
	// return jwt.GenToken(user.UserID,user.UserName)
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}

// GetAllUsers 获取所有用户信息
func GetAllUsers() ([]*entity.User, error) {
	// 查询数据库获取用户信息
	users, err := mysql.QueryAllUsers()
	if err != nil {
		log.Printf("查询用户信息失败：%v\n", err)
		return nil, err
	}

	// 为每个用户生成JWT令牌
	for _, user := range users {
		accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
		if err != nil {
			log.Printf("生成JWT令牌失败：%v\n", err)
			return nil, err
		}
		user.AccessToken = accessToken
		user.RefreshToken = refreshToken
	}

	return users, nil
}
func GetUserByEmail(mail string) (user *entity.User, error error) {
	user, err := mysql.GetUserByEmail(mail)
	if err != nil {
		return nil, err
	}
	// 生成JWT
	// return jwt.GenToken(user.UserID,user.UserName)
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return user, nil
}

func UpdateUser(user *entity.User) (i int64, error error) {
	res, err := mysql.UpdateUser(user)
	if err != nil {
		return 0, err
	}
	// 生成JWT
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return res.RowsAffected()
}

// UpdateEmail id是用户id，email是将设置的新的邮箱
func UpdateEmail(id int, email string) (i int64, error error) {
	user, err := mysql.InternalGetUserByID(uint64(id))
	if err != nil {
		return 0, err
	}

	// 业务处理 —— 修改邮箱
	user.Email = email
	return UpdateUser(user)
}
