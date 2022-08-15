package publisher

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
	"log"
	"nats_ex/internal/entity"
	"time"
)

type Publisher struct {
	ctx  context.Context
	conn *nats.Conn
}

func NewPublisher(ctx context.Context, conn *nats.Conn) *Publisher {
	return &Publisher{ctx: ctx, conn: conn}
}

func (p Publisher) Run() {
	ticker := time.NewTicker(time.Second * 2)

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-ticker.C:
			var user = entity.User{
				Username: "admin",
				Password: "random_password",
			}

			data, err := jsoniter.Marshal(user)
			if err != nil {
				return
			}

			reply := nats.NewInbox()

			err = p.conn.PublishRequest("user", reply, data)
			if err != nil {
				return
			}

			log.Printf("Sent message: %+v\n", user)
		}
	}
}
