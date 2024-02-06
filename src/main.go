package main

import (
	engine "github.com/andrewjmcgehee/shahgo/engine"
)

// just a driver for random testing for now
func main() {
	engine.InitAttacks()
	att := engine.RookRays(engine.A1, false, 0)
	for i := uint64(0); i < 4096; i++ {
		occ := engine.SetOccupancy(i, engine.CountBits(att), att)
		engine.Display(occ, true)
	}
}
