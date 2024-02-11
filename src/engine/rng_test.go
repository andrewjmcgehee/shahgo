package engine

import "testing"

var test_rand_state = rand_state

func TestXORRand(t *testing.T) {
	rand_state = test_rand_state
	expected := uint32(0x0000000067d33a74)
	if actual := XORRand(); expected != actual {
		t.Errorf("XORRand() first call should return 0x%016x but returned 0x%016x", expected, actual)
	}
}

func TestRand(t *testing.T) {
	rand_state = test_rand_state
	expected := uint64(0x41e0fc38fd7a3a74)
	if actual := Rand(); expected != actual {
		t.Errorf("Rand() first call should return 0x%016x but returned 0x%016x", expected, actual)
	}
}

func TestMagicCandidate(t *testing.T) {
	rand_state = test_rand_state
	expected := uint64(0x0060800004203010)
	if actual := MagicCandidate(); expected != actual {
		t.Errorf("MagicCandidate() first call should return 0x%016x but returned 0x%016x", expected, actual)
	}
}
