package model 

import(
	"time"
)

type MessageService struct{
	ChatId string;
	From_user int;
	Text_message string;
	Time_at time.Time 
}