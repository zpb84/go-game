package main

import (
	"fmt"
	"time"

	"github.com/zpb84/go-game/lib/game"
	"github.com/zpb84/go-game/lib/game/agent"
	"github.com/zpb84/go-game/lib/interplay"
	"github.com/zpb84/go-game/lib/text_view"
)

func main() {
	boardSize := 9
	g := game.NewGame(boardSize)
	bots := map[game.Color]interplay.Agent{
		game.BLACK: agent.NewNewbie(),
		game.WHITE: agent.NewNewbie(),
	}
	var err error
	for !g.IsOver() {
		time.Sleep(6 * time.Millisecond)
		text_view.PrintBoard(g.Board())
		botMove := bots[g.NextPlayer()].SelectMove(g)
		text_view.PrintMove(g.NextPlayer(), botMove)
		if g, err = g.ApplyMove(botMove); err != nil {
			fmt.Printf("ERROR: %v", err)
			return
		}
	}
}
