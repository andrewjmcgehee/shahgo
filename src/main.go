package main

import (
	engine "github.com/andrewjmcgehee/shahgo/engine"
)

// just a driver for random testing for now
func main() {
	engine.InitAttacks()
	// b := engine.ParseFEN(engine.FenEmpty)
	b := engine.ParseFEN(engine.FenStart)
	// b := engine.ParseFEN(engine.FenTricky)
	// b := engine.ParseFEN(engine.FenKill)
	// b := engine.ParseFEN(engine.FenCMK)
	engine.DisplayBitboard(b.OccupancyBitboards[engine.White], true)
	engine.DisplayBitboard(b.OccupancyBitboards[engine.Black], true)
	engine.DisplayBitboard(b.OccupancyBitboards[engine.Both], true)
	engine.Move(b, engine.Pawn, engine.D2, engine.D4)
	engine.Move(b, engine.Pawn, engine.F7, engine.F5)
	engine.Move(b, engine.Pawn, engine.C2, engine.C4)
	engine.Move(b, engine.Pawn, engine.F5, engine.F4)
	engine.DisplayBitboard(engine.MaskRookAttacksWithOccupancy(engine.C1, b.OccupancyBitboards[engine.Both]), true)

	if b.OccupancyBitboards[engine.Both] != b.OccupancyBitboards[engine.White]|b.OccupancyBitboards[engine.Black] {
		panic("blah")
	}
}
