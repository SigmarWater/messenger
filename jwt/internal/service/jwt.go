package service

import(
	"github.com/SigmarWater/messenger/jwt/internal/model"
	"time"
)

type JWTToken interface{
	Login(username, password string) (*model.TokenPair, error)
	GetAccessToken(refreshToken string) (string, time.Time, error)
	GetRefreshToken(refreshToken string) (string, time.Time, error)
}