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
	// Flags
	hostAddr string

	// Clients
	authServerClient             protos.AuthServiceClient
	metadataStorageServiceClient protos.MetadataStorageServiceClient
)

type Smurf struct {
	Logger *service.LoggerClient
}

type Options struct {
	Port int
}

func InitializeClient() (*Smurf, error) {
	flag.StringVar(&hostAddr, "host", "localhost:8080", "host address")
	dial, err := grpc.Dial(hostAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to dial: %+v", err)
		return nil, err
	}

	authServerClient = protos.NewAuthServiceClient(dial)
	metadataStorageServiceClient = protos.NewMetadataStorageServiceClient(dial)

	loggerClient, err := service.NewLoggerClient(dial)
	if err != nil {
		return nil, err
	}

	smurf := &Smurf{
		Logger: loggerClient,
	}

	return smurf, nil
}
