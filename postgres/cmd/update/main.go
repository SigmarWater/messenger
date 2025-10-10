package main

import (
	"context"
	"log"

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

	if err := pool.Ping(ctx); err != nil{
		log.Println(err)
		return
	}

	builderUpdate := sq.Update("users").
	PlaceholderFormat(sq.Dollar). 
	Set("name", "Artem").
	Set("email", "artemudalcov05@gmail.com").
	Set("role", "admin").
	Where(sq.Eq{"id": 1})

	query, args, err := builderUpdate.ToSql()
	if err != nil{
		log.Println(err)
		return 
	}

	log.Println("query:", query)
	log.Println("args:", args)

	res, err := pool.Exec(ctx, query, args...)

	if err != nil{
		log.Println(err)
		return
	}
	log.Println(res.RowsAffected())

	var id int 
	var name string 
	var email string 
	var role string

	err = pool.QueryRow(ctx, "select id, name, email, role from users where id=$1;", 1).Scan(&id, &name, &email, &role)

	if err != nil{
		log.Println(err)
	}
	log.Printf("id: %d name: %s email: %s role: %s\n", id, name, email, role)
}