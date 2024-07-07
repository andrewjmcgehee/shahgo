package engine

import "fmt"

var magicRNG *RNG

func magicCandidate() uint64 {
	// has the effect of reducing the number of 1 bits
	return magicRNG.Rand() & magicRNG.Rand() & magicRNG.Rand()
}

func findMagic(square, relevantBits int, rook bool) uint64 {
	var occupancies [4096]uint64
	var attacks [4096]uint64
	var usedAttacks [4096]uint64
	var attackMask uint64
	if rook {
		attackMask = MaskRookAttacks(square)
	} else {
		attackMask = MaskBishopAttacks(square)
	}
	for i := 0; i < 1<<relevantBits; i++ {
		occupancies[i] = SetOccupancy(i, relevantBits, attackMask)
		if rook {
			attacks[i] = MaskRookAttacksWithOccupancy(square, occupancies[i])
		} else {
			attacks[i] = MaskBishopAttacksWithOccupancy(square, occupancies[i])
		}
	}
	for random := uint64(0); random < 100000000; random++ {
		magic := magicCandidate()
		if CountBits((attackMask*magic)&0xFF00000000000000) < 6 {
			continue
		}
		for i := 0; i < 4096; i++ {
			usedAttacks[i] = 0
		}
		failed := false
		for i := uint64(0); !failed && i < 1<<relevantBits; i++ {
			magicIdx := int(occupancies[i] * magic >> (64 - relevantBits))
			if usedAttacks[magicIdx] == uint64(0) {
				usedAttacks[magicIdx] = attacks[i]
			} else {
				failed = true
			}
		}
		if !failed {
			return magic
		}
	}
	return 0
}

func InitMagics() {
	if magicRNG == nil {
		magicRNG = NewRNG()
	}
	// rook magics
	for square := 0; square < 64; square++ {
		fmt.Printf("  0x%016x\n", findMagic(square, rookOccupancyCounts()[square], true))
	}
	fmt.Println()
	// bishop magics
	for square := 0; square < 64; square++ {
		fmt.Printf("  0x%016x\n", findMagic(square, bishopOccupancyCounts()[square], false))
	}
}
