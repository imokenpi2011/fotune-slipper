package calculation

import (
	"errors"
	"log"
	"math/rand"

	repository "github.com/imokenpi2011/fotune-slipper/server/app/repository/firebase"
)

/**
 * おみくじ番号を取得する
 *
 * @return fotuneNum おみくじ番号
 * @return err エラー
 */
func GenerateFotuneNumber() (fotuneId int, err error) {
	log.Println("Start generateFotuneNumber.")

	// firestoreのfotunenの登録件数を取得
	totalFotunesNum, err := repository.GetFotunesCount()
	if err != nil {
		// 取得が失敗した場合はエラー
		return fotuneId, err
	} else if totalFotunesNum <= 0 {
		// 取得結果が0件の場合はエラー
		return fotuneId, errors.New("Fotunes data was not found.")
	}

	// 1からfirestoreの総件数までのランダムな数を返す(0の場合は再抽選する)
	for fotuneId == 0 {
		fotuneId = rand.Intn(totalFotunesNum)
	}

	log.Println("End generateFotuneNumber. fotuneId:", fotuneId)

	return fotuneId, err
}
