package service

import (
	"github.com/trendyol/smurfs/go/protos"
	"io"
	"log"
)

var _ = protos.LogServiceServer(&logServiceImpl{})

type logServiceImpl struct {
	protos.UnimplementedLogServiceServer
}

func NewLogService() protos.LogServiceServer {
	return &logServiceImpl{}
}

func (s *logServiceImpl) Info(server protos.LogService_InfoServer) error {
	for {
		logRequest, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		log.Printf("[info] %s", logRequest.Msg)
	}

	log.Print("info stream ended successfully")
	return nil
}

func (s *logServiceImpl) Warn(server protos.LogService_WarnServer) error {
	for {
		logRequest, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		log.Printf("[warn] %s", logRequest.Msg)
	}

	log.Print("warn stream ended successfully")
	return nil
}

func (s *logServiceImpl) Error(server protos.LogService_ErrorServer) error {
	for {
		logRequest, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		log.Printf("[error] %s", logRequest.Msg)
	}

	log.Print("error stream ended successfully")
	return nil
}
