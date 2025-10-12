package model 

import(
	"time"
)

type UserService struct{ 
	Id int
	Name string 
	Email string 
	EnterPassword string 
	ConfirmPassword string
	Role string
	CreateAt time.Time 
	UpdateAt time.Time
}