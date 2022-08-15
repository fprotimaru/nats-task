package application

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"nats_ex/internal/config"
	"nats_ex/internal/publisher"
	"nats_ex/internal/subscriber"
	"time"
)

type Application struct {
	cfg      *config.Config
	natsConn *nats.Conn
}

func New(cfg *config.Config) *Application {
	return &Application{cfg: cfg}
}

func (app *Application) Run(ctx context.Context) {
	if len(app.cfg.NatsURL) == 0 {
		app.cfg.NatsURL = nats.DefaultURL
	}

	var err error
	app.natsConn, err = nats.Connect(app.cfg.NatsURL, nats.ReconnectWait(time.Second*5))
	if err != nil {
		log.Fatalln(err)
	}

	pub := publisher.NewPublisher(ctx, app.natsConn)

	sub := subscriber.NewSubscriber(app.natsConn)
	go sub.Run()

	go func() {
		<-ctx.Done()
		// gracefully shutdown
		fmt.Println("shutting down NATS connection...")
		if err := app.natsConn.Drain(); err != nil {
			log.Fatalln(err)
		}
		app.natsConn.Close()
	}()

	pub.Run()
}
