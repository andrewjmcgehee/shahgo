package engine

type RNG struct {
	random_state uint32
}

func NewRNG() RNG {
	return RNG{random_state: 1804289383}
}

func (r *RNG) XORRand() uint32 {
	next := r.random_state
	next ^= next << 13
	next ^= next >> 17
	next ^= next << 5
	r.random_state = next
	return r.random_state
}

func (r *RNG) Rand() uint64 {
	mask := uint32(0xffff)
	return (uint64(r.XORRand()&mask) |
		uint64(r.XORRand()&mask)<<16 |
		uint64(r.XORRand()&mask)<<32 |
		uint64(r.XORRand()&mask)<<48)
}

func (r *RNG) MagicCandidate() uint64 {
	return r.Rand() & r.Rand() & r.Rand()
}
