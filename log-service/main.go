package main

import (
	"common/consts"
	"common/event"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"log-service/handler"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

func main() {
	consulRegistry := consul.NewRegistry(registry.Addrs(consts.ConsulHostPort))
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Registry(consulRegistry),
	)
	srv.Init(
		micro.Name(consts.LogServiceName),
		micro.Version(consts.LogServiceVersion),
	)
	// 注册登录日志事件（用于异步消息）
	err := micro.RegisterSubscriber(event.Login, srv.Server(), new(handler.LogService))
	if err != nil {
		logger.Info("log-service 订阅注册失败, err:", err)
		return
	}
	// Register handler

	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
