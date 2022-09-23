package service

import (
	"context"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/grpc"
)

type Storage interface {
	Set(key string, value string) error
	Get(key string) (*string, error)
}

type StorageClient struct {
}

var (
	storageGrpcClient protos.MetadataStorageServiceClient
)

func NewStorageClient(dial *grpc.ClientConn) (*StorageClient, error) {
	storageGrpcClient = protos.NewMetadataStorageServiceClient(dial)

	return &StorageClient{}, nil
}

func (s *StorageClient) Set(key string, value string) error {
	_, err := storageGrpcClient.Set(context.Background(), &protos.SetMetadataStorageRequest{
		Key:   key,
		Value: value,
	})
	return err
}

func (s *StorageClient) Get(key string) (*string, error) {
	response, err := storageGrpcClient.Get(context.Background(), &protos.GetMetadataStorageRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	}

	return &response.Value, nil
}
