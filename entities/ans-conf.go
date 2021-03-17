package entities

// AnsConf - 回答ファイル
type AnsConf struct {
	Quest Quest
	Ans   Ans
	Trick Trick
}

// Ans - [Ans] テーブル
type Ans struct {
	Bestmove string
}
