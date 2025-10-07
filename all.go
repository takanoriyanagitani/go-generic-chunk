package chunk

// All collects all chunks from a ChunkIterator into a slice.
// It creates a copy of each chunk, so it is safe to use with iterators that reuse the chunk slice.
func All[T any](it ChunkIterator[T]) ([][]T, error) {
	var chunks [][]T
	var errEncountered error

	it(func(chunk []T, err error) bool {
		if err != nil {
			errEncountered = err
			return false
		}
		// Create a copy of the chunk, as the underlying array might be reused.
		chunkCopy := make([]T, len(chunk))
		copy(chunkCopy, chunk)
		chunks = append(chunks, chunkCopy)
		return true
	})

	if errEncountered != nil {
		return nil, errEncountered
	}

	return chunks, nil
}
