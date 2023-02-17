package handler

import (
	"context"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	logservice "log-service/proto"
)

type LogService struct{}

func (log *LogService) Print(ctx context.Context, request *logservice.LogRequest) error {
	md, _ := metadata.FromContext(ctx)
	logger.Infof("Received event %+v with metadata %+v\n", request, md)
	return nil
}
