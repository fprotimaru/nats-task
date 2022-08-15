package main

import (
	"context"
	"log"
	"nats_ex/internal/application"
	"nats_ex/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.New("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, syscall.SIGTERM, os.Interrupt)

	app := application.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quitSignal
		cancel()
	}()

	app.Run(ctx)
}
