package engine

const WHITE = 0
const BLACK = 1

type pair struct {
	row int
	col int
}

var BishopOccupancyCounts [64]uint64 = [64]uint64{
	6, 5, 5, 5, 5, 5, 5, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	6, 5, 5, 5, 5, 5, 5, 6,
}
var RookOccupancyCounts [64]uint64 = [64]uint64{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}

var PawnAttacks [2][64]uint64
var KnightAttacks [64]uint64
var KingAttacks [64]uint64

func InitAttacks() {
	for square := uint64(0); square < 64; square++ {
		PawnAttacks[WHITE][square] = MaskPawnAttacks(WHITE, square)
		PawnAttacks[BLACK][square] = MaskPawnAttacks(BLACK, square)
		KnightAttacks[square] = MaskKnightAttacks(square)
		KingAttacks[square] = MaskKingAttacks(square)
	}
}

func MaskPawnAttacks(side uint64, square uint64) uint64 {
	attacks := uint64(0)
	bitboard := uint64(0)
	SetBit(&bitboard, square)
	if side == WHITE {
		attacks |= bitboard >> 7
		attacks |= bitboard >> 9
	} else {
		attacks |= bitboard << 7
		attacks |= bitboard << 9
	}
	if bitboard&NOT_A == 0 { // pawn is on A file
		attacks &= NOT_H
	} else if bitboard&NOT_H == 0 { // pawn is on H file
		attacks &= NOT_A
	}
	return attacks
}

func MaskKnightAttacks(square uint64) uint64 {
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

func MaskKingAttacks(square uint64) uint64 {
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

func BishopRays(square uint64, edges bool, occupied uint64) uint64 {
	rays := uint64(0)
	r, f := RankOf(square)+1, FileOf(square)+1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r+1, f+1) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r += 1
		f += 1
	}
	r, f = RankOf(square)+1, FileOf(square)-1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r+1, f-1) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r += 1
		f -= 1
	}
	r, f = RankOf(square)-1, FileOf(square)+1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r-1, f+1) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r -= 1
		f += 1
	}
	r = RankOf(square) - 1
	f = FileOf(square) - 1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r-1, f-1) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r -= 1
		f -= 1
	}
	return rays
}

func RelevantBishopOccupants(square uint64) uint64 {
	return BishopRays(square, false, 0)
}

func MaskBishopAttacks(square uint64, occupied uint64) uint64 {
	return BishopRays(square, true, occupied)
}

func RookRays(square uint64, edges bool, occupied uint64) uint64 {
	rays := uint64(0)
	rank, file := RankOf(square)-1, FileOf(square)
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank-1, file) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		rank -= 1
	}
	rank = RankOf(square) + 1
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank+1, file) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		rank += 1
	}
	rank, file = RankOf(square), FileOf(square)-1
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank, file-1) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		file -= 1
	}
	file = FileOf(square) + 1
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank, file+1) {
			break
		}
		bit := uint64(1) << uint64(SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		file += 1
	}
	return rays
}

func RelevantRookOccupants(square uint64) uint64 {
	return RookRays(square, false, 0)
}

func MaskRookAttacks(square uint64, occupied uint64) uint64 {
	return RookRays(square, true, occupied)
}

func SetOccupancy(index uint64, bits uint64, attack_mask uint64) uint64 {
	occupancy := uint64(0)
	for count := uint64(0); count < bits; count++ {
		square := LSBIndex(attack_mask)
		PopBit(&attack_mask, square)
		if index&(1<<count) != 0 {
			occupancy |= 1 << square
		}
	}
	return occupancy
}
