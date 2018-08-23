package concat_slices

import "testing"

var slices = [][]byte{
	[]byte("my first slice"),
	[]byte("second slice"),
	[]byte("third slice"),
	[]byte("fourth slice"),
	[]byte("fifth slice"),
	[]byte("looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong slice"),
}

var B []byte

func BenchmarkConcatCopyPreAllocate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = concatCopyPreAllocate(slices)
	}
}

func BenchmarkConcatAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = concatAppend(slices)
	}
}

func BenchmarkConcatWriteBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = concatWriteBuffer(slices)
	}
}

func testEq(a, b []byte) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestConcatWriteBuffer(t *testing.T) {
	B = concatWriteBuffer(slices)
	expected := concatAppend(slices)
	if !testEq(B, expected) {
		t.Errorf("expected length: %d, got: %d", len(expected), len(B))
		t.Error(B)
	}
}
