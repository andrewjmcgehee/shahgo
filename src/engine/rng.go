package engine

type RNG struct {
	random_state uint32
}

func NewRNG() *RNG {
	return &RNG{random_state: 1804289383}
}

func (r *RNG) XORRand() uint32 {
	// XOR shift algorithm https://en.wikipedia.org/wiki/Xorshift#Example_implementation
	next := r.random_state
	next ^= next << 13
	next ^= next >> 17
	next ^= next << 5
	r.random_state = next
	return r.random_state
}

func (r *RNG) Rand() uint64 {
	mask := uint64(0xffff)
	a := uint64(r.XORRand()) & mask
	b := uint64(r.XORRand()) & mask
	c := uint64(r.XORRand()) & mask
	d := uint64(r.XORRand()) & mask
	return a | b<<16 | c<<32 | d<<48
}
