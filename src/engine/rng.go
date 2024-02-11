package engine

var rand_state uint32 = 1804289383

func XORRand() uint32 {
	number := rand_state
	number ^= number << 13
	number ^= number >> 17
	number ^= number << 5
	rand_state = number
	return rand_state
}

func Rand() uint64 {
	mask := uint32(0xffff)
	return uint64(XORRand()&mask) | uint64(XORRand()&mask)<<16 | uint64(XORRand()&mask)<<32 | uint64(XORRand()&mask)<<48
}

func MagicCandidate() uint64 {
	return Rand() & Rand() & Rand()
}
