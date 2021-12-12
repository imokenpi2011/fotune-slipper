package models

import (
	"log"

	"github.com/imokenpi2011/fotune-slipper/server/app/controllers"
)

type Fotune struct {
	ID      string
	Luck    string
	Wish    string
	Study   string
	Love    string
	Health  string
	Waiting string
}

func RegistFotune() {
	fo := Fotune{
		ID:      "3",
		Luck:    "大吉",
		Wish:    "叶うでしょう",
		Study:   "そのまま励みましょう",
		Love:    "転機が訪れるでしょう",
		Health:  "食生活を見直しましょう",
		Waiting: "訪れるでしょう",
	}

	err := controllers.InsertFotunes(fo.ID, fo.Luck, fo.Wish, fo.Study, fo.Love, fo.Health, fo.Waiting)
	if err != nil {
		log.Fatalln(err)
	}

}
