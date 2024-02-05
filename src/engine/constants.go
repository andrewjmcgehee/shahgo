package engine

// squares
const (
	A1 int = iota
	B1 int = iota
	C1 int = iota
	D1 int = iota
	E1 int = iota
	F1 int = iota
	G1 int = iota
	H1 int = iota
	A2 int = iota
	B2 int = iota
	C2 int = iota
	D2 int = iota
	E2 int = iota
	F2 int = iota
	G2 int = iota
	H2 int = iota
	A3 int = iota
	B3 int = iota
	C3 int = iota
	D3 int = iota
	E3 int = iota
	F3 int = iota
	G3 int = iota
	H3 int = iota
	A4 int = iota
	B4 int = iota
	C4 int = iota
	D4 int = iota
	E4 int = iota
	F4 int = iota
	G4 int = iota
	H4 int = iota
	A5 int = iota
	B5 int = iota
	C5 int = iota
	D5 int = iota
	E5 int = iota
	F5 int = iota
	G5 int = iota
	H5 int = iota
	A6 int = iota
	B6 int = iota
	C6 int = iota
	D6 int = iota
	E6 int = iota
	F6 int = iota
	G6 int = iota
	H6 int = iota
	A7 int = iota
	B7 int = iota
	C7 int = iota
	D7 int = iota
	E7 int = iota
	F7 int = iota
	G7 int = iota
	H7 int = iota
	A8 int = iota
	B8 int = iota
	C8 int = iota
	D8 int = iota
	E8 int = iota
	F8 int = iota
	G8 int = iota
	H8 int = iota
)

// ranks
const RANK_1 uint64 = 0x00000000000000ff
const RANK_2 uint64 = 0x000000000000ff00
const RANK_3 uint64 = 0x0000000000ff0000
const RANK_4 uint64 = 0x00000000ff000000
const RANK_5 uint64 = 0x000000ff00000000
const RANK_6 uint64 = 0x0000ff0000000000
const RANK_7 uint64 = 0x00ff000000000000
const RANK_8 uint64 = 0xff00000000000000

// files
const FILE_A uint64 = 0x0101010101010101
const FILE_B uint64 = 0x0202020202020202
const FILE_C uint64 = 0x0404040404040404
const FILE_D uint64 = 0x0808080808080808
const FILE_E uint64 = 0x1010101010101010
const FILE_F uint64 = 0x2020202020202020
const FILE_G uint64 = 0x4040404040404040
const FILE_H uint64 = 0x8080808080808080

// colors
const WHITE_SQUARES uint64 = 0x55aa55aa55aa55aa
const BLACK_SQUARES uint64 = 0xaa55aa55aa55aa55

// interesting squares
const SQUARE_CENTER uint64 = 0x0000001818000000
const LARGE_SQUARE_CENTER uint64 = 0x00003c3c3c3c0000
const LONG_DIAGONALS uint64 = 0x8142241818244281

// sides of board
const LEFT_HALF uint64 = FILE_A | FILE_B | FILE_C | FILE_D
const RIGHT_HALF uint64 = FILE_E | FILE_F | FILE_G | FILE_H

// special ranks
const PROMOTION_RANKS uint64 = RANK_1 | RANK_8

// all files
var FILES = [8]uint64{FILE_A, FILE_B, FILE_C, FILE_D, FILE_E, FILE_F, FILE_G, FILE_H}

// all ranks
var RANKS = [8]uint64{RANK_1, RANK_2, RANK_3, RANK_4, RANK_5, RANK_6, RANK_7, RANK_8}

// file complements
const NOT_A uint64 = ^FILE_A
const NOT_H uint64 = ^FILE_H
const NOT_AB uint64 = ^FILE_A & ^FILE_B
const NOT_GH uint64 = ^FILE_G & ^FILE_H
