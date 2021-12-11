package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func initContext() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Collection("fotunes").Doc("2").Set(ctx, map[string]interface{}{
		"name":  "Update",
		"age":   11,
		"email": "first@example.com",
	})
	defer client.Close()
}

func getClient() (client *firestore.Client) {
	// Use a service account
	ctx := context.Background()
	app := getAppContext(ctx)

	// クライアントの取得
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func getAppContext(ctx context.Context) (app *firebase.App) {
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	return app
}
