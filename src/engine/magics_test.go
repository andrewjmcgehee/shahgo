package engine

import "testing"

func TestMagicCandidate(t *testing.T) {
	// magic_rng is a singleton
	if magic_rng == nil {
		magic_rng = NewRNG()
	}
	expected := uint64(0x0060800004203010)
	if actual := magic_candidate(); expected != actual {
		t.Errorf("magic_candidate() first call should return 0x%016x but returned 0x%016x", expected, actual)
	}
}
