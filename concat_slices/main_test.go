package concat_slices

import "testing"

var slices = [][]byte{
	[]byte("my first slice"),
	[]byte("second slice"),
	[]byte("third slice"),
	[]byte("fourth slice"),
	[]byte("fifth slice"),
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
