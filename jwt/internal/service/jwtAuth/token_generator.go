package jwtAuth

import(
	"github.com/SigmarWater/messenger/jwt/internal/model"
	"time"
	"github.com/golang-jwt/jwt/v5"
) 

func (s *JWTService) generateTokenPair(user *model.User) (*model.TokenPair, error){
	accessToken, accessExpireAt, err := s.generateAccessToken(user)
	if err != nil{
		return nil, err
	}

	refreshToken, refreshExpireAt, err := s.generateRefrestToken(user)
	if err != nil{
		return nil, err 
	}

	return &model.TokenPair{
		AccessToken: accessToken,
		AccessTokenExpireAt: accessExpireAt,
		RefreshToken: refreshToken,
		RefreshTokenExpireAt: refreshExpireAt,
	}, nil
}

func (s *JWTService) generateAccessToken(user *model.User)(string, time.Time, error){
	expiresAt := time.Now().Add(accessTokenTTL)

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      expiresAt.Unix(),
		"iat":      time.Now().Unix(),
		"type":     "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

func (s *JWTService) generateRefrestToken(user *model.User)(string, time.Time, error){
	expireAt := time.Now().Add(refreshTokenTTL)

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"username": user.Username,
		"expire_at": expireAt.Unix(), 
		"iat":      time.Now().Unix(),
		"type": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(refreshTokenSecret))
	if err != nil{
		return "", time.Time{}, err 
	}

	return tokenString, expireAt, nil
}