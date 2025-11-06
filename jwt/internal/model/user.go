package model

import(
	"time"
)

type User struct{
	ID int64 `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type TokenPair struct{
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessTokenExpireAt time.Time `json:"access_token_expire_at"`
	RefreshTokenExpireAt time.Time `json:"refresh_token_expire_at"`
}

type Claim struct{
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}