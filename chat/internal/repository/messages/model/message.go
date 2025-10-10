package model

import (
	"database/sql"
)

type MessageRepository struct{
	IdMessage int
	IdChat int
	ChatName string
	FromUser string
	TextMessage string 
	TimeAt sql.NullTime 
}