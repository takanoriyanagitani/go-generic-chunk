package chunk

// BySize creates a ToChunk function that groups items into chunks of a fixed size.
// If the total number of items is not a multiple of chunkSize, the final chunk will contain the remaining items.
func BySize[T any](chunkSize int) ToChunk[T] {
	return bySizeWithReset(chunkSize, newChunk[T])
}
