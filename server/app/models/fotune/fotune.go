package models

// おみくじクラス
type Fotune struct {
	ID      string // ID
	Luck    string // 運勢
	Wish    string // 願望
	Study   string // 学業
	Love    string // 恋愛
	Health  string // 健康
	Waiting string // 待ち人
}

//レスポンス
type FotuneResponse struct {
	Status     int    // 結果コード
	ErrMessage string // エラーメッセージ
	Fotune     Fotune // おみくじ結果
}
