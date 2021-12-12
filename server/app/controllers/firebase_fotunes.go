package controllers

import (
	"context"
	"fmt"
	"log"
)

// コレクション名
var collectionPath string = "fotunes"

/**
 * firestoreのfotunesにデータを挿入する.
 *
 * @params ID おみくじ:番号
 * @params luck おみくじ:運勢
 * @params wish おみくじ:願望
 * @params study おみくじ:学業
 * @params love おみくじ:恋愛
 * @params health おみくじ:健康
 * @params waiting おみくじ:待ち人
 * @return err エラー
 */
func InsertFotunes(ID string, luck string, wish string, study string, love string, health string, waiting string) (err error) {
	log.Println(fmt.Sprintf("Start insertFotunes. Fotune:[ID:%s,wish:%s ,luck:%s,study:%s,love:%s,health:%s,waiting:%s]", ID, wish, luck, study, love, health, waiting))
	ctx := context.Background()
	client := createClient(ctx)

	_, err = client.Collection(collectionPath).Doc(ID).Set(ctx, map[string]interface{}{
		"wish":    wish,
		"luck":    luck,
		"study":   study,
		"love":    love,
		"health":  health,
		"waiting": waiting,
	})

	log.Println("End insertFotunes.")

	return err

}
