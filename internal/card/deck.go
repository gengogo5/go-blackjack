package card

import (
	"math/rand"
	"time"
)

// カードデッキ
// InitDeck()で初期化
type Deck struct {
	cards []Card
}

// デッキの初期化用関数
func InitDeck() *Deck {
	deck := Deck{
		cards: make([]Card, 52),
	}

	for i := 0; i < 52; i++ {
		deck.cards[i] = Card(i)
	}
	return &deck
}

// デッキをシャッフルする
func (d *Deck) Suffule() {
	if d.HowMany() < 2 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

// デッキの一番上のカードを取り出す
func (d *Deck) Pop() Card {
	// デッキが枯れたケースは考慮しない
	c := d.cards[0]
	// 先頭の要素を削除。unshiftが欲しい
	d.cards = append(d.cards[:0], d.cards[1:]...)
	return c
}

// デッキの残り枚数を返す
func (d Deck) HowMany() int {
	return len(d.cards)
}
