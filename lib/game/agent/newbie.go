package agent

import (
	"github.com/zpb84/go-game/lib/game"
	"github.com/zpb84/go-game/lib/interplay"
)

// newbie - реализация простого бота, который рандомно ставит ходы.
type newbie struct {
}

// SelectMove реализация интерфейса interplay.Agent
func (n *newbie) SelectMove(gameState *game.GameState) *game.GameState {
	return nil
}

func NewNewbie() interplay.Agent {
	return &newbie{}
}
