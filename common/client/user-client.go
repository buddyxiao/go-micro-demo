package client

import (
	"common/consts"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	user "user-service/proto"
)

var (
	userClient user.UserService
)

func GetUserClient() user.UserService {
	if userClient == nil {
		// 注册中心查找
		consulRegistry := consul.NewRegistry(registry.Addrs(consts.ConsulHostPort))
		service := micro.NewService(
			micro.Registry(consulRegistry),
			micro.Client(grpcc.NewClient()),
			micro.Server(grpcs.NewServer()),
		)
		userClient = user.NewUserService(consts.UserServiceName, service.Client())
	}
	return userClient
}
