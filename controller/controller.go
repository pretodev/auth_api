package controller

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"main/data"
)

type Controllers struct{}

var db data.Database

func init() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase_key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	db = data.NewDatabase(ctx, client)
}
