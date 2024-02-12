package main

import (
	engine "github.com/andrewjmcgehee/shahgo/engine"
)

// just a driver for random testing for now
func main() {
	engine.InitAttacks()
	occupied := uint64(0)
	engine.SetBit(&occupied, engine.C5)
	engine.SetBit(&occupied, engine.F2)
	engine.SetBit(&occupied, engine.G7)
	engine.SetBit(&occupied, engine.B2)
	engine.SetBit(&occupied, engine.G5)
	engine.SetBit(&occupied, engine.E2)
	engine.SetBit(&occupied, engine.E7)
	engine.Display(occupied, true)
	engine.Display(engine.GetRookAttacks(engine.D4, occupied), true)
	engine.Display(engine.GetRookAttacks(engine.E5, occupied), true)
}
