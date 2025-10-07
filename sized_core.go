package chunk

type ResetFn[T any] func(chunk []T, chunkSize int) []T

func newChunk[T any](_ []T, chunkSize int) []T {
	return make([]T, 0, chunkSize)
}

func reuseChunk[T any](chunk []T, chunkSize int) []T {
	return chunk[:0]
}

func bySizeWithReset[T any](chunkSize int, reset ResetFn[T]) ToChunk[T] {
	return func(iterator Iterator[T]) ChunkIterator[T] {
		return func(yield func([]T, error) bool) {
			if chunkSize <= 0 {
				yield(nil, ErrChunkSize)
				return
			}

			chunk := make([]T, 0, chunkSize)

			for value, err := range iterator {
				if err != nil {
					yield(nil, err)
					return // Stop on error
				}

				chunk = append(chunk, value)

				if len(chunk) == chunkSize {
					if !yield(chunk, nil) {
						return // Consumer doesn't want more
					}
					chunk = reset(chunk, chunkSize)
				}
			}

			// Yield the last partial chunk
			if len(chunk) > 0 {
				yield(chunk, nil)
			}
		}
	}
}
