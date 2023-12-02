package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"log"
	"strconv"
	"vTeacher/config"
	"vTeacher/dao/mysql"
	"vTeacher/entity"
	"vTeacher/logic"
)

// SignUpHandler 注册业务
func SignUpHandler(c *gin.Context) {
	// 1.获取请求参数
	var fo *entity.RegisterForm
	// 2.校验数据有效性
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			entity.ResponseError(c, entity.CodeInvalidParams) // user 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		entity.ResponseErrorWithMsg(c, entity.CodeInvalidParams, config.RemoveTopStruct(errs.Translate(config.Trans)))
		return // 翻译错误
	}
	fmt.Printf("fo: %v\n", fo)
	// 3.业务处理 —— 注册用户
	if err := logic.SignUp(fo); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {
			entity.ResponseError(c, entity.CodeUserExist)
			return
		}
		entity.ResponseError(c, entity.CodeServerBusy)
		return
	}
	user, err := logic.GetUserByEmail(fo.Email)
	if err != nil {
		entity.ResponseError(c, entity.CodeServerBusy)
		return
	}
	// 返回响应
	entity.ResponseSuccess(c, user)
}

func LoginHandler(c *gin.Context) {
	// 1、获取请求参数及参数校验
	var u *entity.LoginRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			entity.ResponseError(c, entity.CodeInvalidParams) // 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		entity.ResponseErrorWithMsg(c, entity.CodeInvalidParams, config.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}

	// 2、业务逻辑处理——登录
	user, err := logic.Login(u)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", u.Username), zap.Error(err))
		if err.Error() == mysql.ErrorUserNotExit {
			entity.ResponseError(c, entity.CodeUserNotExist)
			return
		}
		entity.ResponseError(c, entity.CodeInvalidParams)
		return
	}
	// 3、返回响应
	entity.ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserID), // js识别的最大值：id值大于1<<53-1  int64: i<<63-1
		"user_name":     user.UserName,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

// GetUserHandler 注册业务
func GetUserHandler(c *gin.Context) {
	var id, err = strconv.Atoi(c.Param("uid"))
	if err != nil {
		entity.ResponseError(c, entity.CodeInvalidParams)
		return
	}
	user, err := logic.GetUser(id)
	// 3.业务处理 —— 获取用户
	if err != nil {
		zap.L().Error("logic.getUser failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {
			entity.ResponseError(c, entity.CodeUserNotExist)
			return
		}
		entity.ResponseError(c, entity.CodeServerBusy)
		return
	}
	// 返回响应
	entity.ResponseSuccess(c, user)
}

// GetUserHandler 注册业务
func GetAllUserHandler(c *gin.Context) {
	users, err := logic.GetAllUsers()
	if err != nil {
		zap.L().Error("logic.getUser failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {
			entity.ResponseError(c, entity.CodeUserNotExist)
			return
		}
		entity.ResponseError(c, entity.CodeServerBusy)
		return
	}
	// 返回响应
	entity.ResponseSuccess(c, users)
}
func SetUserInformationHandler(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("uid"))
	// 1.获取请求参数
	var fo *entity.SetUserInformationForm
	// 2.校验数据有效性
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SetEmail with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			entity.ResponseError(c, entity.CodeInvalidParams) // user 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		entity.ResponseErrorWithMsg(c, entity.CodeInvalidParams, config.RemoveTopStruct(errs.Translate(config.Trans)))
		return // 翻译错误
	}
	log.Printf("fo: %v\n", fo)
	num, _ := logic.UpdateInformation(id, fo.Nickname, fo.Avatar, fo.Password, fo.Email, fo.IsAdmin)
	if num != 0 {
		entity.ResponseSuccess(c, num)
	} else {
		entity.ResponseError(c, entity.CodeInvalidParams)
	}
}

func SetUserEmailHandler(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("uid"))
	// 1.获取请求参数
	var fo *entity.SetUserEmailForm
	// 2.校验数据有效性
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SetEmail with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			entity.ResponseError(c, entity.CodeInvalidParams) // user 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		entity.ResponseErrorWithMsg(c, entity.CodeInvalidParams, config.RemoveTopStruct(errs.Translate(config.Trans)))
		return // 翻译错误
	}
	log.Printf("fo: %v\n", fo)
	num, _ := logic.UpdateEmail(id, fo.Email)
	if num != 0 {
		entity.ResponseSuccess(c, num)
	} else {
		entity.ResponseError(c, entity.CodeInvalidParams)
	}
}

func SetUserPasswordHandler(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("uid"))
	// 1.获取请求参数
	var fo *entity.SetUserPasswordForm
	// 2.校验数据有效性

	if err := c.ShouldBindJSON(&fo); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SetEmail with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			entity.ResponseError(c, entity.CodeInvalidParams) // user 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		entity.ResponseErrorWithMsg(c, entity.CodeInvalidParams, config.RemoveTopStruct(errs.Translate(config.Trans)))
		return // 翻译错误
	}
	log.Printf("fo: %v\n", fo)
	num, _ := logic.SetNewPassword(id, fo.NewPassword, fo.OldPassword)
	if num != 0 {
		entity.ResponseSuccess(c, num)
	} else {
		entity.ResponseError(c, entity.CodeInvalidParams)
	}
}
