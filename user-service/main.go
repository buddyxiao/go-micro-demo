package main

import (
	"common/client"
	"common/config"
	"common/consts"
	//commonmiddleware "common/middleware"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	opentracingplugins "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4/registry"
	"os"
	"user-service/dao"
	"user-service/handler"
	"user-service/middleware"
	pb "user-service/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

func init() {
	dao.SetDefault(config.DB())
}

func main() {
	// 服务注册
	consultRegistry := consul.NewRegistry(registry.Addrs(consts.ConsulHostPort))
	// 链路追踪
	tracer, closer, err := config.NewJaegerTracer(consts.UserServiceName, consts.JaegerHostPort)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	defer closer.Close()
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Registry(consultRegistry),
	)
	srv.Init(
		micro.Name(consts.UserServiceName),
		micro.Version(consts.UserServiceVersion),
		micro.WrapHandler(middleware.RateLimiterWrapper()),
		//micro.WrapHandler(commonmiddleware.JaegerMiddleware(tracer)),
		micro.WrapHandler(opentracingplugins.NewHandlerWrapper(tracer)),
		micro.WrapClient(opentracingplugins.NewClientWrapper(tracer)),
	)
	// Register handler
	randClient := client.GetRandClient()
	if err = pb.RegisterUserServiceHandler(srv.Server(), handler.NewUserService(srv.Client(), randClient)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
