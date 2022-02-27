package game

// ゲームの勝敗を表す疑似enum型
type Result int

const (
	Win Result = iota
	Lose
	Push
)

func (r Result) String() string {
	switch r {
	case Win:
		return "勝ち"
	case Lose:
		return "負け"
	case Push:
		return "引き分け"
	default:
		// ちょっとかっこ悪いけどswitch使いたかった
		return ""
	}
}
