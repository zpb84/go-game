package game

import (
	"errors"
	"reflect"
)

var (
	ErrGroupIsNil = errors.New("group is nil")
	ErrGroupColor = errors.New("colors differ")
)

type Group struct {
	// Цвет группы
	Color Color
	// Камни
	Stones SetOfPoints
	// Степени свободы группы
	liberties SetOfPoints
}

func NewGroup(color Color, stones SetOfPoints, liberties SetOfPoints) *Group {
	return &Group{
		Color:     color,
		Stones:    stones,
		liberties: liberties,
	}
}

// RemoveLiberty удаление степеней свободы
func (g *Group) RemoveLiberty(points ...Point) {
	for _, p := range points {
		g.liberties.Remove(p)
	}
}

// AddLiberty добавление степеней свободы
func (g *Group) AddLiberty(points ...Point) {
	for _, p := range points {
		g.liberties.Add(p)
	}
}

// Merge объединение цепочек одноцветных камней
func (g *Group) Merge(other *Group) *Group {
	if g == nil || other == nil {
		return nil
	}
	if other.Color != g.Color {
		return nil
	}
	combinedStones := MergePoints(g.Stones, other.Stones)
	combinedLiberties := ExcludePoints(MergePoints(g.liberties, other.liberties), combinedStones)
	return &Group{
		Color:     g.Color,
		Stones:    combinedStones,
		liberties: combinedLiberties,
	}
}

func (g *Group) NumLiberties() int {
	return g.liberties.Len()
}

func (g *Group) Equal(other *Group) bool {
	if g == nil || other == nil {
		return true
	}
	return reflect.DeepEqual(g, other)
}
