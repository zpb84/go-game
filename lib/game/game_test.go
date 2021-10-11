package game

import (
	"encoding/json"
	"errors"
)

// jsonBoardState Описание состояния доски в JSON-формате.
// Камни по правилам будут расстанавливаться по доске
type jsonBoardState struct {
	White []Point `json:"white"`
	Black []Point `json:"black"`
}

// LoadState тестовый метод для отладки. Нужен для загрузки состояния доски
func LoadStateFromJSON(array []byte) (*GameState, error) {
	state := &jsonBoardState{}
	if err := json.Unmarshal(array, state); err != nil {
		return nil, err
	}
	if len(state.Black) != len(state.White) {
		return nil, errors.New("len White != len Black")
	}
	g := NewGame(9)
	updater := func(p Point) error {
		m := Play(p)
		if new, err := g.ApplyMove(m); err != nil {
			return err
		} else {
			g = new
		}
		return nil
	}
	for index := range state.Black {
		if err := updater(state.Black[index]); err != nil {
			return nil, err
		}
		if err := updater(state.White[index]); err != nil {
			return nil, err
		}
	}
	return g, nil
}
