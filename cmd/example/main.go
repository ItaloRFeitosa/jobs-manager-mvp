package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	redisClient "github.com/italorfeitosa/jobs-manager-mvp/manager/infrastructure/redis"
	"github.com/tidwall/gjson"
)

func main() {
	client := redisClient.Get()
	channel := "test_channel"
	pubsub := client.Subscribe(context.Background(), channel)
	defer pubsub.Close()

	controlCh := pubsub.Channel()

	for msg := range controlCh {
		go processMessage(context.TODO(), msg)
	}
}

func processMessage(ctx context.Context, msg *redis.Message) {
	client := redisClient.Get()
	idkey := gjson.Get(msg.Payload, "idempotencyKey")
	ackChannel := fmt.Sprintf("%s:%s:ack", msg.Channel, idkey.String())
	err := client.Publish(ctx, ackChannel, `{"started": "started message"}`).Err()
	log.Printf("ackChannel: %#+v\n", ackChannel)
	if err != nil {
		log.Printf("err: %#+v\n", err)
	}

	log.Printf("processing job %s ...", idkey)
	time.Sleep(5 * time.Second)
	err = client.Publish(ctx, ackChannel, `{"succeeded": "success message"}`).Err()

	if err != nil {
		log.Printf("err: %#+v\n", err)
	}
}
