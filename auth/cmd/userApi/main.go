package main

import (
	"context"
	"flag"
	"github.com/SigmarWater/messenger/auth/internal/app"
	"log"
)

var serviceConf string

func init() {
	flag.StringVar(&serviceConf, "env", "./postgres/migrations/.env", "path to config")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
