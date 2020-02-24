package server

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"
	"github.com/MSU-Bot/Serverless/common/serverutils"
	log "github.com/sirupsen/logrus"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue FirestoreValue `json:"oldValue"`
	Value    FirestoreValue `json:"value"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	Fields interface{} `json:"fields"`
}

// HelloFirestore is triggered by a change to a Firestore document.
func HelloFirestore(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function triggered by change to: %v", meta.Resource)
	log.Printf("%v", e)
	return nil
}

// WelcomeUserHandler sends the user a welcome text to MSUBot.
// Triggers on Firestore user document create
func WelcomeUserHandler(ctx context.Context, e FirestoreEvent) {
	client := http.DefaultClient
	log.WithContext(ctx).Infof("Context loaded. Starting execution.")

	firebaseClient := serverutils.GetFirebaseClient(ctx)
	defer firebaseClient.Close()

	phNum := "TODOFIX"

	userData, uid := serverutils.FetchUserDataWithNumber(ctx, firebaseClient, phNum)
	if userData == nil {
		log.WithContext(ctx).Errorf("User doesn't exist in the database. Userdata: %v", userData)
		return
	}
	welcomeSent, ok := userData["welcomeSent"].(bool)
	if !ok {
		log.WithContext(ctx).Infof("welcomeSent: %v", welcomeSent)
	}
	if welcomeSent {
		log.WithContext(ctx).Infof("Already welcomed user")
		return
	}
	messageText := fmt.Sprintf("Thanks for signing up for MSUBot! We'll text you from this number when a seat opens up. Go Cats!")
	_, err := serverutils.SendText(client, userData["number"].(string), messageText)
	if err != nil {
		log.WithContext(ctx).Errorf("Could not send text to user!")

		return
	}
	firebaseClient.Collection("users").Doc(uid).Set(ctx, map[string]interface{}{
		"welcomeSent": true,
	}, firestore.MergeAll)

}
