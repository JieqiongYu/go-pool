package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"firebase.google.com/go"
	"fmt"
	"io"
	"strconv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"github.com/golang/glog"
)

var userID int64

type externalid struct {
	ExternalID string `firestore:"external_id"`
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
func write(ctx context.Context, client *firestore.Client, userID int64, e externalid) {
	result, err := client.Collection("externalid").Doc(strconv.FormatInt(userID,10)).Set(ctx, e)
	if err != nil {
		glog.Errorln(err)
	}
	// docRef.ID is the ID created automatically when adding record
	fmt.Println(result)
}

func encrypt(userID int64) string {

	text := []byte(string(userID))
	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM handle them
	if err != nil {
		fmt.Println(err)
	}

	// Creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all time,
	// for a given key.
	return base64.StdEncoding.EncodeToString([]byte(gcm.Seal(nonce, nonce, text, nil)))

	// // the WriteFile method returns an error if unsuccessful
	// err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	// // handle this error
	// if err != nil {
	// 	// print it out
	// 	fmt.Println(err)
	// }
}
func main() {
	ctx := context.Background()
	app := initFirebase()
	client, err := app.Firestore(ctx)
	if err != nil {
		glog.Errorln(err)
	}

	defer client.Close()

	for userID = 999999892; userID >= 999999862; userID-- {
		pushRequest := externalid{
			ExternalID: encrypt(userID),
		}
		write(ctx, client, userID, pushRequest)
	}
}
