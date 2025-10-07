package chunk

// BySizeReuse creates a ToChunk function that groups items into chunks of a fixed size.
// If the total number of items is not a multiple of chunkSize, the final chunk will contain the remaining items.
// It re-uses the underlying array for the chunks for efficiency.
// The consumer of the iterator must process the chunk immediately and not hold a reference to it.
func BySizeReuse[T any](chunkSize int) ToChunk[T] {
	return bySizeWithReset(chunkSize, reuseChunk[T])
}
