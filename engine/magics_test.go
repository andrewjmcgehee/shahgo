package engine

import "testing"

func TestMagicCandidate(t *testing.T) {
	// magicRNG is a singleton
	if magicRNG == nil {
		magicRNG = NewRNG()
	}
	expected := uint64(0x0060800004203010)
	if actual := magicCandidate(); expected != actual {
		t.Errorf("magicCandidate() first call should return 0x%016x but returned 0x%016x", expected, actual)
	}
}
