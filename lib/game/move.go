package game

// Ход
type Move struct {
	point    *Point
	isPass   bool
	isResign bool
}

// Play ход игрока
func Play(p Point) Move {
	return Move{
		point: &p,
	}
}

// PassTurn игрок пропускает свой ход
func PassTurn() Move {
	return Move{
		isPass: true,
	}
}

// Resign игрок завершает игру
func Resign() Move {
	return Move{
		isResign: true,
	}
}
