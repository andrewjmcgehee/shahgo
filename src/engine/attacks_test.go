package engine

import "testing"

func TestMaskPawnAttacks(t *testing.T) {
	expected := uint64(0x0000000000000200)
	if actual := MaskPawnAttacks(WHITE, A1); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000000000)
	if actual := MaskPawnAttacks(BLACK, A1); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000200000000)
	if actual := MaskPawnAttacks(WHITE, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000020000)
	if actual := MaskPawnAttacks(BLACK, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000002800000000)
	if actual := MaskPawnAttacks(WHITE, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000280000)
	if actual := MaskPawnAttacks(BLACK, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000004000000000)
	if actual := MaskPawnAttacks(WHITE, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000400000)
	if actual := MaskPawnAttacks(BLACK, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000000000)
	if actual := MaskPawnAttacks(WHITE, H8); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0040000000000000)
	if actual := MaskPawnAttacks(BLACK, H8); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestMaskKnightAttacks(t *testing.T) {
	expected := uint64(0x0000000000020400)
	if actual := MaskKnightAttacks(A1); expected != actual {
		t.Errorf("MaskKnightAttacks(A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000020400040200)
	if actual := MaskKnightAttacks(A4); expected != actual {
		t.Errorf("MaskKnightAttacks(A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000050800080500)
	if actual := MaskKnightAttacks(B4); expected != actual {
		t.Errorf("MaskKnightAttacks(B4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000284400442800)
	if actual := MaskKnightAttacks(E4); expected != actual {
		t.Errorf("MaskKnightAttacks(E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000a0100010a000)
	if actual := MaskKnightAttacks(G4); expected != actual {
		t.Errorf("MaskKnightAttacks(G4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000402000204000)
	if actual := MaskKnightAttacks(H4); expected != actual {
		t.Errorf("MaskKnightAttacks(H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0020400000000000)
	if actual := MaskKnightAttacks(H8); expected != actual {
		t.Errorf("MaskKnightAttacks(H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestMaskKingAttacks(t *testing.T) {
	expected := uint64(0x0000000000000302)
	if actual := MaskKingAttacks(A1); expected != actual {
		t.Errorf("MaskKingAttacks(A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000302030000)
	if actual := MaskKingAttacks(A4); expected != actual {
		t.Errorf("MaskKingAttacks(A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000003828380000)
	if actual := MaskKingAttacks(E4); expected != actual {
		t.Errorf("MaskKingAttacks(E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x000000c040c00000)
	if actual := MaskKingAttacks(H4); expected != actual {
		t.Errorf("MaskKingAttacks(H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x40c0000000000000)
	if actual := MaskKingAttacks(H8); expected != actual {
		t.Errorf("MaskKingAttacks(H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestBishopRays(t *testing.T) {
	expected := uint64(0x0182442800284482)
	if actual := BishopRays(E4, true, 0); expected != actual {
		t.Errorf("BishopRays(E4, true) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0002442800284400)
	if actual := BishopRays(E4, false, 0); expected != actual {
		t.Errorf("BishopRays(E4, false) = 0x%016x, expected 0x%016x", actual, expected)
	}
	blocker := uint64(0)
	SetBit(&blocker, C6)
	SetBit(&blocker, D3)
	expected = uint64(0x0080442800284080)
	if actual := BishopRays(E4, true, blocker); expected != actual {
		t.Errorf("BishopRays(E4, true, 0x%016x) = 0x%016x, expected 0x%016x", blocker, actual, expected)
	}
}

func TestRelevantBishopOccupants(t *testing.T) {
	expected := uint64(0x0040201008040200)
	if actual := RelevantBishopOccupants(A1); expected != actual {
		t.Errorf("RelevantBishopOccupants(A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0008040200020400)
	if actual := RelevantBishopOccupants(A4); expected != actual {
		t.Errorf("RelevantBishopOccupants(A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0010080400040800)
	if actual := RelevantBishopOccupants(B4); expected != actual {
		t.Errorf("RelevantBishopOccupants(B4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0002442800284400)
	if actual := RelevantBishopOccupants(E4); expected != actual {
		t.Errorf("RelevantBishopOccupants(E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0008102000201000)
	if actual := RelevantBishopOccupants(G4); expected != actual {
		t.Errorf("RelevantBishopOccupants(G4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0010204000402000)
	if actual := RelevantBishopOccupants(H4); expected != actual {
		t.Errorf("MaskBishopAttacks(H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0040201008040200)
	if actual := RelevantBishopOccupants(H8); expected != actual {
		t.Errorf("RelevantBishopOccupants(H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestRookRays(t *testing.T) {
	expected := uint64(0x10101010ef101010)
	if actual := RookRays(E4, true, 0); expected != actual {
		t.Errorf("RookRays(E4, true) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x001010106e101000)
	if actual := RookRays(E4, false, 0); expected != actual {
		t.Errorf("RookRays(E4, false) = 0x%016x, expected 0x%016x", actual, expected)
	}
	blocker := uint64(0)
	SetBit(&blocker, C6)
	SetBit(&blocker, D3)
	expected = uint64(0x0808f40808080000)
	if actual := RookRays(D6, true, blocker); expected != actual {
		t.Errorf("RookRays(D6, true, 0x%016x) = 0x%016x, expected 0x%016x", blocker, actual, expected)
	}
}

func TestRelevantRookOccupants(t *testing.T) {
	expected := uint64(0x000101010101017e)
	if actual := RelevantRookOccupants(A1); expected != actual {
		t.Errorf("RelevantRookOccupants(A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x000101017e010100)
	if actual := RelevantRookOccupants(A4); expected != actual {
		t.Errorf("RelevantRookOccupants(A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x000202027c020200)
	if actual := RelevantRookOccupants(B4); expected != actual {
		t.Errorf("RelevantRookOccupants(B4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x001010106e101000)
	if actual := RelevantRookOccupants(E4); expected != actual {
		t.Errorf("RelevantRookOccupants(E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x004040403e404000)
	if actual := RelevantRookOccupants(G4); expected != actual {
		t.Errorf("RelevantRookOccupants(G4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x008080807e808000)
	if actual := RelevantRookOccupants(H4); expected != actual {
		t.Errorf("RelevantRookOccupants(H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x7e80808080808000)
	if actual := RelevantRookOccupants(H8); expected != actual {
		t.Errorf("RelevantRookOccupants(H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
}
