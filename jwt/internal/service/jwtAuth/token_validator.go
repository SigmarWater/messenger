package jwtAuth

import (
	"github.com/SigmarWater/messenger/jwt/internal/model"
	"github.com/golang-jwt/jwt/v5"
	
)


func (s *JWTService) validateRefreshToken(tokenString string) (*model.Claim, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, ErrInvalidToken
		}
		return []byte(refreshTokenSecret), nil 
	})

	if err != nil && !token.Valid{
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok{
		return nil, ErrInvalidToken
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh"{
		return nil, ErrInvalidToken
	}

	userID, ok := claims["user_id"].(float64)
	if !ok{
		return nil,ErrInvalidToken
	}

	username, ok := claims["username"].(string)
	if !ok{
		return nil, ErrInvalidToken
	}

	return &model.Claim{
		UserID: int64(userID),
		Username: username,
	}, nil
}