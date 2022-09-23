package client

import (
	"flag"
	"github.com/trendyol/smurfs/go/client/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type SmurfClient struct {
	Logger  *service.LoggerClient
	Auth    *service.AuthClient
	Storage *service.StorageClient
}

type Options struct {
	HostAddress *string
}

func InitializeClient(opt Options) (*SmurfClient, error) {
	if opt.HostAddress == nil {
		flag.StringVar(opt.HostAddress, "host", "localhost:8080", "host address")
	}

	dial, err := grpc.Dial(*opt.HostAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to dial: %+v", err)
		return nil, err
	}

	loggerClient, err := service.NewLoggerClient(dial)
	if err != nil {
		return nil, err
	}

	authClient, err := service.NewAuthClient(dial)
	if err != nil {
		return nil, err
	}

	storageClient, err := service.NewStorageClient(dial)
	if err != nil {
		return nil, err
	}

	smurf := &SmurfClient{
		Logger:  loggerClient,
		Auth:    authClient,
		Storage: storageClient,
	}

	return smurf, nil
}
