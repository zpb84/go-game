package terminal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/zpb84/go-game/lib/game"
	"github.com/zpb84/go-game/lib/game/agent"
	"github.com/zpb84/go-game/lib/interplay"
)

const (
	cols = "ABCDEFGHJKLMNOPQRSTVWXYZ"
)

var (
	stoneToChar = map[game.Color]rune{
		game.NONE:  ' ',
		game.BLACK: 'X',
		game.WHITE: 'O',
	}
)

func BattlePlayers(boardSize int, l *log.Logger) error {
	var err error
	if err = checkBoardSize(boardSize); err != nil {
		return err
	}
	g := game.NewGame(boardSize)
	reader := bufio.NewReader(os.Stdin)
	var move *game.Move
	for !g.IsOver() {
		printBoard(g.Board())
		if g.NextPlayer() == game.BLACK {
			fmt.Print("BLACK Enter: ")
		} else {
			fmt.Print("WHITE Enter: ")
		}
		text, err := reader.ReadString('\n')
		if err != nil {
			l.Printf("Read string: %v\n", err)
			continue
		}
		p, err := pointFromString(text)
		if err != nil {
			l.Printf("Get point: %v\n", err)
			continue
		}
		move = game.Play(p)
		if g, err = g.ApplyMove(move); err != nil {
			l.Printf("Apply move: %v", err)
		}
	}
	return nil
}

func BattleWithNewbie(boardSize int, l *log.Logger) error {
	var err error
	if err = checkBoardSize(boardSize); err != nil {
		return err
	}
	var move *game.Move
	g := game.NewGame(boardSize)
	bot := agent.NewNewbie()
	reader := bufio.NewReader(os.Stdin)
	for !g.IsOver() {
		printBoard(g.Board())
		if g.NextPlayer() == game.BLACK {
			fmt.Print("Enter: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				l.Printf("Read string: %v\n", err)
				continue
			}
			p, err := pointFromString(text)
			if err != nil {
				l.Printf("Get point: %v\n", err)
				continue
			}
			move = game.Play(p)
		} else {
			move = bot.SelectMove(g)
		}
		printMove(g.NextPlayer(), move)
		if g, err = g.ApplyMove(move); err != nil {
			l.Printf("Apply move: %v", err)
		}
	}
	return nil
}

func BattleNewbies(boardSize int) error {
	if err := checkBoardSize(boardSize); err != nil {
		return err
	}
	g := game.NewGame(boardSize)
	bots := map[game.Color]interplay.Agent{
		game.BLACK: agent.NewNewbie(),
		game.WHITE: agent.NewNewbie(),
	}
	var err error
	for !g.IsOver() {
		time.Sleep(1 * time.Microsecond)
		printBoard(g.Board())
		botMove := bots[g.NextPlayer()].SelectMove(g)
		printMove(g.NextPlayer(), botMove)
		if g, err = g.ApplyMove(botMove); err != nil {
			return err
		}
	}
	return nil
}

func checkBoardSize(size int) error {
	if size <= 0 || size >= len(cols) {
		return fmt.Errorf("wrong board size, it must be greater than 0 and less than %v", len(cols))
	}
	return nil
}

func printMove(player game.Color, move *game.Move) {
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

func printBoard(board *game.Board) {
	fmt.Print("\033[H\033[2J")
	w := tabwriter.NewWriter(os.Stdout, 1, 0, 1, ' ', tabwriter.StripEscape)
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
	writeNamesColumns(w, board.Columns())
	w.Flush()
}

func writeNamesColumns(w io.Writer, count int) {
	writer := bufio.NewWriter(w)
	for _, c := range cols[:count] {
		_, _ = writer.WriteString("\t")
		_, _ = writer.WriteRune(c)
	}
	_, _ = writer.WriteString("\n")
	_ = writer.Flush()
}

func pointFromString(input string) (game.Point, error) {
	zeroPoint := game.Point{}
	if len(input) == 0 {
		return zeroPoint, ErrEmptyInput
	}
	col := strings.Index(cols, strings.ToUpper(input[0:1]))
	if col < 0 {
		return zeroPoint, fmt.Errorf("column name is incorrect")
	}
	row, err := strconv.Atoi(input[1 : len(input)-1])
	if err != nil {
		return zeroPoint, fmt.Errorf("number row is incorrect: %v", err)
	}
	return game.Point{
		Row: row,
		Col: col + 1,
	}, nil
}
