package game

import (
	"testing"
)

func TestGroup(t *testing.T) {
	t.Run("Groups", func(t *testing.T) {
		var g1, g2 *Group
		if !g1.Equal(g2) {
			t.Error("Equal nil")
		}
		s1 := NewSetPoints(
			NewPoint(1, 1),
			NewPoint(1, 2),
			NewPoint(2, 1),
		)
		g1 = NewGroup(WHITE, s1, s1)
		if g1.NumLiberties() != 3 {
			t.Error("NumLiberties: liberties count")
		}
		g1.RemoveLiberty(NewPoint(10, 10), NewPoint(1, 1))
		if g1.NumLiberties() != 2 {
			t.Error("RemoveLiberty: liberties count")
		}
		g1.AddLiberty(NewPoint(10, 10), NewPoint(1, 2))
		if g1.NumLiberties() != 3 {
			t.Error("AddLiberty: liberties count")
		}
		s2 := NewSetPoints(
			NewPoint(1, 1),
			NewPoint(1, 2),
			NewPoint(2, 1),
		)
		g2 = NewGroup(WHITE, s2, s2)
		if g1.Equal(g2) {
			t.Error("Equal")
		}
		s1 = NewSetPoints(
			NewPoint(1, 1),
			NewPoint(1, 2),
			NewPoint(2, 1),
		)
		g1 = NewGroup(WHITE, s1, s1)
		if !g1.Equal(g2) {
			t.Error("Not equal")
		}
	})
	t.Run("Groups.Merge", func(t *testing.T) {
		var g1, g2 *Group
		g3 := g1.Merge(g2)
		if g3 != nil {
			t.Error("Merge nil")
		}
		g1 = NewGroup(WHITE,
			NewSetPoints(
				NewPoint(1, 1),
				NewPoint(2, 1),
				NewPoint(3, 1),
			),
			NewSetPoints(
				NewPoint(1, 0),
				NewPoint(2, 0),
				NewPoint(3, 0),
			),
		)
		g2 = NewGroup(BLACK,
			NewSetPoints(
				NewPoint(1, 1),
				NewPoint(1, 2),
				NewPoint(1, 3),
			),
			NewSetPoints(
				NewPoint(0, 1),
				NewPoint(0, 2),
				NewPoint(0, 3),
			),
		)
		g3 = g1.Merge(g2)
		if g3 != nil {
			t.Error("Merge color")
		}
		g2 = NewGroup(WHITE,
			NewSetPoints(
				NewPoint(1, 1),
				NewPoint(1, 2),
				NewPoint(0, 2),
			),
			NewSetPoints(
				NewPoint(10, 2),
				NewPoint(20, 2),
				NewPoint(20, 1),
			),
		)
		g3 = g1.Merge(g2)
		if g3 == nil {
			t.Error("Merge")
		}
	})
}
