package engine

import "testing"

func TestMaskPawnAttacks(t *testing.T) {
	cases := map[string]struct {
		turn     int
		square   int
		expected uint64
	}{
		"a1 white to move": {White, A1, uint64(0x0002000000000000)},
		"a1 black to move": {Black, A1, uint64(0x0000000000000000)},
		"a4 white to move": {White, A4, uint64(0x0000000002000000)},
		"a4 black to move": {Black, A4, uint64(0x0000020000000000)},
		"e4 white to move": {White, E4, uint64(0x0000000028000000)},
		"e4 black to move": {Black, E4, uint64(0x0000280000000000)},
		"h4 white to move": {White, H4, uint64(0x0000000040000000)},
		"h4 black to move": {Black, H4, uint64(0x0000400000000000)},
		"h8 white to move": {White, H8, uint64(0x0000000000000000)},
		"h8 black to move": {Black, H8, uint64(0x0000000000004000)},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := MaskPawnAttacks(tc.turn, tc.square); tc.expected != actual {
				t.Errorf("MaskPawnAttacks(%v, %v) = 0x%016x, expected 0x%016x",
					humanReadbleTurn(tc.turn), humanReadbleSquare(tc.square), actual, tc.expected)
			}
		})
	}
}

func TestMaskKnightAttacks(t *testing.T) {
	cases := map[string]struct {
		square   int
		expected uint64
	}{
		"a8": {A8, uint64(0x0000000000020400)},
		"a5": {A5, uint64(0x0000020400040200)},
		"b5": {B5, uint64(0x0000050800080500)},
		"e5": {E5, uint64(0x0000284400442800)},
		"g5": {G5, uint64(0x0000a0100010a000)},
		"h5": {H5, uint64(0x0000402000204000)},
		"h1": {H1, uint64(0x0020400000000000)},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := MaskKnightAttacks(tc.square); tc.expected != actual {
				t.Errorf("MaskKnightAttacks(%v) = 0x%016x, expected 0x%016x", humanReadbleSquare(tc.square), actual, tc.expected)
			}
		})
	}
}

func TestMaskKingAttacks(t *testing.T) {
	cases := map[string]struct {
		square   int
		expected uint64
	}{
		"a8": {A8, uint64(0x0000000000000302)},
		"a5": {A5, uint64(0x0000000302030000)},
		"e5": {E5, uint64(0x0000003828380000)},
		"h5": {H5, uint64(0x000000c040c00000)},
		"h1": {H1, uint64(0x40c0000000000000)},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := MaskKingAttacks(tc.square); tc.expected != actual {
				t.Errorf("MaskKingAttacks(%v) = 0x%016x, expected 0x%016x", humanReadbleSquare(tc.square), actual, tc.expected)
			}
		})
	}
}

func TestBishopRays(t *testing.T) {
	cases := map[string]struct {
		square   int
		edges    bool
		occupied uint64
		expected uint64
	}{
		"E5 with edges":                 {E5, true, 0, uint64(0x0182442800284482)},
		"E5 without edges":              {E5, false, 0, uint64(0x0002442800284400)},
		"E5 with blockers on C3 and D6": {E5, true, 0x0000040000080000, uint64(0x0080442800284080)},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := BishopRays(tc.square, tc.edges, tc.occupied); tc.expected != actual {
				t.Errorf("BishopRays(%v, %v, 0x%016x) = 0x%016x, expected 0x%016x",
					humanReadbleSquare(tc.square), tc.edges, tc.occupied, actual, tc.expected)
			}
		})
	}
}

func TestMaskBishopAttacks(t *testing.T) {
	cases := map[string]struct {
		square   int
		expected uint64
	}{
		"A8": {A8, uint64(0x0040201008040200)},
		"A5": {A5, uint64(0x0008040200020400)},
		"B5": {B5, uint64(0x0010080400040800)},
		"E5": {E5, uint64(0x0002442800284400)},
		"G5": {G5, uint64(0x0008102000201000)},
		"H5": {H5, uint64(0x0010204000402000)},
		"H1": {H1, uint64(0x0040201008040200)},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := MaskBishopAttacks(tc.square); tc.expected != actual {
				t.Errorf("MaskBishopAttacks(%v) = 0x%016x, expected 0x%016x", humanReadbleSquare(tc.square), actual, tc.expected)
			}
		})
	}
}

func TestRookRays(t *testing.T) {
	cases := map[string]struct {
		square   int
		edges    bool
		occupied uint64
		expected uint64
	}{
		"E5 with edges":                 {E5, true, 0, uint64(0x10101010ef101010)},
		"E5 without edges":              {E5, false, 0, uint64(0x001010106e101000)},
		"D3 with blockers on C3 and D6": {D3, true, 0x0000040000080000, uint64(0x0808f40808080000)},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := RookRays(tc.square, tc.edges, tc.occupied); tc.expected != actual {
				t.Errorf("RookRays(%v, %v, 0x%016x) = 0x%016x, expected 0x%016x",
					humanReadbleSquare(tc.square), tc.edges, tc.occupied, actual, tc.expected)
			}
		})
	}
}

func TestMaskRookAttacks(t *testing.T) {
	cases := map[string]struct {
		square   int
		expected uint64
	}{
		"A8": {A8, uint64(0x000101010101017e)},
		"A5": {A5, uint64(0x000101017e010100)},
		"B5": {B5, uint64(0x000202027c020200)},
		"E5": {E5, uint64(0x001010106e101000)},
		"G5": {G5, uint64(0x004040403e404000)},
		"H5": {H5, uint64(0x008080807e808000)},
		"H1": {H1, uint64(0x7e80808080808000)},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := MaskRookAttacks(tc.square); tc.expected != actual {
				t.Errorf("MaskRookAttacks(%v) = 0x%016x, expected 0x%016x", humanReadbleSquare(tc.square), actual, tc.expected)
			}
		})
	}
}

func TestSetOccupancy(t *testing.T) {
	attackMap := RookRays(A1, false, 0)
	bits := CountBits(attackMap)
	cases := map[string]struct {
		index    int
		expected uint64
	}{
		"index 1":    {1, 0x0000000000000100},
		"index 3650": {3650, 0x7200000000010000},
		"index 4095": {4095, attackMap},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if actual := SetOccupancy(tc.index, bits, attackMap); tc.expected != actual {
				t.Errorf("SetOccupancy(%v, %v, 0x%016x) = 0x%016x, expected 0x%016x", tc.index, bits, attackMap, actual, tc.expected)
			}
		})
	}
}
