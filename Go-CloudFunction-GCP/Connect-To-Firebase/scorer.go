package scorer

import (
	"context"
	"log"
	"time"

	"firebase.google.com/go/db"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json::"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue struct
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your database.
	// Log an interface{} value and inspect the result to ses a JSON
	// representation of your database fields.
	Fields     Review    `json:"fields"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"updateTime"`
}

// Review reprsents the Firestore schema of a movie review.
type Review struct {
	Author struct {
		Value string `json:"stringValue`
	} `json:"author"`
	Text struct {
		Value string `json:"stringValue"`
	} `json:"text"`
}

var client *db.Client

func init() {
	ctx := context.Background()
	conf := &firebase.Config{
		// https://<PROJECT_ID>.firebaseio.com/
		DatabaseURL: "https://connect-cf-241923.firebaseio.com",
	}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}
	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

// ScoreReview generates the scores for movie reviews and transactionally writes them to the Firebase Realtime Database.
func ScoreReview(ctx context.Context, e FirestoreEvent) error {
	docID := e.Value.Name
	log.println(docID)

	review := e.Value.Fields

	reviewScore := score(review.Text.Value)

	ref := client.NewRef("scores").Child(review.Author.Value)
	updateTxn := func(node db.TransactionNode) (interface{}, error) {
		var currentScore int
		if err := node.Unmarshal(&currentScore); err != nil {
			return nil, err
		}
		return currentScore + reviewScore, nil
	}
	return ref.Transaction(ctx, updateTxn)
}

// score computes the score for a review text.
//
// This is currently just the length of the text.
func score(text string) int {
	return len(text)
}
