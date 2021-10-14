package game

import (
	"errors"
)

var (
	ErrGroupIsNil = errors.New("group is nil")
	ErrGroupColor = errors.New("colors differ")
)

// Group описание цепочки(группы) камней
type Group struct {
	// Цвет группы
	Color Color
	// Камни
	stones *SetOfPoints
	// Степени свободы группы
	liberties *SetOfPoints
}

func NewGroup(color Color, stones *SetOfPoints, liberties *SetOfPoints) *Group {
	return &Group{
		Color:     color,
		stones:    stones,
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
	combinedStones := MergePoints(g.stones, other.stones)
	combinedLiberties := ExcludePoints(MergePoints(g.liberties, other.liberties), combinedStones)
	return &Group{
		Color:     g.Color,
		stones:    combinedStones,
		liberties: combinedLiberties,
	}
}

// NumLiberties возвращает количество степеней свободы группы
func (g *Group) NumLiberties() int {
	return g.liberties.Len()
}

// Equal глубокое сравнение групп
func (g *Group) Equal(other *Group) bool {
	if g == nil && other == nil {
		return true
	}
	if g.stones == nil || g.liberties == nil ||
		other == nil || other.stones == nil || other.liberties == nil {
		return false
	}
	if len(g.stones.points) != len(other.stones.points) ||
		len(g.liberties.points) != len(other.liberties.points) {
		return false
	}
	for k := range g.stones.points {
		if _, ok := other.stones.points[k]; !ok {
			return false
		}
	}
	for k := range g.liberties.points {
		if _, ok := other.liberties.points[k]; !ok {
			return false
		}
	}
	return true
}

// Copy создание глубокой копии группы
func (g *Group) Copy() *Group {
	if g == nil {
		return nil
	}
	return &Group{
		Color:     g.Color,
		liberties: g.liberties.Copy(),
		stones:    g.stones.Copy(),
	}
}
