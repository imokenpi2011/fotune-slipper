package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	models "github.com/imokenpi2011/fotune-slipper/server/app/models/fotune"
)

// コレクション名
var collectionPath string = "fotunes"

/**
 * firestoreのfotunesにデータを挿入する.
 *
 * @params fotune おみくじクラス
 * @return err エラー
 */
func InsertFotunes(fotune models.Fotune) (err error) {
	log.Println(fmt.Sprintf("Start insertFotunes. Fotune:[ID:%s,wish:%s ,luck:%s,study:%s,love:%s,health:%s,waiting:%s]",
		fotune.ID, fotune.Wish, fotune.Luck, fotune.Study, fotune.Love, fotune.Health, fotune.Waiting))
	ctx := context.Background()
	client := createClient(ctx)

	_, err = client.Collection(collectionPath).Doc(fotune.ID).Set(ctx, map[string]interface{}{
		"wish":    fotune.Wish,
		"luck":    fotune.Luck,
		"study":   fotune.Study,
		"love":    fotune.Love,
		"health":  fotune.Health,
		"waiting": fotune.Waiting,
	})

	log.Println("End insertFotunes.")

	return err

}

/**
 * firestoreのfotunesの件数を取得する
 *
 * @return count 件数
 */
func GetFotunesCount() (count int, err error) {
	log.Println("Start getFotunesCount.")

	// クライアントの取得
	ctx := context.Background()
	client := createClient(ctx)

	// 全件取得する(対象が100件想定のため使用)
	fotunes, err := client.Collection(collectionPath).DocumentRefs(ctx).GetAll()
	if err != nil {
		return count, err
	}
	// 件数のみを返却する
	count = len(fotunes)

	log.Println("End getFotunesCount. COUNT:", count)

	return count, err

}

/**
 * IDに応じてfirestoreのfotunesを取得する
 *
 * @param ID おみくじ番号
 * @return count 件数
 */
func GetFotunesById(fotuneId int) (fotune models.Fotune, err error) {

	log.Println("Start getFotunesById from firestore.")

	// IDが不正の場合はエラー
	if fotuneId == 0 {
		return fotune, errors.New("Invalid fotuneId is specified.")
	}

	// idを文字列に変換する
	strFotuneId := strconv.Itoa(fotuneId)
	if err != nil {
		return fotune, err
	}

	// クライアントの取得
	ctx := context.Background()
	client := createClient(ctx)

	// 占い結果の取得
	res, err := client.Collection(collectionPath).Doc(strFotuneId).Get(ctx)
	if err != nil {
		return fotune, err
	}

	// 占い結果に一つでも不正文字がある場合はエラー
	data := res.Data()
	for key, value := range data {
		log.Printf("%v: %v\n", key, value)
		if (value == nil) || value == "" {
			return fotune, errors.New("Contains empty value. at this key:" + key)
		}
	}

	// 結果の代入
	luck, _ := res.DataAt("luck")
	wish, _ := res.DataAt("wish")
	study, _ := res.DataAt("study")
	love, _ := res.DataAt("love")
	health, _ := res.DataAt("health")
	waiting, _ := res.DataAt("waiting")
	fotune = models.Fotune{
		ID:      strFotuneId,
		Luck:    luck.(string),
		Wish:    wish.(string),
		Study:   study.(string),
		Love:    love.(string),
		Health:  health.(string),
		Waiting: waiting.(string),
	}

	log.Println(fmt.Sprintf("Successful getting fotunes. Fotune:[ID:%s,wish:%s ,luck:%s,study:%s,love:%s,health:%s,waiting:%s]",
		fotune.ID, fotune.Wish, fotune.Luck, fotune.Study, fotune.Love, fotune.Health, fotune.Waiting))

	log.Println("End getFotunesById from firestore.")

	return fotune, err
}
