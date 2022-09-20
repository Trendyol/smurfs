package service

import (
	"context"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Auth interface {
	GetToken() (string, error)
	GetUserInfo() (*UserInfo, error)
}

type UserInfo struct {
	Username string
	Email    string
}

type Token struct {
	AccessToken  string
	RefreshToken string
	RptToken     string
}

type AuthClient struct {
}

var (
	authGrpcClient protos.AuthServiceClient
)

func NewAuthClient(dial *grpc.ClientConn) (*AuthClient, error) {
	authGrpcClient = protos.NewAuthServiceClient(dial)

	return &AuthClient{}, nil
}

func (a *AuthClient) GetToken() (*Token, error) {
	token, err := authGrpcClient.GetToken(context.Background(), &emptypb.Empty{})
	if err != nil {
		// TODO: (peacecwz): We can customize error messages. Sometimes rpc call errors are not meaningful.
		return nil, err
	}

	return &Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		RptToken:     token.RptToken,
	}, nil
}
