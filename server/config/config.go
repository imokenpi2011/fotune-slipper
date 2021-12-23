package config

import (
	"log"

	"github.com/imokenpi2011/fotune-slipper/server/utils"
	"gopkg.in/go-ini/ini.v1"
)

// 設定一覧を指定
type ConfigList struct {
	Port    string //ポート番号
	LogFile string //ログファイル名
}

//Configインスタンスを定義
var Config ConfigList

//初期処理
func init() {
	//設定の読み込み処理
	LoadConfig()
	//ログの読み込みを設定
	utils.LoggingSettings(Config.LogFile)
}

//設定を読み込む
func LoadConfig() {
	// 設定を記載したiniファイルを読み込む
	configDir := "config.ini"
	//root直下の設定ファイルを読み込む
	cfg, err := ini.Load(configDir)
	if err != nil {
		log.Fatalln(err)
	}

	//読み込んだ値を設定する
	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").MustString("8080"),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}
}
