package go_board

import "testing"

func TestSets(t *testing.T) {
	t.Run("Sets: count", func(t *testing.T) {
		p1 := NewPoint(1, 1)
		p2 := NewPoint(2, 2)
		s := SetOfPoints{}
		s.Add(p1)
		s.Add(p2)
		if len(s) != 2 {
			t.Error("error count")
		}
		for p := range s {
			if p != p1 && p != p2 {
				t.Error("error equal")
			}
		}
	})
	t.Run("Sets: Merge & Exclude", func(t *testing.T) {
		all := [][]Point{
			{
				NewPoint(1, 1),
				NewPoint(1, 2),
				NewPoint(2, 1),
			},
			{
				NewPoint(2, 2),
				NewPoint(1, 1),
				NewPoint(2, 3),
			},
		}
		s1 := NewSetPoints(all[0]...)
		s2 := NewSetPoints(all[1]...)
		s3 := MergePoints(s1, s2)
		if len(s3) != 5 {
			t.Error("Merge: count")
		}
		if s3.Exists(NewPoint(10, 10)) {
			t.Error("Merge: point not found")
		}
		for _, arr := range all {
			for _, p := range arr {
				if !s3.Exists(p) {
					t.Error("Merge: point from arrays not found")
				}
			}
		}
		s4 := ExcludePoints(s1, s2)
		if len(s4) != 2 {
			t.Error("Exclude: count")
		}
		if s4.Exists(NewPoint(1, 1)) {
			t.Error("Exclude: extra point ")
		}
		if !s4.Exists(NewPoint(1, 2)) {
			t.Error("Exclude: point not found")
		}
		if !s4.Exists(NewPoint(2, 1)) {
			t.Error("Exclude: point not found")
		}
	})
}
