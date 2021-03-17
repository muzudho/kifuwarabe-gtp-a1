package entities

import (
	"github.com/muzudho/kifuwarabe-go-base/entities/phase"
)

// QuestConf - Tomlファイル
type QuestConf struct {
	Quest Quest
	Trick Trick
}

// Quest - [Quest] 区画
type Quest struct {
	BoardSize int8
	Step      int16
	Phase     phase.Phase
	BoardData string
}

// Trick - [Trick] テーブル
type Trick struct {
	Rows    []string
	Columns []int8
}

/*
// GetBoardArray - 盤上の石の色の配列。
func (config QuestConf) GetBoardArray() []int {
	nodes := removeAllWhiteSpace(config.Engine.BoardData)
	array := make([]int, len(nodes))
	for i, s := range nodes {
		switch s {
		case '.':
			array[i] = int(stone.None)
		case 'x':
			array[i] = int(stone.Black)
		case 'o':
			array[i] = int(stone.White)
		case '+':
			array[i] = int(stone.Wall)
		default:
			// Ignored.
		}
	}

	return array
}
*/

// BoardSize - 何路盤か
func (config QuestConf) BoardSize() int {
	return int(config.Quest.BoardSize)
}

// Step - 何手目か
func (config QuestConf) Step() int {
	return int(config.Quest.Step)
}

// Phase - 手番
//func (config QuestConf) Phase() Phase {
//	return int(config.Quest.Phase)
//}
