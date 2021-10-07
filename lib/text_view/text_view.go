package text_view

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/zpb84/go-game/lib/game"
)

const (
	cols = "ABCDEFGHJKLMNOPQRSTVWXYZ"
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
		strMove = fmt.Sprintf("%s%v", string(cols[move.Point().Col-1]), move.Point().Row)
	}
	fmt.Printf("%s %s\n", player, strMove)
}

func getNamesColumns(count int) string {
	var builder strings.Builder
	for index, c := range cols[:count] {
		if index == 0 {
			builder.WriteString("\t")
		} else {
			builder.WriteString("\t\t")
		}
		builder.WriteRune(c)
	}
	return builder.String()
}

func PrintBoard(board *game.Board) {
	fmt.Print("\033[H\033[2J")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.StripEscape)
	var builder strings.Builder
	for row := board.Rows(); row > 0; row-- {
		builder.Reset()
		for col := 1; col <= board.Columns(); col++ {
			stone := board.GetColor(game.Point{
				Row: row,
				Col: col,
			})
			_, _ = builder.WriteRune('\t')
			_, _ = builder.WriteRune(stoneToChar[stone])
		}
		fmt.Fprintf(w, "%v%s\n", row, builder.String())
	}
	fmt.Fprintln(w, getNamesColumns(board.Columns()))
	w.Flush()
}
