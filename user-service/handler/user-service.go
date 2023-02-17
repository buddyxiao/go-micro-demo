package handler

import (
	"common/domain/bo"
	"common/event"
	"context"
	"fmt"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	logservice "log-service/proto"
	rand "rand-service/proto"
	pb "user-service/proto"
	"user-service/service"
)

type UserService struct {
	c    client.Client
	rand rand.RandService
}

func NewUserService(c client.Client, rand rand.RandService) *UserService {
	return &UserService{c: c, rand: rand}
}

func (u *UserService) Register(ctx context.Context, request *pb.UserRegistryRequest, response *pb.UserRegistryResponse) error {
	register, err := service.User().Register(bo.RegistryInput{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	})
	response.Msg = register.Msg
	return err
}

func (u *UserService) Login(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	login, err := service.User().Login(bo.LoginInput{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		go loginLog(ctx, u.c, request, false)
		logger.Debug("登录失败，err:", err)
		return err
	}
	response.Msg = login.Msg
	response.Token = login.Token
	// 模拟调用其他服务
	randResponse, err := u.rand.GetRand(ctx, &rand.RandRequest{Max: 100})
	if err != nil {
		logger.Info("调用rand-service错误,err:", err)
		return err
	}
	logger.Info("rand-service返回随机数: ", randResponse.Result)
	// 记录登录日志
	go loginLog(ctx, u.c, request, true)
	return nil
}

func loginLog(ctx context.Context, c client.Client, req *pb.UserLoginRequest, succ bool) {
	var loginLogData logservice.LogRequest
	if succ {
		loginLogData.Msg = fmt.Sprintf("用户%s登录成功", req.Username)
	} else {
		loginLogData.Msg = fmt.Sprintf("用户%s登录失败", req.Username)
	}
	message := c.NewMessage(event.Login, &loginLogData)
	err := c.Publish(ctx, message)
	if err != nil {
		logger.Info("消息发送失败,err:", err)
	}
}
