package game

import (
	"encoding/json"
	"errors"
	"testing"
)

// jsonBoardState Описание состояния доски в JSON-формате.
// Камни по правилам будут расстанавливаться по доске
type jsonBoardState struct {
	Black []Point `json:"black"`
	White []Point `json:"white"`
}

// LoadGameFromJSON тестовый метод для отладки. Нужен для загрузки состояния доски
func LoadGameFromJSON(array []byte, boardSize int) (*GameState, error) {
	allMoves := &jsonBoardState{}
	if err := json.Unmarshal(array, allMoves); err != nil {
		return nil, err
	}
	if len(allMoves.Black) != len(allMoves.White) {
		return nil, errors.New("len White != len Black")
	}
	game := NewGame(boardSize)
	updater := func(p Point) error {
		m := Play(p)
		if newGame, err := game.ApplyMove(m); err != nil {
			return err
		} else {
			game = newGame
		}
		return nil
	}
	for index := range allMoves.Black {
		if err := updater(allMoves.Black[index]); err != nil {
			return nil, err
		}
		if err := updater(allMoves.White[index]); err != nil {
			return nil, err
		}
	}
	return game, nil
}

func TestGrap(t *testing.T) {
	states := []byte(`
  {
    "black": [
      {"Row": 2,"Col": 1},
      {"Row": 1,"Col": 2},
      {"Row": 1,"Col": 3},
      {"Row": 1,"Col": 4},
      {"Row": 2,"Col": 5},
      {"Row": 3,"Col": 5},
      {"Row": 4,"Col": 4},
      {"Row": 3,"Col": 3},
      {"Row": 3,"Col": 2},
      {"Row": 1,"Col": 1},
      {"Row": 3,"Col": 1},
      {"Row": 4,"Col": 3},
      {"Row": 1,"Col": 5}
    ],
    "white": [
      {"Row": 2,"Col": 2},
      {"Row": 2,"Col": 3},
      {"Row": 2,"Col": 4},
      {"Row": 3,"Col": 4},
      {"Row": 6,"Col": 1},
      {"Row": 6,"Col": 2},
      {"Row": 6,"Col": 3},
      {"Row": 6,"Col": 4},
      {"Row": 6,"Col": 5},
      {"Row": 7,"Col": 2},
      {"Row": 7,"Col": 3},
      {"Row": 7,"Col": 4},
      {"Row": 7,"Col": 5}
    ]
  }`)
	game, err := LoadGameFromJSON(states, 9)
	if err != nil {
		t.Error(err)
	}
	resultGroups := []*Group{
		{
			Color: BLACK,
			stones: NewSetPoints(NewPoint(3, 1),
				NewPoint(2, 1), NewPoint(1, 1), NewPoint(1, 2), NewPoint(1, 3),
				NewPoint(1, 4), NewPoint(1, 5), NewPoint(2, 5), NewPoint(3, 5),
				NewPoint(3, 2), NewPoint(3, 3), NewPoint(4, 3), NewPoint(4, 4),
			),
			liberties: NewSetPoints(
				NewPoint(1, 6), NewPoint(2, 6), NewPoint(3, 6), NewPoint(4, 5),
				NewPoint(5, 4), NewPoint(5, 3), NewPoint(4, 2), NewPoint(4, 1),
				NewPoint(2, 2), NewPoint(2, 3), NewPoint(2, 4), NewPoint(3, 4),
			),
		},
		{
			Color: WHITE,
			stones: NewSetPoints(
				NewPoint(6, 1), NewPoint(6, 2), NewPoint(6, 3),
				NewPoint(6, 4), NewPoint(6, 5), NewPoint(7, 2),
				NewPoint(7, 3), NewPoint(7, 4), NewPoint(7, 5),
			),
			liberties: NewSetPoints(
				NewPoint(5, 1), NewPoint(8, 4), NewPoint(8, 5), NewPoint(5, 5),
				NewPoint(8, 2), NewPoint(7, 1), NewPoint(7, 6), NewPoint(6, 6),
				NewPoint(5, 2), NewPoint(5, 3), NewPoint(5, 4), NewPoint(8, 3),
			),
		},
	}
	for _, group := range game.board.grid {
		find := false
		for _, model := range resultGroups {
			if model.Equal(group) {
				find = true
				break
			}
		}
		if !find {
			t.Error("group not found")
		}
	}
}
