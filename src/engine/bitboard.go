package engine

import (
	"fmt"
	"math/bits"
)

type BoardState struct {
	PieceBitboards     [12]uint64
	OccupancyBitboards [3]uint64
	EnPassant          int
	Side               int
	CastleRights       int
}

func NewStartBoardState() *BoardState {
	return &BoardState{
		PieceBitboards:     newPieceBitboards(),
		OccupancyBitboards: newOccupancyBitBoards(),
		EnPassant:          NoSquare,
		Side:               White,
		CastleRights:       WhiteKingside | WhiteQueenside | BlackKingside | BlackQueenside,
	}
}

func NewEmptyBoardState() *BoardState {
	return &BoardState{
		PieceBitboards:     [12]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		OccupancyBitboards: [3]uint64{0, 0, 0},
		EnPassant:          NoSquare,
		Side:               White,
		CastleRights:       0,
	}
}

func SafeCoord(rank, file int) bool {
	return 0 <= rank && rank < 8 && 0 <= file && file < 8
}

func SafeSquare(square int) bool {
	return 0 <= square && square < 64
}

func SafeDim(dim int) bool {
	return 0 <= dim && dim < 8
}

func FileOf(square int) int {
	if !SafeSquare(square) {
		err := fmt.Sprintf("FileOf received unsafe square: %d", square)
		panic(err)
	}
	return square % 8
}

func RankOf(square int) int {
	if !SafeSquare(square) {
		err := fmt.Sprintf("RankOf received unsafe square: %d", square)
		panic(err)
	}
	return square / 8
}

func SquareFrom(rank, file int) int {
	if !SafeDim(rank) {
		err := fmt.Sprintf("SquareFrom received unsafe rank: %d", rank)
		panic(err)
	}
	if !SafeDim(file) {
		err := fmt.Sprintf("SquareFrom received unsafe file: %d", file)
		panic(err)
	}
	return rank*8 + file
}

func TestBit(bitboard uint64, square int) bool {
	if !SafeSquare(square) {
		err := fmt.Sprintf("TestBit received unsafe square: %d", square)
		panic(err)
	}
	return bitboard&(1<<square) != 0
}

func SetBit(bitboard *uint64, square int) {
	if !SafeSquare(square) {
		return
	}
	*bitboard |= (1 << square)
}

func FlipBit(bitboard *uint64, square int) {
	if !SafeSquare(square) {
		return
	}
	*bitboard ^= (1 << square)
}

func PopBit(bitboard *uint64, square int) {
	if TestBit(*bitboard, square) {
		FlipBit(bitboard, square)
	}
}

func CountBits(bitboard uint64) int {
	return bits.OnesCount64(bitboard)
}

func LSBIndex(bitboard uint64) int {
	if bitboard == 0 {
		return 64 // special value to indicate no LSB
	}
	return bits.OnesCount64((bitboard & -bitboard) - 1)
}

func DisplayBitboard(bitboard uint64, stdout bool) string {
	var repr string
	for rank := 0; rank < 8; rank++ {
		repr += fmt.Sprintf("%d ", 8-rank)
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
		fmt.Printf("\n   0x%016x\n\n", bitboard)
	}
	return repr
}

func getPieceIndex(b *BoardState, pieceType uint64) uint64 {
	if b.Side == Black {
		return pieceType + 6
	}
	return pieceType
}

func Move(b *BoardState, pieceType uint64, startSquare, endSquare int) {
	pieceIdx := getPieceIndex(b, pieceType)
	FlipBit(&b.PieceBitboards[pieceIdx], startSquare)
	FlipBit(&b.PieceBitboards[pieceIdx], endSquare)
	FlipBit(&b.OccupancyBitboards[Both], startSquare)
	FlipBit(&b.OccupancyBitboards[Both], endSquare)
	FlipBit(&b.OccupancyBitboards[b.Side], startSquare)
	FlipBit(&b.OccupancyBitboards[b.Side], endSquare)
	if b.Side == White {
		b.Side = Black
	} else {
		b.Side = White
	}
}

func ParseFEN(fen string) *BoardState {
	b := NewEmptyBoardState()
	charIdx := 0
	pieceIndices := pieceBytesToIndicesMap()
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := SquareFrom(rank, file)
			currentChar := fen[charIdx]
			if currentChar == '/' {
				charIdx++
				currentChar = fen[charIdx]
			}
			if (currentChar >= 'A' && currentChar <= 'Z') || (currentChar >= 'a' && currentChar <= 'z') {
				pieceIdx := pieceIndices[currentChar]
				SetBit(&b.PieceBitboards[pieceIdx], square)
			} else if currentChar >= '0' && currentChar <= '9' {
				count := int(currentChar - '0')
				file += count - 1
			}
			charIdx++
		}
	}
	for i := 0; i < 6; i++ {
		b.OccupancyBitboards[White] |= b.PieceBitboards[i]
	}
	for i := 6; i < 12; i++ {
		b.OccupancyBitboards[Black] |= b.PieceBitboards[i]
	}
	b.OccupancyBitboards[Both] = b.OccupancyBitboards[White] | b.OccupancyBitboards[Black]
	// parse side
	charIdx++
	if fen[charIdx] == 'w' {
		b.Side = White
	} else if fen[charIdx] == 'b' {
		b.Side = Black
	}
	// parse castling rights
	charIdx += 2
	for fen[charIdx] != ' ' {
		switch fen[charIdx] {
		case 'K':
			b.CastleRights |= WhiteKingside
		case 'Q':
			b.CastleRights |= WhiteQueenside
		case 'k':
			b.CastleRights |= BlackKingside
		case 'q':
			b.CastleRights |= BlackQueenside
		case '-':
			break
		}
		charIdx++
	}
	// parse en passant
	charIdx++
	if fen[charIdx] != '-' {
		file := fen[charIdx]
		rank := fen[charIdx+1]
		b.EnPassant = SquareFrom(int(rank-'0'), int(file-'a'))
	}
	fmt.Println(fen[charIdx:])
	DisplayBoard(b, true)
	fmt.Printf("%s\n\n", fen)
	return b
}

func DisplayBoard(b *BoardState, stdout bool) string {
	var repr string
	for rank := 0; rank < 8; rank++ {
		repr += fmt.Sprintf("%d ", 8-rank)
		for file := 0; file < 8; file++ {
			square := SquareFrom(rank, file)
			if TestBit(b.OccupancyBitboards[Both], square) {
				if TestBit(b.OccupancyBitboards[White], square) {
					for i := WhitePawns; i <= WhiteKing; i++ {
						if TestBit(b.PieceBitboards[i], square) {
							repr += fmt.Sprintf(" %s", unicodePieceFromCode(i))
						}
					}
				} else {
					for i := BlackPawns; i <= BlackKing; i++ {
						if TestBit(b.PieceBitboards[i], square) {
							repr += fmt.Sprintf(" %s", unicodePieceFromCode(i))
						}
					}
				}
			} else {
				repr += fmt.Sprintf(" .")
			}
		}
		repr += "\n"
	}
	repr += "\n   a b c d e f g h\n"
	if stdout {
		fmt.Println(repr)
		if b.Side == White {
			fmt.Printf("    To Move:  WHITE\n")
		} else {
			fmt.Printf("    To Move:  BLACK\n")
		}
		var castlingRepr string
		castlingMap := getCastlingRightsMap()
		for i := 0; i < 4; i++ {
			index := 1 << i
			if b.CastleRights&index != 0 {
				castlingRepr += castlingMap[index]
			} else {
				castlingRepr += "-"
			}
		}
		fmt.Printf("    Castling: %s\n", castlingRepr)
		var enPassRepr string
		if b.EnPassant != NoSquare {
			rank := RankOf(b.EnPassant)
			file := FileOf(b.EnPassant)
			enPassRepr += fmt.Sprintf("%v%d", 'a', file) + fmt.Sprintf("%v%d", '0', rank)
		} else {
			enPassRepr = "n/a"
		}
		fmt.Printf("    En Pass:  %s\n\n", enPassRepr)
	}
	return repr
}
