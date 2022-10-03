package service

import (
	"context"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/grpc"
)

type LoggerClient struct {
}

var (
	client       protos.LogServiceClient
	debugService protos.LogService_DebugClient
	infoService  protos.LogService_InfoClient
	warnService  protos.LogService_WarnClient
	errorService protos.LogService_ErrorClient
	fatalService protos.LogService_FatalClient
)

func NewLoggerClient(dial *grpc.ClientConn, ctx context.Context) (*LoggerClient, error) {
	client = protos.NewLogServiceClient(dial)
	var err error

	infoService, err = client.Info(ctx)
	if err != nil {
		return nil, err
	}

	debugService, err = client.Debug(ctx)
	if err != nil {
		return nil, err
	}

	warnService, err = client.Warn(ctx)
	if err != nil {
		return nil, err
	}

	errorService, err = client.Error(ctx)
	if err != nil {
		return nil, err
	}

	fatalService, err = client.Fatal(ctx)
	if err != nil {
		return nil, err
	}

	return &LoggerClient{}, nil
}

func (l *LoggerClient) Debug(message string, args ...interface{}) {
	err := debugService.Send(&protos.LogRequest{
		Msg: message,
	})

	if err != nil {
		return
	}
}

func (l *LoggerClient) Info(message string, args ...interface{}) {
	err := infoService.Send(&protos.LogRequest{
		Msg: message,
	})

	if err != nil {
		return
	}
}

func (l *LoggerClient) Warn(message string, args ...interface{}) {
	err := warnService.Send(&protos.LogRequest{
		Msg: message,
	})

	if err != nil {
		return
	}
}

func (l *LoggerClient) Error(message string, args ...interface{}) {
	err := errorService.Send(&protos.LogRequest{
		Msg: message,
	})

	if err != nil {
		return
	}
}

func (l *LoggerClient) Fatal(message string, args ...interface{}) {
	err := fatalService.Send(&protos.LogRequest{
		Msg: message,
	})

	if err != nil {
		return
	}
}
