package game

import "fmt"

var (
	UseFastViolateKO = true
)

// GameState описывает текущее состояние игры.
// A через поле previousState можно получить состояния всех предыдущих ходов
type GameState struct {
	board      *Board
	nextPlayer Color
	lastMove   Move

	previousState *GameState
	previousHash  hashState
}

type hashState struct {
	state map[keyHashState]struct{}
}

type keyHashState struct {
	Player Color
	Hash   uint64
}

func (hs hashState) Copy() hashState {
	newHash := map[keyHashState]struct{}{}
	for k := range hs.state {
		newHash[k] = struct{}{}
	}
	return hashState{
		state: newHash,
	}
}

func (hs hashState) AddState(player Color, hash uint64) hashState {
	result := hs.Copy()
	result.state[keyHashState{
		Player: player,
		Hash:   hash,
	}] = struct{}{}
	return result
}

func (hs hashState) Exist(player Color, hash uint64) bool {
	_, ok := hs.state[keyHashState{
		Player: player,
		Hash:   hash,
	}]
	return ok
}

func (gs *GameState) Board() *Board {
	return gs.board
}

func NewGameState(board *Board, nextPlayer Color, previousState *GameState, lastMove Move) *GameState {
	return &GameState{
		board:         board,
		nextPlayer:    nextPlayer,
		previousState: previousState,
		lastMove:      lastMove,
	}
}

// ApplyMove применение нового хода, в результате должно получиться новое состояние игры
func (gs *GameState) ApplyMove(move Move) (*GameState, error) {
	var nextBoard *Board
	if !move.IsNil() {
		nextBoard = gs.board.Copy()
		if err := nextBoard.PlaceStone(gs.nextPlayer, move.Point()); err != nil {
			return nil, err
		}
	} else {
		nextBoard = gs.board
	}
	var hs hashState
	if gs != nil {
		hs = gs.previousHash.AddState(gs.nextPlayer, gs.board.ZobristHash())
	} else {
		hs.state = map[keyHashState]struct{}{}
	}
	return &GameState{
		board:      nextBoard,
		nextPlayer: gs.nextPlayer.Other(),
		lastMove:   move,

		previousState: gs,
		previousHash:  hs,
	}, nil
}

// IsOver определяет конец игры
func (gs *GameState) IsOver() bool {
	if gs.lastMove.isResign {
		return true
	}
	if gs.previousState == nil {
		return false
	}
	secondLastMove := gs.previousState.lastMove
	return secondLastMove.isPass && gs.lastMove.isPass
}

// isMoveSelfCapture проверяет на самозахват группы (когда степени свободы обнуляются) или на наличие ошибки
// Если true, то такой ход делать нельзя
func (gs *GameState) isMoveSelfCapture(player Color, move Move) bool {
	if move.IsNil() {
		return false
	}
	nextBoard := gs.board.Copy()
	if err := nextBoard.PlaceStone(player, move.Point()); err != nil {
		return true
	}
	group := nextBoard.GetGroup(move.Point())
	return group.NumLiberties() == 0
}

// DoesMoveViolateKO проверяет нарушает ли ход правило КО
// Пробегает по всем состояниям и смотрит, не была ли доска в таком состоянии ранее
func (gs *GameState) doesMoveViolateKO(player Color, move Move) bool {
	if UseFastViolateKO {
		return gs.fastViolateKO(player, move)
	}
	return gs.slowViolateKO(player, move)
}

func (gs *GameState) slowViolateKO(player Color, move Move) bool {
	if move.IsNil() {
		return false
	}

	nextBoard := gs.board.Copy()
	if err := nextBoard.PlaceStone(player, move.Point()); err != nil {
		return true
	}
	nextPlayer := player.Other()

	pastState := gs.previousState
	for pastState != nil {
		if pastState.board.Equal(nextBoard) && pastState.nextPlayer == nextPlayer {
			return true
		}
		pastState = pastState.previousState
	}
	return false
}

func (gs *GameState) fastViolateKO(player Color, move Move) bool {
	if move.IsNil() {
		return false
	}

	nextBoard := gs.board.Copy()
	if err := nextBoard.PlaceStone(player, move.Point()); err != nil {
		panic(fmt.Sprintf("PlaceStone return error: %s", err))
	}

	return gs.previousHash.Exist(player.Other(), nextBoard.ZobristHash())
}

// IsValidMove провреяет можно ли сделать ход
func (gs *GameState) IsValidMove(move Move) bool {
	if gs.IsOver() {
		return false
	}
	if move.isPass || move.isResign {
		return true
	}
	return gs.board.GetGroup(move.Point()) == nil &&
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
