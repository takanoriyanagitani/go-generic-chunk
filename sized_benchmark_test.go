package chunk

import (
	"testing"
)

func BenchmarkBySize(b *testing.B) {
	it := newIntIterator(0, 1000)
	toChunk := BySize[int](100)
	chunkIt := toChunk(it)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chunkIt(func(chunk []int, err error) bool {
			// Do nothing with the chunk, just consume it.
			return true
		})
	}
}

func BenchmarkBySizeReuse(b *testing.B) {
	it := newIntIterator(0, 1000)
	toChunk := BySizeReuse[int](100)
	chunkIt := toChunk(it)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chunkIt(func(chunk []int, err error) bool {
			// Do nothing with the chunk, just consume it.
			return true
		})
	}
}
