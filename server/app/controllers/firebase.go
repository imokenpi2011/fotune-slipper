package controllers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

/**
 * Firestoreのクライアントを取得する。
 *
 * @params ctx コンテキスト
 * @return client firestoreクライアント
 */
func createClient(ctx context.Context) (client *firestore.Client) {
	log.Println("Start authenticate firebase ")
	// 認証情報を定義
	//ctx = context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")

	// firebaseで認証する
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	// クライアントの作成
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	} else if client == nil {
		log.Fatalln("Failed to reate firestore client.")
	}
	log.Println("End authenticate firebase ")

	return client
}
