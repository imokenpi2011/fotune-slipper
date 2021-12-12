package controllers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Firestoreのクライアント
var Client *firestore.Client
var Ctx context.Context

/**
* 初期処理
 */
func init() {
	// firestoreクライアントの作成
	createClient()
}

/**
* Firebase関連の初期設定をする。
 */
func createClient() {
	log.Println("Start authenticate firebase ")
	// 認証情報を定義
	Ctx = context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")

	// firebaseで認証する
	app, err := firebase.NewApp(Ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	// クライアントの作成
	Client, err := app.Firestore(Ctx)
	if err != nil {
		log.Fatalln(err)
	} else if Client == nil {
		log.Fatalln("Failed to reate firestore client.")
	}
}
