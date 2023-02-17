package client

import (
	"common/consts"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	rand "rand-service/proto"
)

var (
	randClient rand.RandService
)

func GetRandClient() rand.RandService {
	if randClient == nil {
		consulRegistry := consul.NewRegistry(
			registry.Addrs(consts.ConsulHostPort))
		service := micro.NewService(
			//micro.Server(grpcs.NewServer()),
			micro.Client(grpcc.NewClient()),
			micro.Registry(consulRegistry))
		randClient = rand.NewRandService(consts.RandServiceName, service.Client())
	}
	return randClient

}
