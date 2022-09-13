package client

import (
	"flag"
	"github.com/trendyol/smurfs/client/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	// Flags
	hostAddr string

	// Clients
	logServiceClient             protos.LogServiceClient
	authServerClient             protos.AuthServiceClient
	metadataStorageServiceClient protos.MetadataStorageServiceClient
)

func Initialize() {
	flag.StringVar(&hostAddr, "host", "localhost:8080", "host address")
	dial, err := grpc.Dial(hostAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to dial: %+v", err)
	}

	logServiceClient = protos.NewLogServiceClient(dial)
	authServerClient = protos.NewAuthServiceClient(dial)
	metadataStorageServiceClient = protos.NewMetadataStorageServiceClient(dial)
}
