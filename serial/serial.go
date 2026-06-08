package serial

import "io"

// Port is a byte-oriented serial device.
type Port interface {
	io.Reader
	io.Writer
}

// BytePort is a serial device optimized for single-byte operations.
type BytePort interface {
	io.ByteReader
	io.ByteWriter
}

// Flusher flushes buffered serial output.
type Flusher interface {
	// Flush writes buffered output to the device.
	Flush() error
}

// ReadReady reports whether a serial port can be read without blocking.
type ReadReady interface {
	// ReadReady reports whether a read can complete immediately.
	ReadReady() (bool, error)
}

// WriteReady reports whether a serial port can be written without blocking.
type WriteReady interface {
	// WriteReady reports whether a write can complete immediately.
	WriteReady() (bool, error)
}
