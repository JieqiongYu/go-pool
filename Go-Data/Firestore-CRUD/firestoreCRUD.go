package main

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/golang/glog"
)

type history struct {
	UserID     int       `firestore:"user_id"`
	Message    string    `firestore:"message"`
	SendStatus int       `firestore:"send_status"`
	Created    time.Time `firestore:"created"`
	Updated    time.Time `firestore:"updated"`
}

func initFirebase() *firebase.App {
	opt := option.WithCredentialsFile("/Users/jieqiong-yu/Desktop/key-connect-cf.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		glog.Errorln(err)
	}
	return app
}

// Write
func write(ctx context.Context, client *firestore.Client, h history) {
	_, err := client.Collection("history").Doc("BJ").Set(ctx, h, firestore.MergeAll)
	if err != nil {
		glog.Errorln(err)
	}
	//// docRef.ID is the ID created automatically when adding record
	//fmt.Println(docRef.ID)
}

// Read
func read(ctx context.Context, client *firestore.Client) {
	fmt.Println("All unsend messages for reading")
	iter := client.Collection("history").Where("send_status", "==", 12).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			glog.Errorln(err)
		}
		fmt.Println(doc.Data())
	}
}

//Update
func update(ctx context.Context, client *firestore.Client) {
	fmt.Println("All unsend messages for update")
	iter := client.Collection("history").Where("send_status", "==", 1).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			glog.Errorln(err)
		}
		fmt.Println(doc.Ref.ID)
		client.Collection("history").Doc(doc.Ref.ID).Update(ctx, []firestore.Update{{Path: "send_status", Value: 12}, {Path: "message", Value: "hahaha"}})
	}
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
		UserID:     2,
		Message:    "Push Message 3",
		SendStatus: 2,
		Created:    time.Now(),
		Updated:    time.Now(),
	}

	write(ctx, client, pushRequest)
	//read(ctx, client)
	update(ctx, client)
}
