package engine

import (
	"fmt"
	"math/bits"
)

func SafeCoord(rank uint64, file uint64) bool {
	return rank < 8 && file < 8
}

func SafeSquare(square uint64) bool {
	return square < 64
}

func SafeDim(dim uint64) bool {
	return dim < 8
}

func FileOf(square uint64) uint64 {
	if !SafeSquare(square) {
		err := fmt.Sprintf("file_of received unsafe square: %d", square)
		panic(err)
	}
	return square % 8
}

func RankOf(square uint64) uint64 {
	if !SafeSquare(square) {
		err := fmt.Sprintf("rank_of received unsafe square: %d", square)
		panic(err)
	}
	return square / 8
}

func SquareFrom(rank uint64, file uint64) uint64 {
	if !SafeDim(rank) {
		err := fmt.Sprintf("square_from received unsafe rank: %d", rank)
		panic(err)
	}
	if !SafeDim(file) {
		err := fmt.Sprintf("square_from received unsafe file: %d", file)
		panic(err)
	}
	return rank*8 + file
}

func TestBit(bitboard uint64, square uint64) bool {
	if !SafeSquare(square) {
		err := fmt.Sprintf("test_bit received unsafe square: %d", square)
		panic(err)
	}
	return bitboard&(1<<square) != 0
}

func SetBit(bitboard *uint64, square uint64) {
	if !SafeSquare(square) {
		return
	}
	*bitboard |= (1 << square)
}

func FlipBit(bitboard *uint64, square uint64) {
	if !SafeSquare(square) {
		return
	}
	*bitboard ^= (1 << square)
}

func PopBit(bitboard *uint64, square uint64) {
	if TestBit(*bitboard, square) {
		FlipBit(bitboard, square)
	}
}

func CountBits(bitboard uint64) uint64 {
	return uint64(bits.OnesCount64(bitboard))
}

func LSBIndex(bitboard uint64) uint64 {
	if bitboard == 0 {
		return 64 // special value to indicate no LSB
	}
	return uint64(bits.OnesCount64((bitboard & -bitboard) - 1))
}

func Display(bitboard uint64, stdout bool) string {
	var repr string
	for rank := uint64(0); rank < 8; rank++ {
		repr += fmt.Sprintf("%d ", 8-rank)
		for file := uint64(0); file < 8; file++ {
			if TestBit(bitboard, SquareFrom(rank, file)) {
				repr += fmt.Sprintf(" 1")
			} else {
				repr += fmt.Sprintf(" 0")
			}
		}
		repr += "\n"
	}
	repr += "\n   a b c d e f g h"
	if stdout {
		fmt.Println(repr)
		fmt.Printf("\n   0x%016x\n", bitboard)
	}
	return repr
}
