package chunk

import (
	"errors"
	"reflect"
	"testing"
)

func TestBySizeReuse(t *testing.T) {
	testCases := []struct {
		name        string
		chunkSize   int
		iterator    Iterator[int]
		expected    [][]int
		expectedErr error
	}{
		{
			name:      "regular case",
			chunkSize: 3,
			iterator:  newIntIterator(0, 10),
			expected: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9},
			},
		},
		{
			name:      "empty iterator",
			chunkSize: 3,
			iterator:  newIntIterator(0, 0),
			expected:  [][]int{},
		},
		{
			name:      "smaller than chunk size",
			chunkSize: 5,
			iterator:  newIntIterator(0, 3),
			expected: [][]int{
				{0, 1, 2},
			},
		},
		{
			name:      "exact multiple of chunk size",
			chunkSize: 2,
			iterator:  newIntIterator(0, 6),
			expected: [][]int{
				{0, 1},
				{2, 3},
				{4, 5},
			},
		},
		{
			name:        "zero chunk size",
			chunkSize:   0,
			iterator:    newIntIterator(0, 5),
			expectedErr: ErrChunkSize,
		},
		{
			name:        "negative chunk size",
			chunkSize:   -1,
			iterator:    newIntIterator(0, 5),
			expectedErr: ErrChunkSize,
		},
		{
			name:        "error from iterator",
			chunkSize:   3,
			iterator:    newErrorIterator[int](ErrIterator),
			expectedErr: ErrIterator,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			toChunk := BySizeReuse[int](tc.chunkSize)
			chunkIt := toChunk(tc.iterator)

			chunks := make([][]int, 0)
			var errEncountered error

			chunkIt(func(chunk []int, err error) bool {
				if err != nil {
					errEncountered = err
					return false
				}
				// Create a copy of the chunk, as the underlying array is reused.
				chunkCopy := make([]int, len(chunk))
				copy(chunkCopy, chunk)
				chunks = append(chunks, chunkCopy)
				return true
			})

			if tc.expectedErr != nil {
				if !errors.Is(errEncountered, tc.expectedErr) {
					t.Errorf("expected error '%v', got '%v'", tc.expectedErr, errEncountered)
				}
			} else if errEncountered != nil {
				t.Errorf("unexpected error: %v", errEncountered)
			}

			if tc.expectedErr == nil {
				if !reflect.DeepEqual(chunks, tc.expected) {
					t.Errorf("expected chunks %v, got %v", tc.expected, chunks)
				}
			}
		})
	}
}
