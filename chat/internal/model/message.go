package model 

import(
	"time"
)

type MessageService struct{
	IdMessage int
	ChatId int
	ChatName string
	FromUser string
	TextMessage string
	TimeAt time.Time 
}