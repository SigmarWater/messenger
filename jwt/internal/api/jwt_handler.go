package api

import (
	"context"

	"github.com/SigmarWater/messenger/jwt/internal/service"
	"github.com/SigmarWater/messenger/jwt/internal/service/jwtAuth"
	pb "github.com/SigmarWater/messenger/jwt/pkg/api/jwt_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"errors"
)

type JWTHandler struct{
	pb.UnimplementedJWTServiceServer 
	jwtService service.JWTToken
}

func NewJWTHandler() *JWTHandler{
	service := jwtAuth.NewJWTService()
	return &JWTHandler{
		jwtService: service,
	}
}

func (j *JWTHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error){
	if req.GetUsername() == "" || req.GetPassword() == ""{
		return nil, status.Error(codes.InvalidArgument, "username and password are required")
	}
	
	tokenPair, err := j.jwtService.Login(req.GetUsername(), req.GetPassword())
	if err != nil{
		if errors.Is(err, jwtAuth.ErrInvalidCredentials){
			return nil, status.Error(codes.Unauthenticated, "invalid credentials")
		}
		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &pb.LoginResponse{
		AccessToken: tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		AccessTokenExpireAt: timestamppb.New(tokenPair.AccessTokenExpireAt),
		RefreshTokenExpireAt: timestamppb.New(tokenPair.RefreshTokenExpireAt),
	}, nil
}


func (j *JWTHandler) GetAccessToken(ctx context.Context, req *pb.GetAccessTokenRequest) (*pb.GetAccessTokenResponse, error){
	if req.RefreshToken == ""{
		return nil, status.Error(codes.InvalidArgument, "refresh token is required")
	}

	accessToken, expireAccessToken, err := j.jwtService.GetAccessToken(req.GetRefreshToken())
	if err != nil {
		if errors.Is(err, jwtAuth.ErrInvalidToken) {
			return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
		}
		return nil, status.Error(codes.Internal, "failed to get access token")
	}

	return &pb.GetAccessTokenResponse{
		AccessToken: accessToken,
		AccessTokenExpireAt: timestamppb.New(expireAccessToken),
	}, nil
}

func (j *JWTHandler) GetRefreshToken(ctx context.Context, req *pb.GetRefreshTokenRequest) (*pb.GetRefreshTokenResponse, error){
	if req.RefreshToken == ""{
		return nil, status.Error(codes.InvalidArgument, "refresh token is required")
	}

	refreshToken, expireRefreshToken, err := j.jwtService.GetRefreshToken(req.GetRefreshToken())
	if err != nil {
		if errors.Is(err, jwtAuth.ErrInvalidToken) {
			return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
		}
		return nil, status.Error(codes.Internal, "failed to get access token")
	}

	return &pb.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
		RefreshTokenExpireAt: timestamppb.New(expireRefreshToken),
	}, nil
}
