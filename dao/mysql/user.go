package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"vTeacher/entity"
	"vTeacher/pkg/snowflake"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "nginx.show"

// encryptPassword 对密码进行加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(user_name string) (error error) {
	sqlStr := `select count(user_id) from user where user_name = ?`
	var count int
	if err := db.Get(&count, sqlStr, user_name); err != nil {
		return err
	}
	if count > 0 {
		return errors.New(ErrorUserExit)
	}
	return
}

// InsertUser 注册业务-向数据库中插入一条新的用户
func InsertUser(user entity.User) (error error) {
	// 对密码进行加密
	user.Password = encryptPassword([]byte(user.Password))
	// 执行SQL语句入库
	sqlstr := `insert into user(user_name,password,email) values(?,?,?)`
	_, err := db.Exec(sqlstr, user.UserName, user.Password, user.Email)
	return err
}

func Register(user *entity.User) (err error) {
	sqlStr := "select count(user_id) from user where user_name = ?"
	var count int64
	err = db.Get(&count, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		// 用户已存在
		return errors.New(ErrorUserExit)
	}
	// 生成user_id
	userID, err := snowflake.GetID()
	if err != nil {
		return ErrorGenIDFailed
	}
	// 生成加密密码
	password := encryptPassword([]byte(user.Password))
	// 把用户插入数据库
	sqlStr = "insert into user(user_id, user_name, password,email) values (?,?,?,?)"
	_, err = db.Exec(sqlStr, userID, user.UserName, password, user.Email)
	return
}

// Login 登录业务
func Login(user *entity.User) (err error) {
	originPassword := user.Password // 记录一下原始密码(用户登录的密码)
	sqlStr := "select user_id, user_name, password from user where user_name = ?"
	err = db.Get(user, sqlStr, user.UserName)
	// 查询数据库出错
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	// 用户不存在
	if err == sql.ErrNoRows {
		return errors.New(ErrorUserNotExit)
	}
	// 生成加密密码与查询到的密码比较
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return errors.New(ErrorPasswordWrong)
	}
	return nil
}

// GetUserByID 根据ID查询用户信息
func GetUserByID(id uint64) (user *entity.User, err error) {
	user = new(entity.User)
	sqlStr := `select user_id, user_name,email,is_admin,level,nick_name from user where user_id = ?`
	err = db.Get(user, sqlStr, id)
	return
}

// GetUserByID 根据ID查询用户信息
func GetUserByEmail(mail string) (user *entity.User, err error) {
	user = new(entity.User)
	sqlStr := `select user_id, user_name,email from user where email = ?`
	err = db.Get(user, sqlStr, mail)
	return
}

// queryAllUsersFromDB 从数据库查询所有用户信息
func QueryAllUsers() ([]*entity.User, error) {
	var users []*entity.User
	sqlStr := `select user_id, user_name, email, is_admin, level, nick_name from user`
	err := db.Select(&users, sqlStr)
	if err != nil {
		log.Printf("从数据库查询用户信息失败：%v\n", err)
		return nil, err
	}
	return users, nil
}

// 获取用户所有信息
func InternalGetUserByID(id uint64) (user *entity.User, err error) {
	user = new(entity.User)
	sqlStr := `select user_id, user_name,email,password,avatar,is_admin,level,nick_name from user where user_id = ?`
	err = db.Get(user, sqlStr, id)
	return
}

func GetSceneByName() {

}

func UpdateUser(user *entity.User) (res sql.Result, err error) {
	sqlStr := `UPDATE vTeacher.user
    SET avatar = :avatar, email = :email, is_admin = :is_admin, level = :level, nick_name = :nick_name, password = :password, user_name = :user_name
WHERE user_id=:user_id;`
	res, err = db.NamedExec(sqlStr, user)
	if err != nil {
		log.Printf("更新用户信息失败：%v\n", err)
		return nil, err
	}
	return res, nil
}
