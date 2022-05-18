package data

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"main/models"
)

type Database struct {
	ctx    context.Context
	client *firestore.Client
}

func NewDatabase(ctx context.Context, client *firestore.Client) Database {
	return Database{ctx, client}
}

func (db Database) ExistsUsername(username string) bool {
	query := db.client.Collection("users").Where("Username", "==", username).Documents(db.ctx)
	docs, _ := query.GetAll()
	return len(docs) > 0
}

func (db Database) ExistsId(id string) bool {
	doc, _ := db.client.Collection("users").Doc(id).Get(db.ctx)
	return doc.Exists()
}

func (db Database) SaveUser(user models.EditingUser) {
	ref := db.client.Collection("users").NewDoc()
	_, err := ref.Set(db.ctx, user)
	if err != nil {
		panic(err)
	}
}

func (db Database) CheckCredential(credential models.Credentials) (string, error) {
	query := db.client.Collection("users").Where("Username", "==", credential.UserName).Documents(db.ctx)
	docs, _ := query.GetAll()
	if len(docs) == 0 {
		return "", fmt.Errorf("user not found")
	}
	password := docs[0].Data()["Password"].(string)
	if password != credential.Password {
		return "", fmt.Errorf("wrong password")
	}
	return docs[0].Ref.ID, nil
}

func (db Database) GetProfile(id string) (*models.Profile, error) {
	doc, _ := db.client.Collection("profiles").Doc(id).Get(db.ctx)
	if !doc.Exists() {
		return nil, fmt.Errorf("user not found")
	}
	data := doc.Data()
	profile := models.Profile{
		FullName: data["FullName"].(string),
		Email:    data["Email"].(string),
		Github:   data["Github"].(string),
		Linkedin: data["Linkedin"].(string),
		Whatsapp: data["Whatsapp"].(string),
	}
	return &profile, nil
}

func (db Database) SaveProfile(userId string, profile models.Profile) {
	ref := db.client.Collection("profiles").Doc(userId)
	_, err := ref.Set(db.ctx, profile)
	if err != nil {
		panic(err)
	}
}
