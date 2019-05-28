package main

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"github.com/golang/glog"
)

type history struct {
	UserID     int64     `firestore:"user_id"`
	Message    string    `firestore:"message"`
	SendStatus int       `firestore:"send_status"`
	Created    time.Time `firestore:"created"`
	Updated    time.Time `firestore:"updated"`
}

func initFirebase() *firebase.App {
	opt := option.WithCredentialsFile("/Users/jieqiong-yu/Desktop/connect-cf.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		glog.Errorln("Error")
	}
	return app
}

// Write
func write(ctx context.Context, client *firestore.Client, h history) {
	docRef, _, err := client.Collection("history").Add(ctx, h)
	if err != nil {
		glog.Errorln(err)
	}
	// docRef.ID is the ID created automatically when adding record
	fmt.Println(docRef.ID)
}

func main() {
	ctx := context.Background()
	app := initFirebase()
	client, err := app.Firestore(ctx)
	if err != nil {
		glog.Errorln(err)
	}

	defer client.Close()

	pushRequest := history{
		UserID:     1,
		Message:    "Push Message 1",
		SendStatus: 0,
		Created:    time.Now(),
		Updated:    time.Now(),
	}

	write(ctx, client, pushRequest)

}
