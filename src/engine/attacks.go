package engine

const WHITE = 0
const BLACK = 1

type pair struct {
	row int
	col int
}

var PawnAttacks [2][64]uint64

func InitAttacks() {
	// knight_delta := [8]pair{{-2, -1}, {-2, -1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	// king_delta := [8]pair{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for square := 0; square < 64; square++ {
		PawnAttacks[WHITE][square] = MaskPawnAttacks(WHITE, square)
		PawnAttacks[BLACK][square] = MaskPawnAttacks(BLACK, square)
	}
}

func MaskPawnAttacks(side int, square int) uint64 {
	attacks := uint64(0)
	bitboard := uint64(0)
	SetBit(&bitboard, square)
	if side == 0 {
		if bitboard<<7&NOT_H != 0 {
			attacks |= bitboard << 7
		}
		if bitboard<<9&NOT_A != 0 {
			attacks |= bitboard << 9
		}
	} else {
		if bitboard>>7&NOT_A != 0 {
			attacks |= bitboard >> 7
		}
		if bitboard>>9&NOT_H != 0 {
			attacks |= bitboard >> 9
		}
	}
	return attacks
}
