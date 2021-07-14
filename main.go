package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	go nc.Publish("foo", []byte("Hello World"))
	go nc.Publish("foo", []byte("Hello World2"))

	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("foo", ch)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	for msg := range ch {
		fmt.Println(string(msg.Data))
	}

	defer sub.Unsubscribe()
	defer sub.Drain()
}
