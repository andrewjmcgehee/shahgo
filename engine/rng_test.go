package engine

import "testing"

func TestNewRNG(t *testing.T) {
	r := NewRNG()
	expected := uint32(1804289383)
	if actual := r.randomState; expected != actual {
		t.Errorf("NewRNG() randomState should be 0x%016x but is 0x%016x", expected, actual)
	}
}

func TestXORRand(t *testing.T) {
	r := NewRNG()
	expected := uint32(0x0000000067d33a74)
	if actual := r.XORRand(); expected != actual {
		t.Errorf("XORRand() first call should return 0x%016x but returned 0x%016x", expected, actual)
	}
}

func TestRand(t *testing.T) {
	r := NewRNG()
	expected := uint64(0x41e0fc38fd7a3a74)
	if actual := r.Rand(); expected != actual {
		t.Errorf("Rand() first call should return 0x%016x but returned 0x%016x", expected, actual)
	}
}
