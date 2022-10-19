package storage

import (
	"context"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MetadataStorage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type MetadataStorageService struct {
	protos.MetadataStorageServiceServer
	root MetadataStorage
}

func NewMetadataStorageService(metadataStorage MetadataStorage) *MetadataStorageService {
	return &MetadataStorageService{
		root: metadataStorage,
	}
}

func (m MetadataStorageService) Get(ctx context.Context, request *protos.GetMetadataStorageRequest) (*protos.MetadataStorageResponse, error) {
	value, err := m.root.Get(request.Key)
	if err != nil {
		return nil, err
	}

	return &protos.MetadataStorageResponse{
		Value: value,
	}, nil
}

func (m MetadataStorageService) Set(ctx context.Context, request *protos.SetMetadataStorageRequest) (*emptypb.Empty, error) {
	err := m.root.Set(request.Key, request.Value)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
