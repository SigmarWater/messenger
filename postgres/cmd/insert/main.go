package main

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	sq "github.com/Masterminds/squirrel"
)

const(
	dbDNS = "host=84.22.148.185 port=50000 dbname=messenger user=sigmawater password=sigmawater sslmode=disable"
)

func main(){
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDNS)

	if err != nil{
		log.Println(err)
		return 
	}

	defer pool.Close()

	builderInsert := sq.Insert("users").
	PlaceholderFormat(sq.Dollar).
	Columns("name", "email", "password_hash").
	Values(gofakeit.Name(), gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 16)). 
	Suffix("RETURNING id")

	query, arguments, err := builderInsert.ToSql()
	if err != nil{
		log.Println(err)
		return
	}

	var userID string 
	err = pool.QueryRow(ctx, query, arguments...).Scan(&userID)
	if err != nil{
		log.Println(err)
		return
	}

	log.Printf("inserted user with id: %s\n", userID)
}