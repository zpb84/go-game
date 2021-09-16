package game

// SetOfPoints множество точек
type SetOfPoints struct {
	points map[Point]struct{}
}

// MergePoints объединяет множества точек в новое множество
func MergePoints(sets ...SetOfPoints) SetOfPoints {
	result := SetOfPoints{
		points: map[Point]struct{}{},
	}
	for _, s := range sets {
		for k := range s.points {
			result.points[k] = struct{}{}
		}
	}
	return result
}

// ExcludePoints формирует новое множество точек из разницы a-b
func ExcludePoints(a, b SetOfPoints) SetOfPoints {
	result := SetOfPoints{
		points: map[Point]struct{}{},
	}
	for k := range a.points {
		if _, ok := b.points[k]; !ok {
			result.points[k] = struct{}{}
		}
	}
	return result
}

// Add добавление точки в множество
func (s *SetOfPoints) Add(p Point) {
	if s.points == nil {
		s.points = map[Point]struct{}{}
	}
	s.points[p] = struct{}{}
}

// Remove удаление точки из множества
func (s *SetOfPoints) Remove(p Point) {
	delete(s.points, p)
}

// Len количество точек в множестве
func (s *SetOfPoints) Len() int {
	return len(s.points)
}

// Exists входит ли точка в это множество
func (s *SetOfPoints) Exists(p Point) bool {
	_, ok := s.points[p]
	return ok
}

// ToArray преобразование множества в массив точек
func (s *SetOfPoints) ToArray() []Point {
	result := make([]Point, 0, len(s.points))
	for p := range s.points {
		result = append(result, p)
	}
	return result
}

func NewSetPoints(points ...Point) SetOfPoints {
	result := SetOfPoints{
		points: map[Point]struct{}{},
	}
	for _, p := range points {
		result.Add(p)
	}
	return result
}
