package engine

var PawnAttacks [2][64]uint64
var KnightAttacks [64]uint64
var KingAttacks [64]uint64
var BishopMasks [64]uint64
var BishopAttacks [64][512]uint64
var RookMasks [64]uint64
var RookAttacks [64][4096]uint64

func InitAttacks() {
	InitLeaperAttacks()
	InitSliderAttacks(false) // bishops
	InitSliderAttacks(true)  // rooks
}

func InitLeaperAttacks() {
	for square := 0; square < 64; square++ {
		PawnAttacks[White][square] = MaskPawnAttacks(White, square)
		PawnAttacks[Black][square] = MaskPawnAttacks(Black, square)
		KnightAttacks[square] = MaskKnightAttacks(square)
		KingAttacks[square] = MaskKingAttacks(square)
	}
}

func InitSliderAttacks(rook bool) {
	var attackMask uint64
	var relevantBits int
	for square := 0; square < 64; square++ {
		if rook {
			RookMasks[square] = MaskRookAttacks(square)
			attackMask = RookMasks[square]
			relevantBits = CountBits(attackMask)
		} else {
			BishopMasks[square] = MaskBishopAttacks(square)
			attackMask = BishopMasks[square]
			relevantBits = CountBits(attackMask)
		}
		for i := 0; i < 1<<relevantBits; i++ {
			occupied := SetOccupancy(i, relevantBits, attackMask)
			if rook {
				magicIdx := (occupied * rookMagics()[square]) >> (64 - rookOccupancyCounts()[square])
				RookAttacks[square][magicIdx] = MaskRookAttacksWithOccupancy(square, occupied)
			} else {
				magicIdx := (occupied * bishopMagics()[square]) >> (64 - bishopOccupancyCounts()[square])
				BishopAttacks[square][magicIdx] = MaskBishopAttacksWithOccupancy(square, occupied)
			}
		}
	}
}

func MaskPawnAttacks(side, square int) uint64 {
	attacks := uint64(0)
	bitboard := uint64(0)
	SetBit(&bitboard, square)
	if side == White {
		attacks |= bitboard >> 7
		attacks |= bitboard >> 9
	} else {
		attacks |= bitboard << 7
		attacks |= bitboard << 9
	}
	if bitboard&NotFileA == 0 { // pawn is on A file
		attacks &= NotFileH
	} else if bitboard&NotFileH == 0 { // pawn is on H file
		attacks &= NotFileA
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
	if bitboard&NotFilesAB == 0 { // knight is on A or B file
		attacks &= NotFilesGH
	} else if bitboard&NotFilesGH == 0 { // knight is on G or H file
		attacks &= NotFilesAB
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
	if bitboard&NotFileA == 0 { // king is on A file
		attacks &= NotFileH
	} else if bitboard&NotFileH == 0 { // king is on H file
		attacks &= NotFileA
	}
	return attacks
}

func BishopRays(square int, edges bool, occupied uint64) uint64 {
	rays := uint64(0)
	r, f := RankOf(square)+1, FileOf(square)+1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r+1, f+1) {
			break
		}
		bit := uint64(1 << SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r++
		f++
	}
	r, f = RankOf(square)+1, FileOf(square)-1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r+1, f-1) {
			break
		}
		bit := uint64(1 << SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r++
		f--
	}
	r, f = RankOf(square)-1, FileOf(square)+1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r-1, f+1) {
			break
		}
		bit := uint64(1 << SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r--
		f++
	}
	r = RankOf(square) - 1
	f = FileOf(square) - 1
	for SafeCoord(r, f) {
		if !edges && !SafeCoord(r-1, f-1) {
			break
		}
		bit := uint64(1 << SquareFrom(r, f))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		r--
		f--
	}
	return rays
}

func MaskBishopAttacks(square int) uint64 {
	return BishopRays(square, false, 0)
}

func MaskBishopAttacksWithOccupancy(square int, occupied uint64) uint64 {
	return BishopRays(square, true, occupied)
}

func GetBishopAttacks(square int, occupied uint64) uint64 {
	occupied &= BishopMasks[square]
	occupied *= bishopMagics()[square]
	occupied >>= 64 - bishopOccupancyCounts()[square]
	return BishopAttacks[square][occupied]
}

func RookRays(square int, edges bool, occupied uint64) uint64 {
	rays := uint64(0)
	rank, file := RankOf(square)-1, FileOf(square)
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank-1, file) {
			break
		}
		bit := uint64(1 << SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		rank--
	}
	rank = RankOf(square) + 1
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank+1, file) {
			break
		}
		bit := uint64(1 << SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		rank++
	}
	rank, file = RankOf(square), FileOf(square)-1
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank, file-1) {
			break
		}
		bit := uint64(1 << SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		file--
	}
	file = FileOf(square) + 1
	for SafeCoord(rank, file) {
		if !edges && !SafeCoord(rank, file+1) {
			break
		}
		bit := uint64(1 << SquareFrom(rank, file))
		rays |= bit
		if occupied&bit != 0 {
			break
		}
		file++
	}
	return rays
}

func MaskRookAttacks(square int) uint64 {
	return RookRays(square, false, 0)
}

func MaskRookAttacksWithOccupancy(square int, occupied uint64) uint64 {
	return RookRays(square, true, occupied)
}

func GetRookAttacks(square int, occupied uint64) uint64 {
	occupied &= RookMasks[square]
	occupied *= rookMagics()[square]
	occupied >>= 64 - rookOccupancyCounts()[square]
	return RookAttacks[square][occupied]
}

func SetOccupancy(index, relevantBits int, attackMask uint64) uint64 {
	occupied := uint64(0)
	for count := 0; count < relevantBits; count++ {
		square := LSBIndex(attackMask)
		PopBit(&attackMask, square)
		if index&(1<<count) != 0 {
			occupied |= 1 << square
		}
	}
	return occupied
}
