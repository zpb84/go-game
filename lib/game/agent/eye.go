package agent

import "github.com/zpb84/go-game/lib/game"

// Глаз - степень свободы группы, которую не может занять противник

// Проверка, является ли точка "глазом" группы
func IsPointAnEye(board *game.Board, point game.Point, color game.Color) bool {
	if board.GetColor(point) != game.NONE {
		return false
	}
	// Для проверки "глаза" все соседние точки должны быть цвета игрока
	for _, neighbor := range point.Neighbors() {
		if board.IsOnGrid(neighbor) {
			neighborColor := board.GetColor(neighbor)
			if neighborColor != color {
				return false
			}
		}
	}
	// Минимальный контроль не менее 3-х углов(точки по диагонали), если точка находится в середине доски
	friendlyCorners := 0
	// Переменная для хранения точек вне доски
	offBoardCorners := 0

	// Коорлинаты всех углов
	corners := []game.Point{
		game.NewPoint(point.Row-1, point.Col-1),
		game.NewPoint(point.Row-1, point.Col+1),
		game.NewPoint(point.Row+1, point.Col-1),
		game.NewPoint(point.Row+1, point.Col+1),
	}
	for _, corner := range corners {
		if board.IsOnGrid(corner) {
			cornerColor := board.GetColor(corner)
			if cornerColor == color {
				friendlyCorners++
			}
		} else {
			offBoardCorners++
		}
	}

	if offBoardCorners > 0 {
		// Точка находится на краю или в углу доски, должны контролироваться все 4-е точки
		return (offBoardCorners + friendlyCorners) == 4
	}
	return friendlyCorners >= 3
}
