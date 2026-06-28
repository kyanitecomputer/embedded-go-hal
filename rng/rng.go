package rng

import "io"

// Source fills byte slices with random data.
type Source interface {
	io.Reader
}
