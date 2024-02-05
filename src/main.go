package main

import (
	engine "github.com/andrewjmcgehee/shahgo/engine"
)

// just a driver for random testing for now
func main() {
	// bitboard := uint64(0)
	// square := engine.H4
	// engine.SetBit(&bitboard, square)
	// engine.Display(bitboard, true)
	// engine.Display(engine.MaskPawnAttacks(square, engine.BLACK), true)
	squares := [3]int{engine.A4, engine.H4, engine.E4}
	engine.InitAttacks()
	for _, square := range squares {
		engine.Display(engine.PawnAttacks[engine.WHITE][square], true)
		engine.Display(engine.PawnAttacks[engine.BLACK][square], true)
	}
}
