package game

import (
	"math/rand"
	"time"
)

const (
	MAX63       = uint64(0x7fffffffffffffff)
	EMPTY_BOARD = uint64(0)
)

type ZobristHash struct {
	table map[pointState]uint64
}

type pointState struct {
	Point Point
	State Color
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Генерация уникальных значений для каждого хода
func NewZobristHash(size int) *ZobristHash {
	table := map[pointState]uint64{}
	exist := map[uint64]struct{}{}
	states := []Color{BLACK, WHITE}
	for r := 1; r <= size; r++ {
		for c := 1; c <= size; c++ {
			for _, s := range states {
				var rnd uint64
				for {
					rnd = rand.Uint64()
					_, ok := exist[rnd]
					if rnd <= MAX63 && !ok {
						exist[rnd] = struct{}{}
						break
					}
				}
				table[pointState{
					Point: NewPoint(r, c),
					State: s,
				}] = rnd
			}
		}
	}
	return &ZobristHash{
		table: table,
	}
}

func (z *ZobristHash) Get(p Point, state Color) uint64 {
	key := pointState{
		Point: p,
		State: state,
	}
	val, ok := z.table[key]
	if !ok {
		panic("ZobristHash: key not found")
	}
	return val
}
