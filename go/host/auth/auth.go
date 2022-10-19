package auth

import (
	"context"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	RptToken     string `json:"rpt_token"`
}

type UserInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Auth interface {
	GetToken() (TokenResponse, error)
	GetUserInfo() (UserInfo, error)
}

//goland:noinspection GoNameStartsWithPackageName
type AuthService struct {
	root Auth
	protos.UnimplementedAuthServiceServer
}

func NewAuthService(auth Auth) *AuthService {
	return &AuthService{
		root: auth,
	}
}

func (a AuthService) GetToken(ctx context.Context, empty *emptypb.Empty) (*protos.TokenResponse, error) {
	token, err := a.root.GetToken()
	if err != nil {
		return nil, err
	}
	return &protos.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		RptToken:     token.RptToken,
	}, nil
}

func (a AuthService) GetUserInfo(ctx context.Context, empty *emptypb.Empty) (*protos.UserInfo, error) {
	userInfo, err := a.root.GetUserInfo()
	if err != nil {
		return nil, err
	}
	return &protos.UserInfo{
		Username: userInfo.Username,
		Email:    userInfo.Email,
	}, nil
}
