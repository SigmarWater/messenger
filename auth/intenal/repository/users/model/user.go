package model 

import "database/sql"

type UserRepository struct{
	Id int 
	Name string
	Email string 
	Password string
	Role string 
	CreateAt sql.NullTime
	UpdateAt sql.NullTime
}