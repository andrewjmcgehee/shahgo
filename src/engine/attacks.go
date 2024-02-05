package engine

const WHITE = 0
const BLACK = 1

type pair struct {
	row int
	col int
}

var PawnAttacks [2][64]uint64
var KnightAttacks [64]uint64
var KingAttacks [64]uint64

func InitAttacks() {
	for square := 0; square < 64; square++ {
		PawnAttacks[WHITE][square] = MaskPawnAttacks(WHITE, square)
		PawnAttacks[BLACK][square] = MaskPawnAttacks(BLACK, square)
		KnightAttacks[square] = MaskKnightAttacks(square)
		KingAttacks[square] = MaskKingAttacks(square)
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
	attacks |= bitboard << 15
	attacks |= bitboard << 10
	attacks |= bitboard << 6
	attacks |= bitboard >> 6
	attacks |= bitboard >> 10
	attacks |= bitboard >> 15
	attacks |= bitboard >> 17
	if bitboard&NOT_AB == 0 { // knight is on A or B file
		attacks &= NOT_GH
	} else if bitboard&NOT_GH == 0 { // knight is on G or H file
		attacks &= NOT_AB
	}
	return attacks
}

func MaskKingAttacks(square int) uint64 {
	attacks := uint64(0)
	bitboard := uint64(0)
	SetBit(&bitboard, square)
	attacks |= bitboard << 9
	attacks |= bitboard << 8
	attacks |= bitboard << 7
	attacks |= bitboard << 1
	attacks |= bitboard >> 1
	attacks |= bitboard >> 7
	attacks |= bitboard >> 8
	attacks |= bitboard >> 9
	if bitboard&NOT_A == 0 { // king is on A file
		attacks &= NOT_H
	} else if bitboard&NOT_H == 0 { // king is on H file
		attacks &= NOT_A
	}
	return attacks
}

func MaskBishopAttacks(square int) uint64 {
	attacks := uint64(0)
	rank := RankOf(square)
	file := FileOf(square)
	for r, f := rank+1, file+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		attacks |= uint64(1) << uint(r*8+f)
	}
	for r, f := rank+1, file-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		attacks |= uint64(1) << uint(r*8+f)
	}
	for r, f := rank-1, file+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		attacks |= uint64(1) << uint(r*8+f)
	}
	for r, f := rank-1, file-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		attacks |= uint64(1) << uint(r*8+f)
	}
	return attacks
}

func MaskRookAttacks(square int) uint64 {
	attacks := uint64(0)
	rank := RankOf(square)
	file := FileOf(square)
	for r := 1; r <= 6; r++ {
		if r != rank {
			attacks |= uint64(1) << uint(r*8+file)
		}
	}
	for f := 1; f <= 6; f++ {
		if f != file {
			attacks |= uint64(1) << uint(rank*8+f)
		}
	}
	return attacks
}
