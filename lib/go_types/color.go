package go_types

type Color int

const (
	BLACK Color = iota
	WHITE
)

func (c Color) String() string {
	return [...]string{"BLACK", "WHITE"}[c]
}
