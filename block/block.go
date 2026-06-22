package block

import (
	"errors"
	"io"
)

// ErrInvalidGeometry reports an unusable logical block layout.
var ErrInvalidGeometry = errors.New("block: invalid geometry")

// ErrOutOfRange reports a block or offset outside the device.
var ErrOutOfRange = errors.New("block: out of range")

// ErrUnaligned reports an operation that violates write alignment.
var ErrUnaligned = errors.New("block: unaligned operation")

// Device provides logical block access.
type Device interface {
	io.ReaderAt
	io.WriterAt

	// BlockSize returns the number of bytes in each logical block.
	BlockSize() int

	// BlockCount returns the number of logical blocks.
	BlockCount() int

	// WriteSize returns the required write alignment in bytes.
	WriteSize() int

	// EraseBlock erases one logical block.
	EraseBlock(block int) error
}

// Syncer commits buffered block-device state to storage.
type Syncer interface {
	// Sync commits buffered state to storage.
	Sync() error
}

// Validate reports whether dev exposes usable geometry.
func Validate(dev Device) error {
	if dev == nil {
		return ErrInvalidGeometry
	}
	if dev.BlockSize() <= 0 || dev.BlockCount() <= 0 || dev.WriteSize() <= 0 {
		return ErrInvalidGeometry
	}
	if dev.BlockSize()%dev.WriteSize() != 0 {
		return ErrInvalidGeometry
	}
	return nil
}

// CheckRange validates a block operation range.
func CheckRange(dev Device, block int, off int, n int) error {
	if err := Validate(dev); err != nil {
		return err
	}
	if block < 0 || block >= dev.BlockCount() || off < 0 || n < 0 || off > dev.BlockSize()-n {
		return ErrOutOfRange
	}
	return nil
}

// CheckWrite validates a block write range and alignment.
func CheckWrite(dev Device, block int, off int, n int) error {
	if err := CheckRange(dev, block, off, n); err != nil {
		return err
	}
	if off%dev.WriteSize() != 0 || n%dev.WriteSize() != 0 {
		return ErrUnaligned
	}
	return nil
}

// CheckByteRange validates a byte-addressed operation range.
func CheckByteRange(dev Device, off int64, n int) error {
	if err := Validate(dev); err != nil {
		return err
	}
	total := int64(dev.BlockSize()) * int64(dev.BlockCount())
	if off < 0 || n < 0 || off > total-int64(n) {
		return ErrOutOfRange
	}
	return nil
}

// CheckByteWrite validates a byte-addressed write range and alignment.
func CheckByteWrite(dev Device, off int64, n int) error {
	if err := CheckByteRange(dev, off, n); err != nil {
		return err
	}
	writeSize := int64(dev.WriteSize())
	if off%writeSize != 0 || int64(n)%writeSize != 0 {
		return ErrUnaligned
	}
	return nil
}
