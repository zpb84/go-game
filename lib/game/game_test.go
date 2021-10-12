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
	g := NewGame(boardSize)
	updater := func(p Point) error {
		m := Play(p)
		if new, err := g.ApplyMove(m); err != nil {
			return err
		} else {
			g = new
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
	return g, nil
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
      {"Row": 3,"Col": 2}
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
      {"Row": 6,"Col": 5}
    ]
  }`)
	_, err := LoadGameFromJSON(states, 9)
	if err != nil {
		t.Error(err)
	}
}
