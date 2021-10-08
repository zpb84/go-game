package main

import (
	"flag"
	"log"
	"os"

	"github.com/zpb84/go-game/lib"
	"github.com/zpb84/go-game/lib/terminal"
)

const (
	boardSize = 9
)

func main() {
	flagBots := flag.Bool("bots", false, "Battle bots")
	flagInfo := flag.Bool("info", false, "Show info")
	flagPlayers := flag.Bool("players", false, "Game for two players")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	var err error
	switch {
	case *flagInfo:
		lib.ShowInfo(os.Stdout)
	case *flagBots:
		err = terminal.BattleNewbies(boardSize)
	case *flagPlayers:
		err = terminal.BattlePlayers(boardSize, logger)
	default:
		err = terminal.BattleWithNewbie(boardSize, logger)
	}
	if err != nil {
		logger.Fatal(err)
	}
}
