package reg

import (
	"testing"
	"unsafe"
)

func TestReadWriteWidths(t *testing.T) {
	t.Run("8", func(t *testing.T) {
		var v uint8
		addr := uintptr(unsafe.Pointer(&v))
		Write8(addr, 0xa5)
		if got := Read8(addr); got != 0xa5 {
			t.Fatalf("Read8() = %#x, want 0xa5", got)
		}
	})

	t.Run("16", func(t *testing.T) {
		var v uint16
		addr := uintptr(unsafe.Pointer(&v))
		Write16(addr, 0xa55a)
		if got := Read16(addr); got != 0xa55a {
			t.Fatalf("Read16() = %#x, want 0xa55a", got)
		}
	})

	t.Run("32", func(t *testing.T) {
		var v uint32
		addr := uintptr(unsafe.Pointer(&v))
		Write32(addr, 0xa55af00d)
		if got := Read32(addr); got != 0xa55af00d {
			t.Fatalf("Read32() = %#x, want 0xa55af00d", got)
		}
	})

	t.Run("64", func(t *testing.T) {
		var v uint64
		addr := uintptr(unsafe.Pointer(&v))
		Write64(addr, 0xa55af00d11223344)
		if got := Read64(addr); got != 0xa55af00d11223344 {
			t.Fatalf("Read64() = %#x, want 0xa55af00d11223344", got)
		}
	})
}

func TestReadWriteCompatibility(t *testing.T) {
	if unsafe.Sizeof(uintptr(0)) < 8 {
		t.Skip("uint32 compatibility test requires test address below 4 GiB")
	}

	var v uint32
	addr := uintptr(unsafe.Pointer(&v))
	if addr > uintptr(^uint32(0)) {
		t.Skip("test address does not fit generated PAC uint32 API")
	}

	Write(uint32(addr), 0x12345678)
	if got := Read(uint32(addr)); got != 0x12345678 {
		t.Fatalf("Read() = %#x, want 0x12345678", got)
	}
}

func TestBitHelpers(t *testing.T) {
	var v uint32 = 0b1010
	addr := uintptr(unsafe.Pointer(&v))

	SetBits32(addr, 0b0101)
	if got := Read32(addr); got != 0b1111 {
		t.Fatalf("after SetBits32 got %#b, want 0b1111", got)
	}

	ClearBits32(addr, 0b0011)
	if got := Read32(addr); got != 0b1100 {
		t.Fatalf("after ClearBits32 got %#b, want 0b1100", got)
	}

	MaskWrite32(addr, 0b0110, 0b0010)
	if got := Read32(addr); got != 0b1010 {
		t.Fatalf("after MaskWrite32 got %#b, want 0b1010", got)
	}
}
