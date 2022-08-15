package main

import (
	jsoniter "github.com/json-iterator/go"
	"log"

	"github.com/nats-io/nats.go"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	user1 := User{
		Username: "admin",
		Password: "admin12345",
	}

	data, err := jsoniter.Marshal(user1)
	if err != nil {
		log.Fatalf("error on marshalling user: %+v\n", user1)
	}

	// Simple Publisher
	err = nc.Publish("user", data)
	if err == nil {
		log.Println("Message published")
	}

	defer nc.Close()
}
