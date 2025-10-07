// Package chunk provides generic functions for splitting iterators into chunks.
package chunk

import (
	"iter"
)

// Iterator defines a generic iterator that yields a value of type T and an error.
// It follows the pattern of Go's `iter.Seq2` type.
type Iterator[T any] iter.Seq2[T, error]

// ChunkIterator defines a generic iterator that yields a slice of T values and an error.
// It follows the pattern of Go's `iter.Seq2` type.
type ChunkIterator[T any] iter.Seq2[[]T, error]

// ToChunk is a function that converts an Iterator into a ChunkIterator.
// This is the core transformation function provided by the package.
type ToChunk[T any] func(Iterator[T]) ChunkIterator[T]
