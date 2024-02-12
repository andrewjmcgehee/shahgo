package engine

import "fmt"

var magic_rng *RNG

func magic_candidate() uint64 {
	// has the effect of reducing the number of 1 bits
	return magic_rng.Rand() & magic_rng.Rand() & magic_rng.Rand()
}

func find_magic(square uint64, relevant_bits uint64, rook bool) uint64 {
	var occupancies [4096]uint64
	var attacks [4096]uint64
	var used_attacks [4096]uint64
	var attack_mask uint64
	if rook {
		attack_mask = MaskRookAttacks(square)
	} else {
		attack_mask = MaskBishopAttacks(square)
	}
	for i := uint64(0); i < 1<<relevant_bits; i++ {
		occupancies[i] = SetOccupancy(i, relevant_bits, attack_mask)
		if rook {
			attacks[i] = MaskRookAttacksWithOccupancy(square, occupancies[i])
		} else {
			attacks[i] = MaskBishopAttacksWithOccupancy(square, occupancies[i])
		}
	}
	for random := uint64(0); random < 100000000; random++ {
		magic := magic_candidate()
		if CountBits((attack_mask*magic)&0xFF00000000000000) < 6 {
			continue
		}
		for i := 0; i < 4096; i++ {
			used_attacks[i] = 0
		}
		failed := false
		for i := uint64(0); !failed && i < 1<<relevant_bits; i++ {
			magic_idx := int(occupancies[i] * magic >> (64 - relevant_bits))
			if used_attacks[magic_idx] == uint64(0) {
				used_attacks[magic_idx] = attacks[i]
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
	if magic_rng == nil {
		magic_rng = NewRNG()
	}
	// rook magics
	for square := uint64(0); square < 64; square++ {
		fmt.Printf("  0x%016x\n", find_magic(square, getRookOccupancyCounts()[square], true))
	}
	fmt.Println()
	// bishop magics
	for square := uint64(0); square < 64; square++ {
		fmt.Printf("  0x%016x\n", find_magic(square, getBishopOccupancyCounts()[square], false))
	}
}
