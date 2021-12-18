package services

import (
	models "github.com/imokenpi2011/fotune-slipper/server/app/models/fotune"
	"github.com/imokenpi2011/fotune-slipper/server/app/repository/calculation"
	repository "github.com/imokenpi2011/fotune-slipper/server/app/repository/firebase"
)

/**
 * おみくじ結果を取得する.
 *
 * @return fotune おみくじ結果
 * @return err エラー
 */
func GetFotuneResult() (fotune models.Fotune, err error) {

	// おみくじ番号を生成
	fotuneId, err := calculation.GenerateFotuneNumber()
	if err != nil {
		return fotune, err
	}

	// おみくじ番号を元に結果を取得する
	fotune, err = repository.GetFotunesById(fotuneId)

	return fotune, err
}
