package engine

// colors
const (
	White = iota
	Black
	Both
)

// squares
const (
	A8 = iota
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	NoSquare
)

// ranks
const Rank1 uint64 = 0xff00000000000000
const Rank2 uint64 = 0x00ff000000000000
const Rank3 uint64 = 0x0000ff0000000000
const Rank4 uint64 = 0x000000ff00000000
const Rank5 uint64 = 0x00000000ff000000
const Rank6 uint64 = 0x0000000000ff0000
const Rank7 uint64 = 0x000000000000ff00
const Rank8 uint64 = 0x00000000000000ff

// files
const FileA uint64 = 0x0101010101010101
const FileB uint64 = 0x0202020202020202
const FileC uint64 = 0x0404040404040404
const FileD uint64 = 0x0808080808080808
const FileE uint64 = 0x1010101010101010
const FileF uint64 = 0x2020202020202020
const FileG uint64 = 0x4040404040404040
const FileH uint64 = 0x8080808080808080

// colors
const WhiteSquares uint64 = 0xaa55aa55aa55aa55
const BlackSquares uint64 = 0x55aa55aa55aa55aa

// interesting squares
const SquareCenter uint64 = 0x0000001818000000
const LargeSquareCenter uint64 = 0x00003c3c3c3c0000
const LongDiagonals uint64 = 0x8142241818244281

// sides of board
const LeftHalf uint64 = FileA | FileB | FileC | FileD
const RightHalf uint64 = FileE | FileF | FileG | FileH

// special ranks
const PromotionRanks uint64 = Rank1 | Rank8

// castling
const (
	WhiteKingside = 1 << iota
	WhiteQueenside
	BlackKingside
	BlackQueenside
)

func getCastlingRightsMap() map[int]string {
	return map[int]string{
		1: "K",
		2: "Q",
		4: "k",
		8: "q",
	}
}

// piece types
const (
	WhitePawns = iota
	WhiteKnights
	WhiteBishops
	WhiteRooks
	WhiteQueen
	WhiteKing
	BlackPawns
	BlackKnights
	BlackBishops
	BlackRooks
	BlackQueen
	BlackKing
)
const (
	Pawn = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func newPieceBitboards() [12]uint64 {
	return [12]uint64{
		Rank2,              // white pawns
		0x2400000000000000, // white knights
		0x4200000000000000, // white bishops
		0x8100000000000000, // white rooks
		0x0800000000000000, // white queen
		0x1000000000000000, // white king,
		Rank7,              // black pawns
		0x0000000000000024, // black knights
		0x0000000000000042, // black bishops
		0x0000000000000081, // black rooks
		0x0000000000000008, // black queen
		0x0000000000000010, // black king,
	}
}

func newOccupancyBitBoards() [3]uint64 {
	pieces := newPieceBitboards()
	white := pieces[0] | pieces[1] | pieces[2] | pieces[3] | pieces[4] | pieces[5]
	black := pieces[6] | pieces[7] | pieces[8] | pieces[9] | pieces[10] | pieces[11]
	return [3]uint64{
		white,
		black,
		white | black,
	}
}

// piece printing
func pieceBytesToIndicesMap() map[byte]int {
	return map[byte]int{
		'P': 0,
		'N': 1,
		'B': 2,
		'R': 3,
		'Q': 4,
		'K': 5,
		'p': 6,
		'n': 7,
		'b': 8,
		'r': 9,
		'q': 10,
		'k': 11,
	}
}

func unicodePieceStrings() []string {
	return []string{
		"♟",
		"♞",
		"♝",
		"♜",
		"♛",
		"♚",
		"♙",
		"♘",
		"♗",
		"♖",
		"♕",
		"♔",
	}
}

func unicodePieceFromCode(code int) string {
	return unicodePieceStrings()[code]
}

// FEN codes
const FenEmpty = "8/8/8/8/8/8/8/8 w - -"
const FenStart = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const FenTricky = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
const FenKill = "rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1"
const FenCMK = "r2q1rk1/ppp2ppp/2n1bn2/2b1p3/3pP3/3P1NPP/PPP1NPB1/R1BQ1RK1 b - - 0 9"

// all files
func allFiles() [8]uint64 {
	return [8]uint64{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}
}

// all ranks
func allRanks() [8]uint64 {
	return [8]uint64{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}
}

// file complements
const NotFileA uint64 = ^FileA
const NotFileH uint64 = ^FileH
const NotFilesAB uint64 = ^FileA & ^FileB
const NotFilesGH uint64 = ^FileG & ^FileH

func bishopOccupancyCounts() [64]int {
	return [64]int{
		6, 5, 5, 5, 5, 5, 5, 6,
		5, 5, 5, 5, 5, 5, 5, 5,
		5, 5, 7, 7, 7, 7, 5, 5,
		5, 5, 7, 9, 9, 7, 5, 5,
		5, 5, 7, 9, 9, 7, 5, 5,
		5, 5, 7, 7, 7, 7, 5, 5,
		5, 5, 5, 5, 5, 5, 5, 5,
		6, 5, 5, 5, 5, 5, 5, 6,
	}
}

func bishopMagics() [64]uint64 {
	return [64]uint64{
		0x0040040844404084,
		0x002004208a004208,
		0x0010190041080202,
		0x0108060845042010,
		0x0581104180800210,
		0x2112080446200010,
		0x1080820820060210,
		0x03c0808410220200,
		0x0004050404440404,
		0x0000021001420088,
		0x24d0080801082102,
		0x0001020a0a020400,
		0x0000040308200402,
		0x0004011002100800,
		0x0401484104104005,
		0x0801010402020200,
		0x00400210c3880100,
		0x0404022024108200,
		0x0810018200204102,
		0x0004002801a02003,
		0x0085040820080400,
		0x810102c808880400,
		0x000e900410884800,
		0x8002020480840102,
		0x0220200865090201,
		0x2010100a02021202,
		0x0152048408022401,
		0x0020080002081110,
		0x4001001021004000,
		0x800040400a011002,
		0x00e4004081011002,
		0x001c004001012080,
		0x8004200962a00220,
		0x8422100208500202,
		0x2000402200300c08,
		0x8646020080080080,
		0x80020a0200100808,
		0x2010004880111000,
		0x623000a080011400,
		0x42008c0340209202,
		0x0209188240001000,
		0x400408a884001800,
		0x00110400a6080400,
		0x1840060a44020800,
		0x0090080104000041,
		0x0201011000808101,
		0x1a2208080504f080,
		0x8012020600211212,
		0x0500861011240000,
		0x0180806108200800,
		0x4000020e01040044,
		0x300000261044000a,
		0x0802241102020002,
		0x0020906061210001,
		0x5a84841004010310,
		0x0004010801011c04,
		0x000a010109502200,
		0x0000004a02012000,
		0x500201010098b028,
		0x8040002811040900,
		0x0028000010020204,
		0x06000020202d0240,
		0x8918844842082200,
		0x4010011029020020,
	}
}

func rookOccupancyCounts() [64]int {
	return [64]int{
		12, 11, 11, 11, 11, 11, 11, 12,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		11, 10, 10, 10, 10, 10, 10, 11,
		12, 11, 11, 11, 11, 11, 11, 12,
	}
}

func rookMagics() [64]uint64 {
	return [64]uint64{
		0x8a80104000800020,
		0x0140002000100040,
		0x02801880a0017001,
		0x0100081001000420,
		0x0200020010080420,
		0x03001c0002010008,
		0x8480008002000100,
		0x2080088004402900,
		0x0000800098204000,
		0x2024401000200040,
		0x0100802000801000,
		0x0120800800801000,
		0x0208808088000400,
		0x0002802200800400,
		0x2200800100020080,
		0x0801000060821100,
		0x0080044006422000,
		0x0100808020004000,
		0x12108a0010204200,
		0x0140848010000802,
		0x0481828014002800,
		0x8094004002004100,
		0x4010040010010802,
		0x0000020008806104,
		0x0100400080208000,
		0x2040002120081000,
		0x0021200680100081,
		0x0020100080080080,
		0x0002000a00200410,
		0x0000020080800400,
		0x0080088400100102,
		0x0080004600042881,
		0x4040008040800020,
		0x0440003000200801,
		0x0004200011004500,
		0x0188020010100100,
		0x0014800401802800,
		0x2080040080800200,
		0x0124080204001001,
		0x0200046502000484,
		0x0480400080088020,
		0x1000422010034000,
		0x0030200100110040,
		0x0000100021010009,
		0x2002080100110004,
		0x0202008004008002,
		0x0020020004010100,
		0x2048440040820001,
		0x0101002200408200,
		0x0040802000401080,
		0x4008142004410100,
		0x02060820c0120200,
		0x0001001004080100,
		0x020c020080040080,
		0x2935610830022400,
		0x0044440041009200,
		0x0280001040802101,
		0x2100190040002085,
		0x80c0084100102001,
		0x4024081001000421,
		0x00020030a0244872,
		0x0012001008414402,
		0x02006104900a0804,
		0x0001004081002402,
	}
}

func humanReadbleSquare(square int) string {
	return string([]byte{byte(FileOf(square) + 'a'), byte(RankOf(square) + '1')})
}

func humanReadbleTurn(turn int) string {
	if turn == White {
		return "white"
	}
	return "black"
}
