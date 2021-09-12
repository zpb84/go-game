package go_board

type SetOfPoints map[Point]struct{}

// MergePoints объединяет множества точек в новое множество
func MergePoints(sets ...SetOfPoints) SetOfPoints {
	result := make(SetOfPoints)
	for _, s := range sets {
		for k := range s {
			result[k] = struct{}{}
		}
	}
	return result
}

// ExcludePoints формирует новое множество точек из разницы a-b
func ExcludePoints(a, b SetOfPoints) SetOfPoints {
	result := make(SetOfPoints)
	for k := range a {
		if _, ok := b[k]; !ok {
			result[k] = struct{}{}
		}
	}
	return result
}

func (s *SetOfPoints) Add(p Point) {
	(*s)[p] = struct{}{}
}

func (s *SetOfPoints) Exists(p Point) bool {
	_, ok := (*s)[p]
	return ok
}

func NewSetPoints(points ...Point) SetOfPoints {
	result := SetOfPoints{}
	for _, p := range points {
		result.Add(p)
	}
	return result
}
