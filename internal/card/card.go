package card

// カードを表す型
// 0〜51で1周する
// 52以上でも一応ゲームはできる。カジノだと2セット以上使う場合があるらしい
// 構造体でも良かったがユーザ定義型で十分かもしれないと思った
type Card int

var suitMark = [4]string{"♠", "♥", "◆", "♣"}
var displayNumber = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
var score = [13]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

// 表示用の文字列を返す
func (c Card) String() string {
	return suitMark[c.suitNumber()] + displayNumber[c.number()]
}

// 0〜12でカードの数を表す
func (c Card) number() int {
	return int(c) % 13
}

// 0〜3でスートを表す
func (c Card) suitNumber() int {
	return int(c) / 13
}

// ブラックジャックとして使われる数
func (c Card) Score() int {
	return score[c.number()]
}
