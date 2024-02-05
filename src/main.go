package main

import (
	engine "github.com/andrewjmcgehee/shahgo/engine"
)

// just a driver for random testing for now
func main() {
	squares := [7]int{engine.A1, engine.A4, engine.B4, engine.E4, engine.G4, engine.H4, engine.H8}
	for _, square := range squares {
		engine.Display(engine.MaskRookAttacks(square), true)
	}
}
