package main

import (
	engine "github.com/andrewjmcgehee/shahgo/engine"
)

// just a driver for random testing for now
func main() {
	blocker := uint64(0)
	engine.SetBit(&blocker, engine.C6)
	engine.SetBit(&blocker, engine.D3)
	squares := [1]int{engine.D6}
	for _, square := range squares {
		engine.Display(engine.RookRays(square, true, blocker), true)
	}
}
