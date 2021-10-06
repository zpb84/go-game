package text_view

import (
	"fmt"
	"strings"

	"github.com/zpb84/go-game/lib/game"
)

const (
	cols = "ABCDEFGHJKLMNOPQRST"
)

var (
	stoneToChar = map[game.Color]rune{
		game.NONE:  '🔲',
		game.BLACK: '🌚',
		game.WHITE: '🌝',
	}
)

func PrintMove(player game.Color, move *game.Move) {
	strMove := ""
	switch {
	case move.IsPass():
		strMove = "passes"
	case move.IsResign():
		strMove = "resigns"
	default:
		strMove = fmt.Sprintf("%v%v", cols[move.Point().Col-1], move.Point().Row)
	}
	fmt.Printf("%s %s\n", player, strMove)
}

func PrintBoard(board *game.Board) {
	var builder strings.Builder
	for row := board.Rows(); row > 0; row-- {
		builder.Reset()
		for col := 1; col <= board.Columns(); col++ {
			stone := board.GetColor(game.Point{
				Row: row,
				Col: col,
			})
			_, _ = builder.WriteRune(stoneToChar[stone])
		}
		fmt.Printf("%s\n", builder.String())
	}
}
