package reg

import "unsafe"

// Read reads a 32-bit register at addr.
//
// Read exists for compatibility with generated PAC code that uses uint32
// addresses.
//
//go:nosplit
func Read(addr uint32) uint32 {
	return Read32(uintptr(addr))
}

// Write writes a 32-bit register at addr.
//
// Write exists for compatibility with generated PAC code that uses uint32
// addresses.
//
//go:nosplit
func Write(addr uint32, val uint32) {
	Write32(uintptr(addr), val)
}

// Read8 reads an 8-bit register at addr.
//
//go:nosplit
func Read8(addr uintptr) uint8 {
	return *(*uint8)(unsafe.Pointer(addr)) //nolint:unsafeptr
}

// Write8 writes an 8-bit register at addr.
//
//go:nosplit
func Write8(addr uintptr, val uint8) {
	*(*uint8)(unsafe.Pointer(addr)) = val //nolint:unsafeptr
}

// Read16 reads a 16-bit register at addr.
//
//go:nosplit
func Read16(addr uintptr) uint16 {
	return *(*uint16)(unsafe.Pointer(addr)) //nolint:unsafeptr
}

// Write16 writes a 16-bit register at addr.
//
//go:nosplit
func Write16(addr uintptr, val uint16) {
	*(*uint16)(unsafe.Pointer(addr)) = val //nolint:unsafeptr
}

// Read32 reads a 32-bit register at addr.
//
//go:nosplit
func Read32(addr uintptr) uint32 {
	return *(*uint32)(unsafe.Pointer(addr)) //nolint:unsafeptr
}

// Write32 writes a 32-bit register at addr.
//
//go:nosplit
func Write32(addr uintptr, val uint32) {
	*(*uint32)(unsafe.Pointer(addr)) = val //nolint:unsafeptr
}

// Read64 reads a 64-bit register at addr.
//
//go:nosplit
func Read64(addr uintptr) uint64 {
	return *(*uint64)(unsafe.Pointer(addr)) //nolint:unsafeptr
}

// Write64 writes a 64-bit register at addr.
//
//go:nosplit
func Write64(addr uintptr, val uint64) {
	*(*uint64)(unsafe.Pointer(addr)) = val //nolint:unsafeptr
}

// SetBits32 sets mask bits in a 32-bit register at addr.
//
//go:nosplit
func SetBits32(addr uintptr, mask uint32) {
	Write32(addr, Read32(addr)|mask)
}

// ClearBits32 clears mask bits in a 32-bit register at addr.
//
//go:nosplit
func ClearBits32(addr uintptr, mask uint32) {
	Write32(addr, Read32(addr)&^mask)
}

// MaskWrite32 replaces mask bits in a 32-bit register at addr with val bits.
//
//go:nosplit
func MaskWrite32(addr uintptr, mask, val uint32) {
	Write32(addr, (Read32(addr)&^mask)|(val&mask))
}
