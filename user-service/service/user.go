package service

import (
	"common/domain/bo"
	"common/xerror"
	"github.com/hashicorp/go-uuid"
	"go-micro.dev/v4/logger"
	"time"
	"user-service/dao"
	"user-service/model"
)

type userservice struct {
}

func User() userservice {
	return userservice{}
}

func (u userservice) Login(intput bo.LoginInput) (output bo.LoginOutput, err error) {
	user, _ := dao.User.Where(dao.User.Username.Eq(intput.Username)).Take()
	if user == nil {
		err = xerror.LoginFailError
		return
	}
	if *user.Password == intput.Password {
		output.Msg = "登录成功"
		output.Token, _ = uuid.GenerateUUID()
		return
	}
	err = xerror.LoginFailError
	return
}

func (u userservice) Register(input bo.RegistryInput) (output bo.RegistryOutput, err error) {
	// 数据校验
	user, _ := dao.User.Where(dao.User.Username.Eq(input.Username)).Take()
	if user != nil {
		return output, xerror.UserExistError
	}
	now := time.Now()
	addUser := model.User{
		Username:   &input.Username,
		Password:   &input.Password,
		Email:      &input.Email,
		CreateTime: &now,
	}
	err = dao.User.Create(&addUser)
	if err != nil {
		logger.Debug("创建用户失败:", err)
		err = xerror.CreateFailError
	}
	output.Msg = "注册成功"
	return
}
