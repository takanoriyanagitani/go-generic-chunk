package chunk

import "errors"

var (
	ErrChunkSize = errors.New("chunk size must be positive")
	ErrIterator  = errors.New("iterator error")
)
