package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	sq "github.com/Masterminds/squirrel"
)

const(
	dbDNS = "host=84.22.148.185 port=50000 user=sigmawater password=sigmawater dbname=messenger sslmode=disable"
)

func main(){
	ctx := context.Background() 

	pool, err := pgxpool.Connect(ctx, dbDNS)
	if err != nil{
		log.Println(err)
		return 
	}

	if err := pool.Ping(ctx); err != nil{
		log.Println(err)
		return 
	}

	builderDelete := sq.Delete("users").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id":2}).Suffix("RETURNING id")

	query, args, err := builderDelete.ToSql()

	log.Println("query:", query)
	log.Println("args:", args)

	if err != nil{
		log.Printf("to sql error: %v\n", err)
		return 
	}

	var id int 
	err = pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil{
		log.Printf("querry row error: %v\n", err)
		return 
	}

	log.Printf("deletted id: %d\n", id)
}