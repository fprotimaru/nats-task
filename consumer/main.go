package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	var users []User

	// Simple Async Subscriber
	nc.Subscribe("user", func(msg *nats.Msg) {
		var user User
		if err := jsoniter.Unmarshal(msg.Data, &user); err != nil {
			log.Printf("error on unmarshalling json data: %s", msg.Data)
		}
		users = append(users, user)

		log.Println(users)
	})

	// Keep the connection alive
	runtime.Goexit()
}
