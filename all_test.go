package chunk

import (
	"errors"
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	it := newIntIterator(0, 10)
	toChunk := BySize[int](3)
	chunkIt := toChunk(it)

	chunks, err := All(chunkIt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9},
	}

	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected %v, got %v", expected, chunks)
	}
}

func TestAllWithReuse(t *testing.T) {
	it := newIntIterator(0, 10)
	toChunk := BySizeReuse[int](3)
	chunkIt := toChunk(it)

	chunks, err := All(chunkIt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9},
	}

	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected %v, got %v", expected, chunks)
	}
}

func TestAllWithError(t *testing.T) {
	it := newErrorIterator[int](ErrIterator)
	toChunk := BySize[int](3)
	chunkIt := toChunk(it)

	_, err := All(chunkIt)
	if !errors.Is(err, ErrIterator) {
		t.Errorf("expected error %v, got %v", ErrIterator, err)
	}
}

func BenchmarkAll(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := newIntIterator(0, 1000)
		toChunk := BySize[int](100)
		chunkIt := toChunk(it)
		_, _ = All(chunkIt)
	}
}

func BenchmarkAllWithReuse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := newIntIterator(0, 1000)
		toChunk := BySizeReuse[int](100)
		chunkIt := toChunk(it)
		_, _ = All(chunkIt)
	}
}
