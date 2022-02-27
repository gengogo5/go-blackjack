# go-blackjack
Golang学習用にブラックジャックを実装

# 仕様
## 実装済
- 標準的なブラックジャックのルール
- プレイヤーとディーラーの1対1
- Aの特殊ルール(1または11どちらか有利な方を選択できる)

## 未実装
- スプリット
- チップやベットなどの賭け要素
- 勝敗の履歴記録
- 複数プレイヤー

# 構成
## パッケージ
- cmd
  - br/main
- internal
  - card
  - game
  - player

## 型
### Card
カード1枚を表すユーザ定義型(int)

自分のスート(マーク)や数字、ブラックジャックのルール上での数字を知っている

### Deck
カードの集合を表す構造体

メンバに`Card`のスライスを持つ。ユーザ定義型でも良かったかもしれない

### Player
ブラックジャックのプレイヤー(ディーラー含む)を表す構造体

`Card`のスライスをハンドとしてメンバに持つ

自身のハンドに関する点数計算やルールを知っている

### Game
デッキとプレイヤーを持つ最上位の構造体

ゲームの進行や勝敗判定を行う

### Result
勝敗の構造体

enum的なものを作ってみたかったので作った

# 実行ファイル作成

```bash
go build -o ./blackjack cmd/bj/main.go
```
