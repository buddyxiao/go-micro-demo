package main

import (
	"common/client"
	"common/config"
	"common/consts"
	"gin-api-gateway/handler"
	"gin-api-gateway/middleware"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
)

func main() {
	r := gin.Default()
	jaeger, closer, err := config.NewJaegerTracer("api-gateway", consts.JaegerHostPort)
	defer closer.Close()
	if err != nil {
		logger.Info("api-gateway init jaeger failed, err:", err)
	}
	r.Use(
		middleware.JaegerGatewayMiddleware(jaeger),
		middleware.ErrorMiddleware(),
		middleware.ResponseMiddleware(),
		//opengintracing.InjectToHeaders(jaeger, false),
	)
	rand := r.Group("/v1/rand")
	apiHandler := handler.GetAPIHandler(client.GetRandClient(), client.GetUserClient())
	rand.GET("", apiHandler.Rand)
	r.Group("/v1/user").
		POST("/registry", apiHandler.Registry).
		POST("/login", apiHandler.Login)
	err = r.Run(":8080")
	if err != nil {
		panic(err)
		return
	}
}
