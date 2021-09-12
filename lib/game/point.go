package game

type Point struct {
	Row int
	Col int
}

// Neighbors получение соседних точек
func (p Point) Neighbors() []Point {
	return []Point{
		{p.Row - 1, p.Col},
		{p.Row + 1, p.Col},
		{p.Row, p.Col - 1},
		{p.Row, p.Col + 1},
	}
}

func NewPoint(row int, col int) Point {
	return Point{
		Row: row,
		Col: col,
	}
}
