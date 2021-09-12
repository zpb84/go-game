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
