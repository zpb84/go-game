package agent

import (
	"math/rand"
	"time"

	"github.com/zpb84/go-game/lib/game"
	"github.com/zpb84/go-game/lib/interplay"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// newbie - реализация простого бота, который рандомно ставит ходы.
type newbie struct {
}

// SelectMove реализация интерфейса interplay.Agent
// FIX ME: Можно распаралелить вычисления
func (n *newbie) SelectMove(gameState *game.GameState) *game.Move {
	candidates := []game.Point{}
	for r := 1; r <= gameState.Board().Rows()+1; r++ {
		for c := 1; c <= gameState.Board().Columns()+1; c++ {
			candidate := game.Point{Row: r, Col: c}
			if gameState.IsValidMove(game.Play(candidate)) && !IsPointAnEye(gameState.Board(), candidate, gameState.NextPlayer()) {
				candidates = append(candidates, candidate)
			}
		}
	}
	if len(candidates) == 0 {
		return game.PassTurn()
	}
	return game.Play(candidates[rand.Intn(len(candidates))])
}

func NewNewbie() interplay.Agent {
	return &newbie{}
}
