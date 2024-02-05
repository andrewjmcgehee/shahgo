package engine

import "testing"

func TestMaskPawnAttacks(t *testing.T) {
	expected := uint64(0x0000000200000000)
	if actual := MaskPawnAttacks(WHITE, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000020000)
	if actual := MaskPawnAttacks(BLACK, A4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, A4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000004000000000)
	if actual := MaskPawnAttacks(WHITE, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000400000)
	if actual := MaskPawnAttacks(BLACK, H4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, H4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000002800000000)
	if actual := MaskPawnAttacks(WHITE, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(WHITE, E4) = %v, expected %v", actual, expected)
	}
	expected = uint64(0x0000000000280000)
	if actual := MaskPawnAttacks(BLACK, E4); expected != actual {
		t.Errorf("MaskPawnAttacks(BLACK, E4) = %v, expected %v", actual, expected)
	}
}
