package driver

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// Firebase ...
type Firebase struct {
	Client *auth.Client
}

var ctx = context.Background()
var firebaseConn = &Firebase{}

//Init creates a firebase instance from a service accoutn file
func Init() error {
	opt := option.WithCredentialsFile("service-acc.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return err
	}

	firebaseConn.Client = client

	return err
}

//CreateCustomToken creates a new token from a token received as a parameter
func CreateCustomToken(token string) (string, error) {
	newToken, err := firebaseConn.Client.CustomToken(ctx, token)
	if err != nil {
		fmt.Println("Failed to create token:", err.Error())
	}

	return newToken, err
}
