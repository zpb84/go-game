package game

type Color int

const (
	BLACK Color = iota
	WHITE
	NONE
)

func (c Color) String() string {
	return [...]string{"BLACK", "WHITE", "NONE"}[c]
}

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
