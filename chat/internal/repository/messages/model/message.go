package model

import (
	"database/sql"
)

type MessageRepository struct{
	Id_chat int
	From_user string
	Text_message string 
	Time_at sql.NullTime 
}