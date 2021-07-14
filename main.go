package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	//defer nc.Close()
	nc.Publish("foo", []byte("Hello World"))
	nc.Publish("foo", []byte("Hello World"))

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("foo", ch)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	//msg := <-ch
	for msg := range ch {
		fmt.Println(string(msg.Data))
	}

	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()
}
