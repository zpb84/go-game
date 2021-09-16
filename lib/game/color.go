package game

// Цвет игроков
type Color int

const (
	BLACK Color = iota
	WHITE
	NONE
)

func (c Color) String() string {
	return [...]string{"BLACK", "WHITE", "NONE"}[c]
}

// Other - метод возвращает цвет следующего игрока
func (c Color) Other() Color {
	switch c {
	case BLACK:
		return WHITE
	case WHITE:
		return BLACK
	default:
		return NONE
	}
}
