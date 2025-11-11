package model

import (
	"database/sql"
)

type UserRepository struct {
	Id       int
	Name     string
	Email    string
	Password string
	Role     string
	CreateAt sql.NullTime
	UpdateAt sql.NullTime
}

type CacheUser struct {
	UUID     string `redis:"uuid"`
	Name     string `redis:"name"`
	Email    string `redis:"email"`
	Role     string `redis:"role"`
	CreateAt int64  `redis:"create_at"`
	UpdateAt int64  `redis:"update_at, omitempty"`
}
