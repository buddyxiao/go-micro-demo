package main

import (
	"common/config"
	"common/consts"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4/registry"
	"os"
	"rand-service/handler"
	pb "rand-service/proto"
	"user-service/dao"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

func init() {
	dao.SetDefault(config.DB())
}
func main() {
	// 添加注册中心
	r := consul.NewRegistry(registry.Addrs(consts.ConsulHostPort))
	// 链路追踪
	tracer, closer, err := config.NewJaegerTracer(consts.RandServiceName, consts.JaegerHostPort)
	defer closer.Close()
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Registry(r),
	)
	srv.Init(
		micro.Name(consts.RandServiceName),
		micro.Version(consts.RandServiceVersion),
		//micro.WrapHandler(middleware.JaegerMiddleware(tracer)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
		micro.WrapClient(opentracing.NewClientWrapper(tracer)),
	)
	// Register handler
	if err := pb.RegisterRandHandler(srv.Server(), new(handler.RandService)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
