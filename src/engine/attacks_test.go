package engine

import "testing"

func TestMaskPawnAttacks(t *testing.T) {
	expected := uint64(0x0000000000000200)
	if actual := MaskPawnAttacks(WHITE, A1); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, A1) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000000000)
	if actual := MaskPawnAttacks(BLACK, A1); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, A1) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000200000000)
	if actual := MaskPawnAttacks(WHITE, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000020000)
	if actual := MaskPawnAttacks(BLACK, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000002800000000)
	if actual := MaskPawnAttacks(WHITE, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, E4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000280000)
	if actual := MaskPawnAttacks(BLACK, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, E4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000004000000000)
	if actual := MaskPawnAttacks(WHITE, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000400000)
	if actual := MaskPawnAttacks(BLACK, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000000000)
	if actual := MaskPawnAttacks(WHITE, H8); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, H8) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0040000000000000)
	if actual := MaskPawnAttacks(BLACK, H8); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, H8) = %v, expected %v", actual, expected)
	}
}

func TestMaskKnightAttacks(t *testing.T) {
	expected := uint64(0x0000000000020400)
	if actual := MaskKnightAttacks(A1); expected != actual {
		t.Errorf("MaskKnightAttacks(A1) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000020400040200)
	if actual := MaskKnightAttacks(A4); expected != actual {
		t.Errorf("MaskKnightAttacks(A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000050800080500)
	if actual := MaskKnightAttacks(B4); expected != actual {
		t.Errorf("MaskKnightAttacks(B4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000284400442800)
	if actual := MaskKnightAttacks(E4); expected != actual {
		t.Errorf("MaskKnightAttacks(E4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000a0100010a000)
	if actual := MaskKnightAttacks(G4); expected != actual {
		t.Errorf("MaskKnightAttacks(G4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000402000204000)
	if actual := MaskKnightAttacks(H4); expected != actual {
		t.Errorf("MaskKnightAttacks(H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0020400000000000)
	if actual := MaskKnightAttacks(H8); expected != actual {
		t.Errorf("MaskKnightAttacks(H8) = %v, expected %v", actual, expected)
	}
}

func TestMaskKingAttacks(t *testing.T) {
	expected := uint64(0x0000000000000302)
	if actual := MaskKingAttacks(A1); expected != actual {
		t.Errorf("MaskKingAttacks(A1) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000302030000)
	if actual := MaskKingAttacks(A4); expected != actual {
		t.Errorf("MaskKingAttacks(A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000003828380000)
	if actual := MaskKingAttacks(E4); expected != actual {
		t.Errorf("MaskKingAttacks(E4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x000000c040c00000)
	if actual := MaskKingAttacks(H4); expected != actual {
		t.Errorf("MaskKingAttacks(H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x40c0000000000000)
	if actual := MaskKingAttacks(H8); expected != actual {
		t.Errorf("MaskKingAttacks(H8) = %v, expected %v", actual, expected)
	}
}

func TestMaskBishopAttacks(t *testing.T) {
	expected := uint64(0x0040201008040200)
	if actual := MaskBishopAttacks(A1); expected != actual {
		t.Errorf("MaskBishopAttacks(A1) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0008040200020400)
	if actual := MaskBishopAttacks(A4); expected != actual {
		t.Errorf("MaskBishopAttacks(A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0010080400040800)
	if actual := MaskBishopAttacks(B4); expected != actual {
		t.Errorf("MaskBishopAttacks(B4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0002442800284400)
	if actual := MaskBishopAttacks(E4); expected != actual {
		t.Errorf("MaskBishopAttacks(E4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0008102000201000)
	if actual := MaskBishopAttacks(G4); expected != actual {
		t.Errorf("MaskBishopAttacks(G4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0010204000402000)
	if actual := MaskBishopAttacks(H4); expected != actual {
		t.Errorf("MaskBishopAttacks(H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0040201008040200)
	if actual := MaskBishopAttacks(H8); expected != actual {
		t.Errorf("MaskBishopAttacks(H8) = %v, expected %v", actual, expected)
	}
}

func TestMaskRookAttacks(t *testing.T) {
	expected := uint64(0x000101010101017e)
	if actual := MaskRookAttacks(A1); expected != actual {
		t.Errorf("MaskRookAttacks(A1) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x000101017e010100)
	if actual := MaskRookAttacks(A4); expected != actual {
		t.Errorf("MaskRookAttacks(A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x000202027c020200)
	if actual := MaskRookAttacks(B4); expected != actual {
		t.Errorf("MaskRookAttacks(B4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x001010106e101000)
	if actual := MaskRookAttacks(E4); expected != actual {
		t.Errorf("MaskRookAttacks(E4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x004040403e404000)
	if actual := MaskRookAttacks(G4); expected != actual {
		t.Errorf("MaskRookAttacks(G4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x008080807e808000)
	if actual := MaskRookAttacks(H4); expected != actual {
		t.Errorf("MaskRookAttacks(H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x7e80808080808000)
	if actual := MaskRookAttacks(H8); expected != actual {
		t.Errorf("MaskRookAttacks(H8) = %v, expected %v", actual, expected)
	}
}
