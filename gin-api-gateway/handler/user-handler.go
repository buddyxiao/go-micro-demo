package handler

import (
	"common/domain/dto"
	"common/domain/vo"
	"context"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
	userservice "user-service/proto"
)

func (a *APIHandler) Registry(ctx *gin.Context) {
	var req dto.RegistryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Set("error", err)
		return
	}
	register, err := a.userClient.Register(context.Background(), &userservice.UserRegistryRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		ctx.Set("error", err)
		return
	}
	ctx.Set("data", vo.RegistryVo{Msg: register.Msg})
}

func (a *APIHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Set("error", err)
		return
	}
	login, err := a.userClient.Login(ctx.Request.Context(), &userservice.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		ctx.Set("error", err)
		logger.Info("登录失败：err:", err)
		return
	}
	ctx.Set("data", login)
}
