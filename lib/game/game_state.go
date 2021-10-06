package game

import "reflect"

// GameState описывает текущее состояние игры.
// A через поле previousState можно получить состояния всех предыдущих ходов
type GameState struct {
	board         *Board
	nextPlayer    Color
	previousState *GameState
	lastMove      *Move
}

func (gs *GameState) Board() *Board {
	return gs.board
}

func NewGameState(board *Board, nextPlayer Color, previousState *GameState, lastMove *Move) *GameState {
	return &GameState{
		board:         board,
		nextPlayer:    nextPlayer,
		previousState: previousState,
		lastMove:      lastMove,
	}
}

// ApplyMove применение нового хода, в результате должно получиться новое состояние игры
func (gs *GameState) ApplyMove(move *Move) (*GameState, error) {
	var nextBoard *Board
	if move.point != nil {
		nextBoard = gs.board.Copy()
		if err := nextBoard.PlaceStone(gs.nextPlayer, *move.point); err != nil {
			return nil, err
		}
	} else {
		nextBoard = gs.board
	}
	return &GameState{
		board:         nextBoard,
		nextPlayer:    gs.nextPlayer.Other(),
		previousState: gs,
		lastMove:      move,
	}, nil
}

// IsOver определяет конец игры
func (gs *GameState) IsOver() bool {
	if gs.lastMove == nil {
		return false
	}
	if gs.lastMove.isResign {
		return true
	}
	if gs.previousState == nil {
		return false
	}
	secondLastMove := gs.previousState.lastMove
	if secondLastMove == nil {
		return false
	}
	return secondLastMove.isPass && gs.lastMove.isPass
}

// IsMoveSelfCapture проверяет на самозахват группы (когда степени свободы обнуляются) или на наличие ошибки
// Если true, то такой ход делать нельзя
func (gs *GameState) isMoveSelfCapture(player Color, move *Move) bool {
	if move.point == nil {
		return false
	}
	nextBoard := gs.board.Copy()
	if err := nextBoard.PlaceStone(player, *move.point); err != nil {
		return true
	}
	group := nextBoard.GetGroup(*move.point)
	return group.NumLiberties() == 0
}

// DoesMoveViolateKO проверяет нарушает ли ход правило КО
// Пробегает по всем состояниям и смотрит, не была ли доска в таком состоянии ранее
func (gs *GameState) doesMoveViolateKO(player Color, move *Move) bool {
	if move.point == nil {
		return false
	}

	nextBoard := gs.board.Copy()
	if err := nextBoard.PlaceStone(player, *move.point); err != nil {
		return true
	}
	nextPlayer := player.Other()

	pastState := gs.previousState
	for pastState != nil {
		if reflect.DeepEqual(pastState.board.grid, nextBoard) && pastState.nextPlayer == nextPlayer {
			return true
		}
		pastState = pastState.previousState
	}
	return false
}

// IsValidMove провреяет можно ли сделать ход
func (gs *GameState) IsValidMove(move *Move) bool {
	if gs.IsOver() {
		return false
	}
	if move.isPass || move.isResign {
		return true
	}
	return gs.board.GetGroup(*move.point) == nil &&
		!gs.isMoveSelfCapture(gs.nextPlayer, move) &&
		!gs.doesMoveViolateKO(gs.nextPlayer, move)
}

func (gs *GameState) NextPlayer() Color {
	return gs.nextPlayer
}

func NewGame(boardSize int) *GameState {
	return &GameState{
		board:      NewBoard(boardSize, boardSize),
		nextPlayer: BLACK,
	}
}
