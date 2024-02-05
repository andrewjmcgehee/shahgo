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
