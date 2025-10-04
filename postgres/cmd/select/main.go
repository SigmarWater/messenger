package main

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
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

	if err := pool.Ping(ctx); err != nil{
		log.Println(err)
		return
	}

	builderSelect := sq.Select("id", "name", "email", "role").
	From("users").
	PlaceholderFormat(sq.Dollar). 
	OrderBy("id DESC").
	Limit(10) 

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to select notes: %v", err)
	}

	var id int 
	var name string 
	var email string 
	var role string 

	for rows.Next(){
		if err := rows.Scan(&id, &name, &email, &role); err != nil{
			log.Fatalf("failed to scan note: %v", err)
			return
		}

		log.Printf("id: %d name: %s email: %s role: %s\n", id, name, email, role)
	}

}