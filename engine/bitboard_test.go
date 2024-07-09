package engine

import (
	"fmt"
	"testing"
)

func TestSafeCoord(t *testing.T) {
	cases := map[string]struct {
		rank, file int
		safe       bool
	}{
		"0,0":  {0, 0, true},
		"7,0":  {7, 0, true},
		"7,7":  {7, 7, true},
		"0,7":  {0, 7, true},
		"8,0":  {8, 0, false},
		"0,8":  {0, 8, false},
		"-1,0": {-1, 0, false},
		"0,-1": {0, -1, false},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := SafeCoord(tc.rank, tc.file); tc.safe != actual {
				t.Errorf("SafeCoord(%d, %d) = %v, expected %v", tc.rank, tc.file, actual, tc.safe)
			}
		})
	}
}

func TestSafeSquare(t *testing.T) {
	cases := map[string]struct {
		square int
		safe   bool
	}{
		"0":  {0, true},
		"63": {63, true},
		"-1": {-1, false},
		"64": {64, false},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := SafeSquare(tc.square); tc.safe != actual {
				t.Errorf("SafeSquare(%d) = %v, expected %v", tc.square, actual, tc.safe)
			}
		})
	}
}

func TestSafeDim(t *testing.T) {
	cases := map[string]struct {
		dim  int
		safe bool
	}{
		"0":  {0, true},
		"7":  {7, true},
		"8":  {8, false},
		"-1": {-1, false},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := SafeDim(tc.dim); tc.safe != actual {
				t.Errorf("SafeDim(%d) = %v, expected %v", tc.dim, actual, tc.safe)
			}
		})
	}
}

func TestFileOf(t *testing.T) {
	cases := map[string]struct {
		square, file int
	}{
		"8":  {8, 0},
		"9":  {9, 1},
		"15": {15, 7},
		"16": {16, 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := FileOf(tc.square); tc.file != actual {
				t.Errorf("FileOf(%d) = %v, expected %v", tc.square, actual, tc.file)
			}
		})
	}
}

func TestRankOf(t *testing.T) {
	// UnsafeSquare already tested, only checking happy paths
	cases := map[string]struct {
		square, rank int
	}{
		"1":  {1, 0},
		"8":  {8, 1},
		"15": {15, 1},
		"16": {16, 2},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := RankOf(tc.square); tc.rank != actual {
				t.Errorf("RankOf(%d) = %v, expected %v", tc.square, actual, tc.rank)
			}
		})
	}
}

func TestSquareFrom(t *testing.T) {
	// UnsafeDim already tested, only checking happy paths
	cases := map[string]struct {
		rank, file, square int
	}{
		"0,0": {0, 0, 0},
		"0,1": {0, 1, 1},
		"1,0": {1, 0, 8},
		"7,7": {7, 7, 63},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := SquareFrom(tc.rank, tc.file); tc.square != actual {
				t.Errorf("SquareFrom(%d, %d) = %v, expected %v", tc.rank, tc.file, actual, tc.square)
			}
		})
	}
}

func TestTestBit(t *testing.T) {
	// UnsafeSquare already tested, only checking happy paths
	bitboard := Rank8
	for square := 0; square < 64; square++ {
		expected := true
		if square >= 8 {
			expected = false
		}
		t.Run(fmt.Sprintf("rank 8 square %d", square), func(t *testing.T) {
			if actual := TestBit(bitboard, square); expected != actual {
				t.Errorf("TestBit(0x%016x, %d) = %v, expected %v", bitboard, square, actual, expected)
			}
		})
	}
	bitboard = Rank7
	for square := 0; square < 64; square++ {
		expected := true
		if square < 8 || square >= 16 {
			expected = false
		}
		t.Run(fmt.Sprintf("rank 7 square %d", square), func(t *testing.T) {
			if actual := TestBit(bitboard, square); expected != actual {
				t.Errorf("TestBit(0x%016x, %d) = %v, expected %v", bitboard, square, actual, expected)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	cases := map[string]struct {
		pre, expected uint64
		square        int
	}{
		"Rank 8": {0x00000000000000bf, Rank8, 6},
		"1":      {1, 1, 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := tc.pre
			SetBit(&actual, tc.square)
			if tc.expected != actual {
				t.Errorf("SetBit(&(0x%016x), %d) = 0x%016x, expected 0x%016x", tc.pre, tc.square, actual, tc.expected)
			}
		})
	}
}

func TestFlipBit(t *testing.T) {
	cases := map[string]struct {
		pre, expected uint64
		square        int
	}{
		"Rank 8": {0x00000000000000bf, Rank8, 6},
		"1":      {1, 0, 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := tc.pre
			FlipBit(&actual, tc.square)
			if tc.expected != actual {
				t.Errorf("FlipBit(&(0x%016x), %d) = 0x%016x, expected 0x%016x", tc.pre, tc.square, actual, tc.expected)
			}
		})
	}
}

func TestCountBits(t *testing.T) {
	cases := map[string]struct {
		bitboard uint64
		expected int
	}{
		"5":   {5, 2},
		"MAX": {uint64(1)<<63 + (uint64(1)<<63 - 1), 64},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := CountBits(tc.bitboard); tc.expected != actual {
				t.Errorf("CountBits(0x%016x) = %v, expected %v", tc.bitboard, actual, tc.expected)
			}
		})
	}
}

func TestLSBIndex(t *testing.T) {
	cases := map[string]struct {
		bitboard uint64
		expected int
	}{
		"0":     {0, 64},
		"1<<63": {uint64(1) << 63, 63},
		"1<<1":  {uint64(1) << 1, 1},
		"7":     {7, 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := LSBIndex(tc.bitboard); tc.expected != actual {
				t.Errorf("LSBIndex(0x%016x) = %v, expected %v", tc.bitboard, actual, tc.expected)
			}
		})
	}
}

func TestDisplayBitboard(t *testing.T) {
	row0 := "1 0 1 0 1 0 1 0\n"
	row1 := "0 1 0 1 0 1 0 1\n"
	expected := "8  " + row0
	expected += "7  " + row1
	expected += "6  " + row0
	expected += "5  " + row1
	expected += "4  " + row0
	expected += "3  " + row1
	expected += "2  " + row0
	expected += "1  " + row1
	expected += "\n   a b c d e f g h"
	if actual := DisplayBitboard(WhiteSquares, false); expected != actual {
		t.Errorf("Display(0x55aa55aa55aa55aa, false) = %v, expected %v", actual, expected)
	}
}
