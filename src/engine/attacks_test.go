package engine

import "testing"

func TestMaskPawnAttacks(t *testing.T) {
	expected := uint64(0x0002000000000000)
	if actual := MaskPawnAttacks(WHITE, A1); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000000000)
	if actual := MaskPawnAttacks(BLACK, A1); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, A1) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000002000000)
	if actual := MaskPawnAttacks(WHITE, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000020000000000)
	if actual := MaskPawnAttacks(BLACK, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, A4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000028000000)
	if actual := MaskPawnAttacks(WHITE, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000280000000000)
	if actual := MaskPawnAttacks(BLACK, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, E4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000040000000)
	if actual := MaskPawnAttacks(WHITE, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000400000000000)
	if actual := MaskPawnAttacks(BLACK, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, H4) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000000000)
	if actual := MaskPawnAttacks(WHITE, H8); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000000004000)
	if actual := MaskPawnAttacks(BLACK, H8); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, H8) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestMaskKnightAttacks(t *testing.T) {
	expected := uint64(0x0000000000020400)
	if actual := MaskKnightAttacks(A8); expected != actual {
		t.Errorf("MaskKnightAttacks(A8) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000020400040200)
	if actual := MaskKnightAttacks(A5); expected != actual {
		t.Errorf("MaskKnightAttacks(A5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000050800080500)
	if actual := MaskKnightAttacks(B5); expected != actual {
		t.Errorf("MaskKnightAttacks(B5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000284400442800)
	if actual := MaskKnightAttacks(E5); expected != actual {
		t.Errorf("MaskKnightAttacks(E5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000a0100010a000)
	if actual := MaskKnightAttacks(G5); expected != actual {
		t.Errorf("MaskKnightAttacks(G5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000402000204000)
	if actual := MaskKnightAttacks(H5); expected != actual {
		t.Errorf("MaskKnightAttacks(H5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0020400000000000)
	if actual := MaskKnightAttacks(H1); expected != actual {
		t.Errorf("MaskKnightAttacks(H1) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestMaskKingAttacks(t *testing.T) {
	expected := uint64(0x0000000000000302)
	if actual := MaskKingAttacks(A8); expected != actual {
		t.Errorf("MaskKingAttacks(A8) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000000302030000)
	if actual := MaskKingAttacks(A5); expected != actual {
		t.Errorf("MaskKingAttacks(A5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0000003828380000)
	if actual := MaskKingAttacks(E5); expected != actual {
		t.Errorf("MaskKingAttacks(E5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x000000c040c00000)
	if actual := MaskKingAttacks(H5); expected != actual {
		t.Errorf("MaskKingAttacks(H5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x40c0000000000000)
	if actual := MaskKingAttacks(H1); expected != actual {
		t.Errorf("MaskKingAttacks(H1) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestBishopRays(t *testing.T) {
	expected := uint64(0x0182442800284482)
	if actual := BishopRays(E5, true, 0); expected != actual {
		t.Errorf("BishopRays(E5, true) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0002442800284400)
	if actual := BishopRays(E5, false, 0); expected != actual {
		t.Errorf("BishopRays(E5, false) = 0x%016x, expected 0x%016x", actual, expected)
	}
	blocker := uint64(0)
	SetBit(&blocker, C3)
	SetBit(&blocker, D6)
	expected = uint64(0x0080442800284080)
	if actual := BishopRays(E5, true, blocker); expected != actual {
		t.Errorf("BishopRays(E5, true, 0x%016x) = 0x%016x, expected 0x%016x", blocker, actual, expected)
	}
}

func TestRelevantBishopOccupants(t *testing.T) {
	expected := uint64(0x0040201008040200)
	if actual := RelevantBishopOccupants(A8); expected != actual {
		t.Errorf("RelevantBishopOccupants(A8) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0008040200020400)
	if actual := RelevantBishopOccupants(A5); expected != actual {
		t.Errorf("RelevantBishopOccupants(A5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0010080400040800)
	if actual := RelevantBishopOccupants(B5); expected != actual {
		t.Errorf("RelevantBishopOccupants(B5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0002442800284400)
	if actual := RelevantBishopOccupants(E5); expected != actual {
		t.Errorf("RelevantBishopOccupants(E5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0008102000201000)
	if actual := RelevantBishopOccupants(G5); expected != actual {
		t.Errorf("RelevantBishopOccupants(G5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0010204000402000)
	if actual := RelevantBishopOccupants(H5); expected != actual {
		t.Errorf("MaskBishopAttacks(H5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x0040201008040200)
	if actual := RelevantBishopOccupants(H1); expected != actual {
		t.Errorf("RelevantBishopOccupants(H1) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestRookRays(t *testing.T) {
	expected := uint64(0x10101010ef101010)
	if actual := RookRays(E5, true, 0); expected != actual {
		t.Errorf("RookRays(E5, true) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x001010106e101000)
	if actual := RookRays(E5, false, 0); expected != actual {
		t.Errorf("RookRays(E5, false) = 0x%016x, expected 0x%016x", actual, expected)
	}
	blocker := uint64(0)
	SetBit(&blocker, C3)
	SetBit(&blocker, D6)
	expected = uint64(0x0808f40808080000)
	if actual := RookRays(D3, true, blocker); expected != actual {
		t.Errorf("RookRays(D3, true, 0x%016x) = 0x%016x, expected 0x%016x", blocker, actual, expected)
	}
}

func TestRelevantRookOccupants(t *testing.T) {
	expected := uint64(0x000101010101017e)
	if actual := RelevantRookOccupants(A8); expected != actual {
		t.Errorf("RelevantRookOccupants(A8) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x000101017e010100)
	if actual := RelevantRookOccupants(A5); expected != actual {
		t.Errorf("RelevantRookOccupants(A5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x000202027c020200)
	if actual := RelevantRookOccupants(B5); expected != actual {
		t.Errorf("RelevantRookOccupants(B5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x001010106e101000)
	if actual := RelevantRookOccupants(E5); expected != actual {
		t.Errorf("RelevantRookOccupants(E5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x004040403e404000)
	if actual := RelevantRookOccupants(G5); expected != actual {
		t.Errorf("RelevantRookOccupants(G5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x008080807e808000)
	if actual := RelevantRookOccupants(H5); expected != actual {
		t.Errorf("RelevantRookOccupants(H5) = 0x%016x, expected 0x%016x", actual, expected)
	}
	expected = uint64(0x7e80808080808000)
	if actual := RelevantRookOccupants(H1); expected != actual {
		t.Errorf("RelevantRookOccupants(H1) = 0x%016x, expected 0x%016x", actual, expected)
	}
}

func TestSetOccupancy(t *testing.T) {
	expected := uint64(0x0000000000000100)
	attack_map := RookRays(A1, false, 0)
	bits := CountBits(attack_map)
	if actual := SetOccupancy(1, bits, attack_map); expected != actual {
		t.Errorf("SetOccupancy(1, %v, 0x%016x) = 0x%016x, expected 0x%016x", bits, attack_map, actual, expected)
	}
	expected = uint64(0x7200000000010000)
	if actual := SetOccupancy(3650, bits, attack_map); expected != actual {
		t.Errorf("SetOccupancy(1, %v, 0x%016x) = 0x%016x, expected 0x%016x", bits, attack_map, actual, expected)
	}
	expected = attack_map
	if actual := SetOccupancy(4095, bits, attack_map); expected != actual {
		t.Errorf("SetOccupancy(1, %v, 0x%016x) = 0x%016x, expected 0x%016x", bits, attack_map, actual, expected)
	}
}
