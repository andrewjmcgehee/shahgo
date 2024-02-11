package main

import (
	engine "github.com/andrewjmcgehee/shahgo/engine"
)

// just a driver for random testing for now
func main() {
	engine.Display(uint64(engine.XORRand()), true)
	engine.Display(uint64(engine.XORRand()&0xffff), true)
	engine.Display(engine.Rand(), true)
	engine.Display(engine.MagicCandidate(), true)
}
