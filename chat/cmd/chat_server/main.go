package main

import (
	"context"
	"flag"
	"github.com/SigmarWater/messenger/chat/internal/app"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "env", "c:\\users\\admin\\desktop\\goproject\\messenger\\postgres\\migrations\\.env", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	application, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = application.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
