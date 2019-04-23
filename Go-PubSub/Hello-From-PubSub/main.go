package pubsub

import (
	"context"
	"encoding/json"
	"log"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Info struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

// Pub/Subからメッセージ受信したらこの関数が起動する
func HelloFromPubSub(ctx context.Context, m PubSubMessage) error {
	var i Info

	err := json.Unmarshal(m.Data, &i)

	if err != nil {
		log.Printf("Error:%T message:%v", err, err)
		return nil
	}

	log.Println("My name is", i.Name, "and my hobby is", i.Hobby)
	return nil
}
