package logger

import (
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

type LogService struct {
	protos.UnimplementedLogServiceServer
	root Logger
}

func NewLogService(logger Logger) *LogService {
	return &LogService{
		root: logger,
	}
}

func (logSrv LogService) Info(stream protos.LogService_InfoServer) error {
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}

		if err != nil {
			return err
		}

		logSrv.root.Info(l.Msg)
	}
}

func (logSrv LogService) Debug(stream protos.LogService_DebugServer) error {
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}

		if err != nil {
			return err
		}

		logSrv.root.Debug(l.Msg)
	}
}

func (logSrv LogService) Warn(stream protos.LogService_WarnServer) error {
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}

		if err != nil {
			return err
		}

		logSrv.root.Warn(l.Msg)
	}
}

func (logSrv LogService) Error(stream protos.LogService_ErrorServer) error {
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}

		if err != nil {
			return err
		}

		logSrv.root.Error(l.Msg)
	}
}

func (logSrv LogService) Fatal(stream protos.LogService_FatalServer) error {
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}

		if err != nil {
			return err
		}

		logSrv.root.Fatal(l.Msg)
	}
}
