package middleware

import (
	"context"
	"errors"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
	"os"
)

func init() {
	// 务必先进行初始化
	err := sentinel.InitDefault()
	if err != nil {
		logger.Debug("限流器启动失败")
		os.Exit(1)
	}
	// 配置规则
	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "user-service",
			Threshold:              1,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		logger.Debug("限流器启动失败")
		os.Exit(1)
	}
}

// RateLimiterWrapper 限流服务访问
func RateLimiterWrapper() server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			entry, blockError := sentinel.Entry("user-service")
			if blockError != nil {
				return errors.New("user-service 太忙了,请重试")
			} else {
				defer entry.Exit()
				return handlerFunc(ctx, req, rsp)
			}
		}
	}
}
