package game

import "errors"

var (
	ErrBoardOutOfBounds    = errors.New("point out of bounds board")
	ErrBoardColorIsNotNone = errors.New("point has color")
	ErrBoardMargin         = errors.New("Merge error")
)

type Board struct {
	numRows int
	numCols int
	grid    map[Point]*Group
}

func NewBoard(numRows int, numCols int) *Board {
	return &Board{
		numRows: numRows,
		numCols: numCols,
		grid:    map[Point]*Group{},
	}
}

// PlaceStone проверка степеней свободы соседних точек
func (b *Board) PlaceStone(player Color, point Point) error {
	// Проверка выхода за доску
	if !b.IsOnGrid(point) {
		return ErrBoardOutOfBounds
	}
	// Проверка, лежит ли камень
	if b.GetColor(point) != NONE {
		return ErrBoardColorIsNotNone
	}
	liberties := SetOfPoints{}
	// Соседние группы того же цвета
	adjacentSameColor := map[*Group]struct{}{}
	// Соседние группы противника
	adjacentOppositeColor := map[*Group]struct{}{}

	// Просмотриваем всех соседей
	for _, neighbor := range point.Neighbors() {
		if !b.IsOnGrid(neighbor) {
			// Сосед вне границы доски
			continue
		}
		// Получаем группу камней по точке соседа

		if neighborGroup := b.GetGroup(neighbor); neighborGroup == nil {
			// Сосед не содержит камня
			liberties.Add(neighbor)
		} else {
			// Соседняя группа принадлежит игроку
			if neighborGroup.Color == player {
				if _, ok := adjacentSameColor[neighborGroup]; !ok {
					adjacentSameColor[neighborGroup] = struct{}{}
				}
			} else {
				// Соседняя группа принадлежит противнику
				if _, ok := adjacentOppositeColor[neighborGroup]; !ok {
					adjacentOppositeColor[neighborGroup] = struct{}{}
				}
			}
		}
	}
	newGroup := NewGroup(player, NewSetPoints(point), liberties)
	// Если в соседях есть цепочки того же цвета, то объединяем их
	for sameColor := range adjacentSameColor {
		newGroup = newGroup.Merge(sameColor)
		if newGroup == nil {
			return ErrBoardMargin
		}
	}
	// Обновляем информацию на доске по группам
	for groupPoint := range newGroup.Stones.points {
		b.grid[groupPoint] = newGroup
	}
	// Уменьшение степеней свободы у цепочек камней противоположного цвета
	for otherColor := range adjacentOppositeColor {
		otherColor.RemoveLiberty(point)
	}
	// Удаление групп противника, у которых степени свободы = 0
	for otherColor := range adjacentOppositeColor {
		if otherColor.NumLiberties() == 0 {
			b.RemoveGroup(otherColor)
		}
	}
	return nil
}

// IsOnGrid проверяет, не выходит ли точка за пределы доски
func (b *Board) IsOnGrid(point Point) bool {
	return 1 <= point.Row && point.Row <= b.numRows &&
		1 <= point.Col && point.Col <= b.numCols
}

// Возвращает цвет камня по координатам доски
func (b *Board) GetColor(point Point) Color {
	g, ok := b.grid[point]
	if !ok {
		return NONE
	}
	return g.Color
}

// GetGroup возвращает группу камней по координатам доски
func (b *Board) GetGroup(point Point) *Group {
	if g, ok := b.grid[point]; ok {
		return g
	}
	return nil
}

// RemoveGroup удаление цепочки камней
func (b *Board) RemoveGroup(g *Group) {
	for point := range g.Stones.points {
		for _, neighbor := range point.Neighbors() {
			if neighborGroup := b.GetGroup(neighbor); neighborGroup != nil {
				// Удаление цепочки приводит к увеличению свобод других групп
				if neighborGroup != g {
					neighborGroup.AddLiberty(point)
				}
			}
		}
		delete(b.grid, point)
	}
}