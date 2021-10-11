package game

// Ход
type Move struct {
	point    Point
	isPass   bool
	isResign bool
}

// Play ход игрока
func Play(p Point) Move {
	return Move{
		point: p,
	}
}

// PassTurn игрок пропускает свой ход
func PassTurn() Move {
	return Move{
		point:  NewPoint(-1, -1),
		isPass: true,
	}
}

// Resign игрок завершает игру
func Resign() Move {
	return Move{
		point:    NewPoint(-1, -1),
		isResign: true,
	}
}

func (m Move) IsPass() bool {
	return m.isPass
}

func (m Move) IsResign() bool {
	return m.isResign
}

func (m Move) IsNil() bool {
	if m.isResign || m.isPass || m.point.Col <= 0 || m.point.Row <= 0 {
		return true
	}
	return false
}

func (m Move) Point() Point {
	return m.point
}
