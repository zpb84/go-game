package go_board

import (
	"errors"
	"reflect"

	"github.com/zpb84/go-game/lib/go_types"
)

var (
	ErrGroupIsNil = errors.New("group is nil")
	ErrGroupColor = errors.New("colors differ")
)

type Group struct {
	// Цвет группы
	color go_types.Color
	// Камни
	stones SetOfPoints
	// Степени свободы группы
	liberties SetOfPoints
}

func NewGroup(color go_types.Color, stones SetOfPoints, liberties SetOfPoints) *Group {
	return &Group{
		color:     color,
		stones:    stones,
		liberties: liberties,
	}
}

// RemoveLiberty удаление степеней свободы
func (g *Group) RemoveLiberty(points ...Point) {
	for _, p := range points {
		delete(g.liberties, p)
	}
}

// AddLiberty добавление степеней свободы
func (g *Group) AddLiberty(points ...Point) {
	for _, p := range points {
		g.liberties.Add(p)
	}
}

// Merge объединение цепочек одноцветных камней
func (g *Group) Merge(other *Group) (*Group, error) {
	if g == nil || other == nil {
		return nil, ErrGroupIsNil
	}
	if other.color != g.color {
		return nil, ErrGroupColor
	}
	combinedStones := MergePoints(g.stones, other.stones)
	combinedLiberties := ExcludePoints(MergePoints(g.liberties, other.liberties), combinedStones)
	return &Group{
		color:     g.color,
		stones:    combinedStones,
		liberties: combinedLiberties,
	}, nil
}

func (g *Group) NumLiberties() int {
	return len(g.liberties)
}

func (g *Group) Equal(other *Group) bool {
	if g == nil || other == nil {
		return true
	}
	return reflect.DeepEqual(g, other)
}
