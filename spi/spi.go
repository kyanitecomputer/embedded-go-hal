package spi

import "time"

// Polarity identifies the idle clock level.
type Polarity uint8

const (
	// IdleLow means the clock is low when idle.
	IdleLow Polarity = iota

	// IdleHigh means the clock is high when idle.
	IdleHigh
)

// Phase identifies the clock edge used to capture data.
type Phase uint8

const (
	// CaptureOnFirstTransition captures data on the first clock transition.
	CaptureOnFirstTransition Phase = iota

	// CaptureOnSecondTransition captures data on the second clock transition.
	CaptureOnSecondTransition
)

// Mode describes SPI clock polarity and phase.
type Mode struct {
	Polarity Polarity
	Phase    Phase
}

// Mode0 is CPOL=0, CPHA=0.
var Mode0 = Mode{Polarity: IdleLow, Phase: CaptureOnFirstTransition}

// Mode1 is CPOL=0, CPHA=1.
var Mode1 = Mode{Polarity: IdleLow, Phase: CaptureOnSecondTransition}

// Mode2 is CPOL=1, CPHA=0.
var Mode2 = Mode{Polarity: IdleHigh, Phase: CaptureOnFirstTransition}

// Mode3 is CPOL=1, CPHA=1.
var Mode3 = Mode{Polarity: IdleHigh, Phase: CaptureOnSecondTransition}

// Operation describes one SPI transaction segment.
//
// If both Write and Read are non-empty, the segment is full duplex. If only
// Write is set, the segment is write-only. If only Read is set, the segment is
// read-only. If Delay is non-zero, the segment delays without transferring data.
type Operation struct {
	Write []byte
	Read  []byte
	Delay time.Duration
}

// Bus transfers bytes on a SPI bus without managing chip select.
type Bus interface {
	Reader
	Writer
	Transferer
	Flusher
}

// Reader reads bytes from a SPI bus.
type Reader interface {
	// Read reads len(buf) bytes from the bus.
	Read(buf []byte) error
}

// Writer writes bytes to a SPI bus.
type Writer interface {
	// Write writes all bytes in buf to the bus.
	Write(buf []byte) error
}

// Transferer performs full-duplex SPI transfers.
type Transferer interface {
	// Transfer writes w while reading into r.
	Transfer(r, w []byte) error

	// TransferInPlace writes and reads using buf as both source and destination.
	TransferInPlace(buf []byte) error
}

// Flusher waits until all queued SPI bus operations are complete.
type Flusher interface {
	// Flush waits until the bus is idle.
	Flush() error
}

// Device performs chip-select-scoped SPI transactions.
type Device interface {
	// Transaction performs operations while the device is selected.
	Transaction(ops []Operation) error
}

// Read performs one read transaction.
func Read(dev Device, buf []byte) error {
	return dev.Transaction([]Operation{{Read: buf}})
}

// Write performs one write transaction.
func Write(dev Device, buf []byte) error {
	return dev.Transaction([]Operation{{Write: buf}})
}

// Transfer performs one full-duplex transaction.
func Transfer(dev Device, r, w []byte) error {
	return dev.Transaction([]Operation{{Read: r, Write: w}})
}

// TransferInPlace performs one in-place full-duplex transaction.
func TransferInPlace(dev Device, buf []byte) error {
	return dev.Transaction([]Operation{{Read: buf, Write: buf}})
}
