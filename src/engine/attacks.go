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
	if side == WHITE {
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

func BishopRays(square int, edges bool, occupied uint64) uint64 {
	rays := uint64(0)
	lo := 1
	hi := 6
	if edges {
		lo = 0
		hi = 7
	}
	rank := RankOf(square)
	file := FileOf(square)
	for r, f := rank+1, file+1; r <= hi && f <= hi; r, f = r+1, f+1 {
		bit := uint64(1) << uint(r*8+f)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	for r, f := rank+1, file-1; r <= hi && f >= lo; r, f = r+1, f-1 {
		bit := uint64(1) << uint(r*8+f)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	for r, f := rank-1, file+1; r >= lo && f <= hi; r, f = r-1, f+1 {
		bit := uint64(1) << uint(r*8+f)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	for r, f := rank-1, file-1; r >= lo && f >= lo; r, f = r-1, f-1 {
		bit := uint64(1) << uint(r*8+f)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	return rays
}

func RelevantBishopOccupants(square int) uint64 {
	return BishopRays(square, false, 0)
}

func MaskBishopAttacks(square int, occupied uint64) uint64 {
	return BishopRays(square, true, occupied)
}

func RookRays(square int, edges bool, occupied uint64) uint64 {
	rays := uint64(0)
	lo := 1
	hi := 6
	if edges {
		lo = 0
		hi = 7
	}
	rank := RankOf(square)
	file := FileOf(square)
	for r := rank - 1; r >= lo; r-- {
		bit := uint64(1) << uint(r*8+file)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	for r := rank + 1; r <= hi; r++ {
		bit := uint64(1) << uint(r*8+file)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	for f := file - 1; f >= lo; f-- {
		bit := uint64(1) << uint(rank*8+f)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	for f := file + 1; f <= hi; f++ {
		bit := uint64(1) << uint(rank*8+f)
		rays |= bit
		if occupied&bit != 0 {
			break
		}
	}
	return rays
}

func RelevantRookOccupants(square int) uint64 {
	return RookRays(square, false, 0)
}

func MaskRookAttacks(square int, occupied uint64) uint64 {
	return RookRays(square, true, occupied)
}
