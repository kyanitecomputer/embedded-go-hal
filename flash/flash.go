package flash

import (
	"errors"
	"io"
)

// ErrInvalidGeometry reports unusable flash geometry.
var ErrInvalidGeometry = errors.New("flash: invalid geometry")

// Geometry describes common flash memory geometry.
type Geometry struct {
	TotalSize      int64
	EraseBlockSize int
	PageSize       int
	EraseValue     byte
}

// Validate reports whether g contains usable geometry values.
func (g Geometry) Validate() error {
	if g.TotalSize <= 0 || g.EraseBlockSize <= 0 || g.PageSize <= 0 {
		return ErrInvalidGeometry
	}
	if int64(g.EraseBlockSize) > g.TotalSize {
		return ErrInvalidGeometry
	}
	if g.EraseBlockSize%g.PageSize != 0 {
		return ErrInvalidGeometry
	}
	if g.TotalSize%int64(g.EraseBlockSize) != 0 {
		return ErrInvalidGeometry
	}
	return nil
}

// Device provides common flash device metadata.
type Device interface {
	// Geometry returns the flash memory geometry.
	Geometry() Geometry
}

// NOR provides byte-addressed NOR flash operations.
type NOR interface {
	Device
	io.ReaderAt
	io.WriterAt

	// EraseBlock erases the block containing addr.
	EraseBlock(addr int64) error
}

// NANDGeometry describes NAND-specific geometry.
type NANDGeometry struct {
	PagesPerBlock int
	OOBSize       int
	MaxBitflips   int
	MaxBadBlocks  int
}

// NAND provides page-addressed NAND flash operations.
type NAND interface {
	Device

	// NANDGeometry returns NAND-specific geometry.
	NANDGeometry() NANDGeometry

	// ReadPage reads one NAND page and optional out-of-band bytes.
	ReadPage(page int64, data []byte, oob []byte) (bitflips int, err error)

	// WritePage writes one NAND page and optional out-of-band bytes.
	WritePage(page int64, data []byte, oob []byte) error

	// EraseBlock erases one NAND erase block.
	EraseBlock(block int64) error

	// IsBadBlock reports whether block is marked bad.
	IsBadBlock(block int64) (bool, error)

	// MarkBadBlock marks block as bad.
	MarkBadBlock(block int64) error
}
