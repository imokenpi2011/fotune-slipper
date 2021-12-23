package controllers

import (
	"log"
	"net/http"

	models "github.com/imokenpi2011/fotune-slipper/server/app/models/fotune"

	services "github.com/imokenpi2011/fotune-slipper/server/app/services/fotune"
)

// 結果コード:成功
var SUCCESS_CODE = 0

// 結果コード:失敗
var FAILED_CODE = 1

// エラーメッセージ
var ERROR_MESSAGE = "おみくじに失敗しました。しばらく時間を置いて再度実施してください。"

/**
 * IDに応じてfirestoreのfotunesを取得する
 *
 * @param w レスポンス
 * @param r リクエスト
 */
func doFotuneSlipper(w http.ResponseWriter, r *http.Request) {
	log.Println("Start getting fotunne.")

	// おみくじの取得
	fotune, err := services.GetFotuneResult()

	// おみくじ結果の作成
	fotuneResponse := models.FotuneResponse{
		Fotune: fotune,
	}

	if err != nil {
		log.Println("Failed to get fotunne.")
		log.Println("ERROR:", err)
		// 結果コードの設定
		fotuneResponse.Status = FAILED_CODE
		// エラーメッセージの設定
		fotuneResponse.ErrMessage = ERROR_MESSAGE

		// レスポンスを返す
		respondWithJson(w, http.StatusBadRequest, fotuneResponse)
		return
	}
	// 結果コードの設定
	fotuneResponse.Status = SUCCESS_CODE

	log.Println("Successful getting fotunne.")

	//レスポンスを返す
	respondWithJson(w, http.StatusOK, fotuneResponse)
}
