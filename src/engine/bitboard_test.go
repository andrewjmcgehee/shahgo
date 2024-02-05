package engine

import "testing"

func TestSafeCoord(t *testing.T) {
	expected := true
	if actual := SafeCoord(0, 0); expected != actual {
		t.Errorf("SafeCoord(0, 0) = %v, expected %v", actual, expected)
	}
	if actual := SafeCoord(7, 0); expected != actual {
		t.Errorf("SafeCoord(7, 0) = %v, expected %v", actual, expected)
	}
	if actual := SafeCoord(7, 7); expected != actual {
		t.Errorf("SafeCoord(7, 7) = %v, expected %v", actual, expected)
	}
	if actual := SafeCoord(0, 7); expected != actual {
		t.Errorf("SafeCoord(0, 7) = %v, expected %v", actual, expected)
	}
	expected = false
	if actual := SafeCoord(-1, 0); expected != actual {
		t.Errorf("SafeCoord(-1, 0) = %v, expected %v", actual, expected)
	}
	if actual := SafeCoord(0, -1); expected != actual {
		t.Errorf("SafeCoord(0, -1) = %v, expected %v", actual, expected)
	}
	if actual := SafeCoord(8, 0); expected != actual {
		t.Errorf("SafeCoord(8, 0) = %v, expected %v", actual, expected)
	}
	if actual := SafeCoord(0, 8); expected != actual {
		t.Errorf("SafeCoord(0, 8) = %v, expected %v", actual, expected)
	}

}

func TestUnsafeSquare(t *testing.T) {
	expected := true
	if actual := UnsafeSquare(-1); expected != actual {
		t.Errorf("UnsafeSquare(-1) = %v, expected %v", actual, expected)
	}
	expected = false
	if actual := UnsafeSquare(0); expected != actual {
		t.Errorf("UnsafeSquare(0) = %v, expected %v", actual, expected)
	}
	if actual := UnsafeSquare(63); expected != actual {
		t.Errorf("UnsafeSquare(63) = %v, expected %v", actual, expected)
	}
	expected = true
	if actual := UnsafeSquare(64); expected != actual {
		t.Errorf("UnsafeSquare(64) = %v, expected %v", actual, expected)
	}
}

func TestUnsafeDim(t *testing.T) {
	expected := true
	if actual := UnsafeDim(-1); expected != actual {
		t.Errorf("UnsafeDim(-1) = %v, expected %v", actual, expected)
	}
	expected = false
	if actual := UnsafeDim(0); expected != actual {
		t.Errorf("UnsafeDim(0) = %v, expected %v", actual, expected)
	}
	if actual := UnsafeDim(7); expected != actual {
		t.Errorf("UnsafeDim(7) = %v, expected %v", actual, expected)
	}
	expected = true
	if actual := UnsafeDim(8); expected != actual {
		t.Errorf("UnsafeDim(8) = %v, expected %v", actual, expected)
	}
}

func TestFileOf(t *testing.T) {
	// UnsafeSquare already tested, only checking happy paths
	expected := 0
	if actual := FileOf(8); expected != actual {
		t.Errorf("FileOf(8) = %v, expected %v", actual, expected)
	}
	expected = 1
	if actual := FileOf(9); expected != actual {
		t.Errorf("FileOf(9) = %v, expected %v", actual, expected)
	}
	expected = 7
	if actual := FileOf(15); expected != actual {
		t.Errorf("FileOf(15) = %v, expected %v", actual, expected)
	}
	expected = 0
	if actual := FileOf(16); expected != actual {
		t.Errorf("FileOf(16) = %v, expected %v", actual, expected)
	}
}

func TestRankOf(t *testing.T) {
	// UnsafeSquare already tested, only checking happy paths
	expected := 0
	if actual := RankOf(1); expected != actual {
		t.Errorf("RankOf(1) = %v, expected %v", actual, expected)
	}
	expected = 1
	if actual := RankOf(8); expected != actual {
		t.Errorf("RankOf(8) = %v, expected %v", actual, expected)
	}
	if actual := RankOf(15); expected != actual {
		t.Errorf("RankOf(15) = %v, expected %v", actual, expected)
	}
	expected = 2
	if actual := RankOf(16); expected != actual {
		t.Errorf("RankOf(16) = %v, expected %v", actual, expected)
	}
}

func TestSquareFrom(t *testing.T) {
	// UnsafeDim already tested, only checking happy paths
	expected := 0
	if actual := SquareFrom(0, 0); expected != actual {
		t.Errorf("SquareFrom(0, 0) = %v, expected %v", actual, expected)
	}
	expected = 1
	if actual := SquareFrom(0, 1); expected != actual {
		t.Errorf("SquareFrom(0, 1) = %v, expected %v", actual, expected)
	}
	expected = 8
	if actual := SquareFrom(1, 0); expected != actual {
		t.Errorf("SquareFrom(1, 0) = %v, expected %v", actual, expected)
	}
	expected = 63
	if actual := SquareFrom(7, 7); expected != actual {
		t.Errorf("SquareFrom(7, 7) = %v, expected %v", actual, expected)
	}
}

func TestTestBit(t *testing.T) {
	// UnsafeSquare already tested, only checking happy paths
	bitboard := RANK_1
	for square := 0; square < 64; square++ {
		expected := true
		if square >= 8 {
			expected = false
		}
		if actual := TestBit(bitboard, square); expected != actual {
			t.Errorf("TestBit(0x00000000000000ff, %d) = %v, expected %v", square, actual, expected)
		}
	}
	bitboard = RANK_2
	for square := 0; square < 64; square++ {
		expected := true
		if square < 8 || square >= 16 {
			expected = false
		}
		if actual := TestBit(bitboard, square); expected != actual {
			t.Errorf("TestBit(0x000000000000ff00, %d) = %v, expected %v", square, actual, expected)
		}
	}
}

func TestFlipBit(t *testing.T) {
	expected := RANK_1
	var actual uint64 = 0x00000000000000bf
	if FlipBit(&actual, 6); expected != actual {
		t.Errorf("FlipBit(0x00000000000000bf, 7) = %x, expected %x", actual, expected)
	}
	expected = uint64(0)
	actual = 1
	if FlipBit(&actual, 0); expected != actual {
		t.Errorf("FlipBit(0x0000000000000000, 0) = %x, expected %x", actual, expected)
	}
}

func TestSetBit(t *testing.T) {
	expected := RANK_1
	var actual uint64 = 0x00000000000000bf
	if SetBit(&actual, 6); expected != actual {
		t.Errorf("SetBit(0x00000000000000bf, 7) = %x, expected %x", actual, expected)
	}
	expected = uint64(1)
	actual = 1
	if SetBit(&actual, 0); expected != actual {
		t.Errorf("SetBit(0x0000000000000001, 0) = %x, expected %x", actual, expected)
	}
}

func TestDisplay(t *testing.T) {
	row_0 := "1 0 1 0 1 0 1 0\n"
	row_1 := "0 1 0 1 0 1 0 1\n"
	expected := "8  " + row_0
	expected += "7  " + row_1
	expected += "6  " + row_0
	expected += "5  " + row_1
	expected += "4  " + row_0
	expected += "3  " + row_1
	expected += "2  " + row_0
	expected += "1  " + row_1
	expected += "\n   a b c d e f g h"
	if actual := Display(WHITE_SQUARES, false); expected != actual {
		t.Errorf("Display(0x55aa55aa55aa55aa, false) = %v, expected %v", actual, expected)
	}
}
