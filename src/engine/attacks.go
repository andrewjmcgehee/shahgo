package engine

const WHITE = 0
const BLACK = 1

type pair struct {
	row int
	col int
}

var PawnAttacks [2][64]uint64
var KnightAttacks [64]uint64

func InitAttacks() {
	for square := 0; square < 64; square++ {
		PawnAttacks[WHITE][square] = MaskPawnAttacks(WHITE, square)
		PawnAttacks[BLACK][square] = MaskPawnAttacks(BLACK, square)
		KnightAttacks[square] = MaskKnightAttacks(square)
	}
}

func MaskPawnAttacks(side int, square int) uint64 {
	attacks := uint64(0)
	bitboard := uint64(0)
	SetBit(&bitboard, square)
	if side == 0 {
		attacks |= bitboard << 7
		attacks |= bitboard << 9
	} else {
		attacks |= bitboard >> 7
		attacks |= bitboard >> 9
	}
	if bitboard&NOT_A == 0 { // pawn is on A file
		attacks &= NOT_H
	} else if bitboard&NOT_H == 0 { // pawn is on H file
		attacks &= NOT_A
	}
	return attacks
}

func MaskKnightAttacks(square int) uint64 {
	attacks := uint64(0)
	bitboard := uint64(0)
	SetBit(&bitboard, square)
	attacks |= bitboard << 17
	attacks |= bitboard >> 17
	attacks |= bitboard << 15
	attacks |= bitboard >> 15
	attacks |= bitboard << 10
	attacks |= bitboard >> 10
	attacks |= bitboard << 6
	attacks |= bitboard >> 6
	if bitboard&NOT_AB == 0 { // knight is on A or B file
		attacks &= NOT_GH
	} else if bitboard&NOT_GH == 0 { // knight is on G or H file
		attacks &= NOT_AB
	}
	return attacks
}
