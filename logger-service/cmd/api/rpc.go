package main

import (
	"context"
	"log"
	"time"

	"github.com/jhonipereira/go-micro-logger/data"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo: ", err)
		return err
	}

	*resp = "Proccessed payload " + payload.Name
	return nil
}
