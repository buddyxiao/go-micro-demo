package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
	"net/http"
	rand "rand-service/proto"
)

func (a *APIHandler) Rand(ctx *gin.Context) {
	var request rand.RandRequest
	ctx.ShouldBindQuery(&request)
	logger.Info("request:", request)
	resp, _ := a.randClient.GetRand(context.Background(), &request)
	ctx.JSON(http.StatusOK, gin.H{
		"result": resp.GetResult(),
	})
}
