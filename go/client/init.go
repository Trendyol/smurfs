package client

import (
	"flag"
	"github.com/trendyol/smurfs/go/client/pkg/service"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	// Clients
	metadataStorageServiceClient protos.MetadataStorageServiceClient
)

type Smurf struct {
	Logger *service.LoggerClient
	Auth   *service.AuthClient
}

type Options struct {
	HostAddress *string
}

func InitializeClient(opt Options) (*Smurf, error) {
	if opt.HostAddress == nil {
		flag.StringVar(opt.HostAddress, "host", "localhost:8080", "host address")
	}

	dial, err := grpc.Dial(*opt.HostAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to dial: %+v", err)
		return nil, err
	}

	metadataStorageServiceClient = protos.NewMetadataStorageServiceClient(dial)

	loggerClient, err := service.NewLoggerClient(dial)
	if err != nil {
		return nil, err
	}

	authClient, err := service.NewAuthClient(dial)
	if err != nil {
		return nil, err
	}

	smurf := &Smurf{
		Logger: loggerClient,
		Auth:   authClient,
	}

	return smurf, nil
}
