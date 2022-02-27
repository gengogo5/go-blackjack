package player

import (
	"blackjack/internal/card"
	"strings"
)

// プレイヤー or ディーラー
type Player struct {
	hands    []card.Card
	name     string
	isDealer bool
}

// プレイヤーを生成するファクトリ関数
func InitPlayer() *Player {
	player := Player{
		name:     "あなた",
		isDealer: false,
	}
	return &player
}

// ディーラーを生成するファクトリ関数
// 2回呼ばないで欲しい
func InitDealer() *Player {
	dealer := Player{
		name:     "ディーラー",
		isDealer: true,
	}
	return &dealer
}

// デッキからカードを1枚引く
func (p *Player) Hit(deck *card.Deck) card.Card {
	card := deck.Pop()
	p.hands = append(p.hands, card)
	return card
}

// ハンドの合計値を返す
func (p Player) HandsNumber() int {
	sum := 0
	ace := 0
	for _, v := range p.hands {
		sum += v.Score()
		// Aの場合は合計に対して例外処理を行う
		if v.Score() == 11 {
			ace++
		}
	}

	// ハンドが21を超える場合はAを1と数える
	// Aの枚数分行う
	for i := ace; 0 < i && 21 < sum; i-- {
		sum -= 10
	}
	return sum
}

// ハンドの文字列表現を返す
func (p Player) HandsString() string {
	cards := make([]string, len(p.hands))
	for i, v := range p.hands {
		cards[i] = v.String()
	}
	return strings.Join(cards, " ")
}

// ハンドがブラック・ジャックかを返す
func (p Player) HasBJ() bool {
	// ハンドが2枚で21の場合はブラックジャック
	return len(p.hands) == 2 && p.HandsNumber() == 21
}

// ハンドがBUSTかを返す
func (p Player) IsBust() bool {
	return p.HandsNumber() > 21
}

// 名前を返す
// 構造体のメンバを開くと外から書き換えられるのでReaderのアクセサを用意した
func (p Player) Name() string {
	return p.name
}
