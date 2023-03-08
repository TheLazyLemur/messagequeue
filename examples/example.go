package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"time"

	"lemur/messagequeue/client"
	"lemur/messagequeue/server"
)

var (
	isClient   bool
	isConsumer bool
	queueName  string
)

func randomString(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(randomBytes)[:length]
}

func main() {
	flag.BoolVar(&isClient, "client", false, "is client")
	flag.BoolVar(&isConsumer, "consumer", true, "consumer mode")
	flag.StringVar(&queueName, "queue", "default", "queue name")
	flag.Parse()

	if isClient {
		if isConsumer {
			client := client.NewSubscriber(":3000", queueName)
			client.ReadFromQueue()
		}

		if !isConsumer {
			client := client.NewPublisher(":3000", queueName)

			for {
				time.Sleep(time.Millisecond * 20)
				client.PublishMessage(randomString(10))
			}
		}
	} else {
		s := server.NewServer()
		s.Start()
	}

}
