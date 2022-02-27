package game

import (
	"blackjack/internal/card"
	"blackjack/internal/player"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Game struct {
	deck   card.Deck
	player player.Player
	dealer player.Player
}

func InitGame() *Game {
	game := Game{}
	game.deck = *card.InitDeck()
	game.player = *player.InitPlayer()
	game.dealer = *player.InitDealer()
	return &game
}

func (g *Game) Play() {
	// デッキをシャッフルする
	g.deck.Suffule()

	// 2枚ずつ配る
	g.dealer.Hit(&g.deck)
	g.dealer.Hit(&g.deck)
	g.player.Hit(&g.deck)
	g.player.Hit(&g.deck)

	// プレイヤーのターン
	// BUSTするまでプレイヤーはカードを引く選択をする
	g.playerTurn()

	// ディーラーのターン
	// プレイヤーがBUSTした場合は勝敗判定へ移る
	if !g.player.IsBust() {
		g.dealerTurn()
	}

	// 勝敗判定
	result, msg := g.judge()
	g.printResult(result, msg)
}

// YesかNoをユーザに問う
func askYorN() string {
	print("(Y/n) >")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "y", "Y":
			return "Y"
		case "n", "N":
			return "n"
		default:
			println("Yかnを入力してください")
		}
	}
}

// ゲームの結果を返す
func (g Game) judge() (Result, string) {
	switch {
	case g.player.IsBust():
		return Lose, g.player.Name() + "のBUSTです"
	case g.dealer.IsBust():
		return Win, g.dealer.Name() + "のBUSTです"
	case g.player.HasBJ() && !g.dealer.HasBJ():
		return Win, g.player.Name() + "のBlackJackです"
	case !g.player.HasBJ() && g.dealer.HasBJ():
		return Lose, g.dealer.Name() + "のBlackJackです"
	case g.player.HasBJ() && g.dealer.HasBJ():
		return Push, "BlackJack同士は引き分けです"
	case g.dealer.HandsNumber() > g.player.HandsNumber():
		return Lose, g.dealer.Name() + "のハンドが上です"
	case g.player.HandsNumber() > g.dealer.HandsNumber():
		return Win, g.player.Name() + "のハンドが上です"
	default:
		return Push, "同点です"
	}
}

// プレイヤーがカードを引く
func (g *Game) playerTurn() {
	delayPrintln(fmt.Sprintf("<<<<< %sのターン >>>>>", g.player.Name()))
	printUpCard(g.dealer)
	// BUSTするまでカードを引く選択を繰り返す
	for !g.player.IsBust() {
		printHands(g.player)
		println()
		print("1枚引きますか？")

		if askYorN() == "Y" {
			card := g.player.Hit(&g.deck)
			println("引いたカード: " + card.String())
		} else {
			break
		}
	}
}

// ディーラーがカードを引く
func (g *Game) dealerTurn() {
	delayPrintln(fmt.Sprintf("<<<<< %sのターン >>>>>", g.dealer.Name()))
	printHands(g.dealer)
	// 17以上になるまで引き続ける
	for g.dealer.HandsNumber() <= 16 {
		card := g.dealer.Hit(&g.deck)
		delayPrintln("引いたカード: " + card.String())
	}
}

// プレイヤーのハンドと数値を標準出力する
func printHands(p player.Player) {
	delayPrintln(fmt.Sprintf("%s >> %s (%d)", p.Name(), p.HandsString(), p.HandsNumber()))
}

// ハンドの1枚目だけを標準出力する
// ディーラー用
func printUpCard(p player.Player) {
	maskHand := strings.Split(p.HandsString(), " ")[0] + " ??"
	println(fmt.Sprintf("%s >> %s", p.Name(), maskHand))
}

// 勝負の結果を標準出力する
func (g Game) printResult(result Result, message string) {
	delayPrintln("==========================")
	printHands(g.dealer)
	printHands(g.player)
	delayPrintln("結果: " + result.String())
	delayPrintln(message)
}

// 少し待ってから表示する
func delayPrintln(str string) {
	time.Sleep(time.Millisecond * 500)
	println(str)
}
