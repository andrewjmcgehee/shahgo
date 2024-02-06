package engine

// squares
const (
	A8 uint64 = iota
	B8 uint64 = iota
	C8 uint64 = iota
	D8 uint64 = iota
	E8 uint64 = iota
	F8 uint64 = iota
	G8 uint64 = iota
	H8 uint64 = iota
	A7 uint64 = iota
	B7 uint64 = iota
	C7 uint64 = iota
	D7 uint64 = iota
	E7 uint64 = iota
	F7 uint64 = iota
	G7 uint64 = iota
	H7 uint64 = iota
	A6 uint64 = iota
	B6 uint64 = iota
	C6 uint64 = iota
	D6 uint64 = iota
	E6 uint64 = iota
	F6 uint64 = iota
	G6 uint64 = iota
	H6 uint64 = iota
	A5 uint64 = iota
	B5 uint64 = iota
	C5 uint64 = iota
	D5 uint64 = iota
	E5 uint64 = iota
	F5 uint64 = iota
	G5 uint64 = iota
	H5 uint64 = iota
	A4 uint64 = iota
	B4 uint64 = iota
	C4 uint64 = iota
	D4 uint64 = iota
	E4 uint64 = iota
	F4 uint64 = iota
	G4 uint64 = iota
	H4 uint64 = iota
	A3 uint64 = iota
	B3 uint64 = iota
	C3 uint64 = iota
	D3 uint64 = iota
	E3 uint64 = iota
	F3 uint64 = iota
	G3 uint64 = iota
	H3 uint64 = iota
	A2 uint64 = iota
	B2 uint64 = iota
	C2 uint64 = iota
	D2 uint64 = iota
	E2 uint64 = iota
	F2 uint64 = iota
	G2 uint64 = iota
	H2 uint64 = iota
	A1 uint64 = iota
	B1 uint64 = iota
	C1 uint64 = iota
	D1 uint64 = iota
	E1 uint64 = iota
	F1 uint64 = iota
	G1 uint64 = iota
	H1 uint64 = iota
)

// ranks
const RANK_1 uint64 = 0xff00000000000000
const RANK_2 uint64 = 0x00ff000000000000
const RANK_3 uint64 = 0x0000ff0000000000
const RANK_4 uint64 = 0x000000ff00000000
const RANK_5 uint64 = 0x00000000ff000000
const RANK_6 uint64 = 0x0000000000ff0000
const RANK_7 uint64 = 0x000000000000ff00
const RANK_8 uint64 = 0x00000000000000ff

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
const WHITE_SQUARES uint64 = 0xaa55aa55aa55aa55
const BLACK_SQUARES uint64 = 0x55aa55aa55aa55aa

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
