package handler

import (
	context "context"
	pb "rand-service/proto"
	"rand-service/service"
)

type RandService struct{}

func (r *RandService) GetRand(ctx context.Context, request *pb.RandRequest, response *pb.RandResponse) error {
	response.Result = service.Rand().GetRand(request.GetMax())
	return nil
}
