package subscriber

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
	"log"
	"nats_ex/internal/entity"
)

type Subscriber struct {
	conn *nats.Conn
}

func NewSubscriber(conn *nats.Conn) *Subscriber {
	return &Subscriber{conn: conn}
}

func (s Subscriber) Run() {
	_, err := s.conn.Subscribe("user", func(msg *nats.Msg) {
		if err := msg.Respond([]byte("got user")); err != nil {
			log.Println(err)
			return
		}
		var user entity.User
		if err := jsoniter.Unmarshal(msg.Data, &user); err != nil {
			log.Println(err)
			return
		}
		log.Printf("Got message: %+v\n", user)
		entity.UsersDB.Add(user)
	})
	if err != nil {
		log.Println(err)
		return
	}
}
