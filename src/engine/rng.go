package engine

type RNG struct {
	randomState uint32
}

func NewRNG() *RNG {
	return &RNG{randomState: 1804289383}
}

func (r *RNG) XORRand() uint32 {
	// XOR shift algorithm https://en.wikipedia.org/wiki/Xorshift#Example_implementation
	next := r.randomState
	next ^= next << 13
	next ^= next >> 17
	next ^= next << 5
	r.randomState = next
	return r.randomState
}

func (r *RNG) Rand() uint64 {
	mask := uint64(0xffff)
	a := uint64(r.XORRand()) & mask
	b := uint64(r.XORRand()) & mask
	c := uint64(r.XORRand()) & mask
	d := uint64(r.XORRand()) & mask
	return a | b<<16 | c<<32 | d<<48
}
