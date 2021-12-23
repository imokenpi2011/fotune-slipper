package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/imokenpi2011/fotune-slipper/server/config"
)

/**
 * レスポンス用のJSONを生成する
 *
 * @param w レスポンス
 * @param code レスポンスコード
 * @param payload JSON生成用のパラメータ
 */
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/**
 * サーバーを起動する
 */
func StartMainServer() error {
	// URLの登録
	http.HandleFunc("/fotune", doFotuneSlipper)

	//サーバを起動する。(起動ポート以外,関係ないページの場合404を返す様にする)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
