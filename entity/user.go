package entity

import (
	"encoding/json"
	"errors"
)

// User 定义请求参数结构体
type User struct {
	UserID       uint64 `json:"userid" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	UserName     string `json:"username" db:"user_name"`
	NickName     string `json:"nickname" db:"nick_name"`
	Password     string `json:"password" db:"password"`
	Email        string `json:"email" db:"email"` // 邮箱
	IsAdmin      bool   `json:"is_admin" db:"is_admin"`
	Level        int    `json:"level" db:"level"`
	AccessToken  string
	RefreshToken string
}

// UnmarshalJSON 为User类型实现自定义的UnmarshalJSON方法
func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"user_name"`
		Password string `json:"password" db:"password"`
		Email    string `json:"email" db:"email"` // 邮箱
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		u.UserName = required.UserName
		u.Password = required.Password
		u.Email = required.Email
	}
	return
}

// RegisterForm 注册请求参数
type RegisterForm struct {
	UserName string `json:"username" binding:"required"` // 用户名
	Email    string `json:"email" binding:"required"`    // 邮箱
	Password string `json:"password" binding:"required"` // 密码
}

// LoginForm 登录请求参数
type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Request
type LoginRequest struct {
	// 用户密码
	Password string `json:"password" binding:"required"`
	// 用户名
	Username string `json:"username" binding:"required"`
}

// UnmarshalJSON 为RegisterForm类型实现自定义的UnmarshalJSON方法
func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username"`
		Email    string `json:"email"`    // 邮箱
		Password string `json:"password"` // 密码
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if len(required.Email) == 0 {
		err = errors.New("缺少必填字段email")
	} else {
		r.UserName = required.UserName
		r.Email = required.Email
		r.Password = required.Password
	}
	return
}

// VoteDataForm 投票数据
type VoteDataForm struct {
	// UserID int 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
}

// UnmarshalJSON 为VoteDataForm类型实现自定义的UnmarshalJSON方法
func (v *VoteDataForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string `json:"post_id"`
		Direction int8   `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}
