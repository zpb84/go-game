package interplay

import "github.com/zpb84/go-game/lib/game"

// Agent интерфейс для бота
type Agent interface {
	SelectMove(gameState *game.GameState) game.Move
}
