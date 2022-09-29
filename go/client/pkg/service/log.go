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

func NewLoggerClient(dial *grpc.ClientConn) (*LoggerClient, error) {
	client = protos.NewLogServiceClient(dial)
	var err error

	debugService, err = client.Debug(context.Background())
	if err != nil {
		return nil, err
	}

	warnService, err = client.Warn(context.Background())
	if err != nil {
		return nil, err
	}

	errorService, err = client.Error(context.Background())
	if err != nil {
		return nil, err
	}

	fatalService, err = client.Fatal(context.Background())
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
	infoService, err := client.Info(context.Background())
	if err != nil {
		return
	}

	err = infoService.Send(&protos.LogRequest{
		Msg: message,
	})

	if err != nil {
		return
	}

	infoService.CloseSend()
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
