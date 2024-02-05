package engine

import (
	"fmt"
	"math/bits"
)

func SafeCoord(rank int, file int) bool {
	return 0 <= rank && rank < 8 && 0 <= file && file < 8
}

func UnsafeSquare(square int) bool {
	return square < 0 || square >= 64
}

func UnsafeDim(dim int) bool {
	return dim < 0 || dim >= 8
}

func FileOf(square int) int {
	if UnsafeSquare(square) {
		err := fmt.Sprintf("file_of received unsafe square: %d", square)
		panic(err)
	}
	return square % 8
}

func RankOf(square int) int {
	if UnsafeSquare(square) {
		err := fmt.Sprintf("rank_of received unsafe square: %d", square)
		panic(err)
	}
	return square / 8
}

func SquareFrom(rank int, file int) int {
	if UnsafeDim(rank) {
		err := fmt.Sprintf("square_from received unsafe rank: %d", rank)
		panic(err)
	}
	if UnsafeDim(file) {
		err := fmt.Sprintf("square_from received unsafe file: %d", file)
		panic(err)
	}
	return rank*8 + file
}

func TestBit(bitboard uint64, square int) bool {
	if UnsafeSquare(square) {
		err := fmt.Sprintf("test_bit received unsafe square: %d", square)
		panic(err)
	}
	return bitboard&(1<<square) != 0
}

func SetBit(bitboard *uint64, square int) {
	if UnsafeSquare(square) {
		return
	}
	*bitboard |= (1 << square)
}

func FlipBit(bitboard *uint64, square int) {
	if UnsafeSquare(square) {
		return
	}
	*bitboard ^= (1 << square)
}

func CountBits(bitboard uint64) int {
	return bits.OnesCount64(bitboard)
}

func MSBIndex(bitboard uint64) int {
	if bitboard == 0 {
		return -1
	}
	return bits.LeadingZeros64(bitboard)
}

func Display(bitboard uint64, stdout bool) string {
	var repr string
	for rank := 7; rank >= 0; rank-- {
		repr += fmt.Sprintf("%d ", rank+1)
		for file := 0; file < 8; file++ {
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
