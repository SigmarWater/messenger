package jwtAuth

import (
	"github.com/SigmarWater/messenger/jwt/internal/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"errors"
	"time"
)

func (s *JWTService) Login(username, password string) (*model.TokenPair, error){
	user, exists := s.users[username]

	if !exists{
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		return nil, ErrInvalidCredentials
	}

	tokenPair, err := s.generateTokenPair(&user)
	return tokenPair, err 
}

func (s *JWTService) GetAccessToken(refreshToken string) (string, time.Time, error){
	claims, err := s.validateRefreshToken(refreshToken)
	if err != nil{
		if errors.Is(err, ErrInvalidToken){
			return "", time.Time{}, status.Error(codes.Internal, "invalid token")
		}
		return "", time.Time{}, err 
	}

	user, exists := s.users[claims.Username]
	if !exists{
		return "", time.Time{}, ErrInvalidToken
	}

	return s.generateAccessToken(&user)
}

func (s *JWTService) GetRefreshToken(refreshToken string) (string, time.Time, error){
	claims, err := s.validateRefreshToken(refreshToken)
	if err != nil{
		if errors.Is(err, ErrInvalidToken){
			return "", time.Time{}, status.Error(codes.Internal, "invalid token")
		}
		return "", time.Time{}, err 
	}

	user, exists := s.users[claims.Username]
	if !exists{
		return "", time.Time{}, ErrInvalidToken
	}

	return s.generateRefrestToken(&user)
}

