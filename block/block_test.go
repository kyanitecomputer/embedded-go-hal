package block

import (
	"errors"
	"testing"
)

type testDevice struct {
	blockSize  int
	blockCount int
	writeSize  int
}

func (d testDevice) BlockSize() int  { return d.blockSize }
func (d testDevice) BlockCount() int { return d.blockCount }
func (d testDevice) WriteSize() int  { return d.writeSize }
func (d testDevice) ReadAt([]byte, int64) (int, error) {
	return 0, nil
}
func (d testDevice) WriteAt([]byte, int64) (int, error) {
	return 0, nil
}
func (d testDevice) EraseBlock(int) error { return nil }

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		dev  Device
		want error
	}{
		{name: "valid", dev: testDevice{blockSize: 512, blockCount: 4, writeSize: 16}},
		{name: "nil", want: ErrInvalidGeometry},
		{name: "zero block size", dev: testDevice{blockCount: 4, writeSize: 16}, want: ErrInvalidGeometry},
		{name: "unaligned write size", dev: testDevice{blockSize: 512, blockCount: 4, writeSize: 24}, want: ErrInvalidGeometry},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.dev); !errors.Is(err, tt.want) {
				t.Fatalf("Validate() = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestCheckWrite(t *testing.T) {
	dev := testDevice{blockSize: 512, blockCount: 4, writeSize: 16}
	if err := CheckWrite(dev, 1, 16, 32); err != nil {
		t.Fatalf("CheckWrite() = %v, want nil", err)
	}
	if err := CheckWrite(dev, 4, 0, 16); !errors.Is(err, ErrOutOfRange) {
		t.Fatalf("CheckWrite() = %v, want %v", err, ErrOutOfRange)
	}
	if err := CheckWrite(dev, 1, 8, 16); !errors.Is(err, ErrUnaligned) {
		t.Fatalf("CheckWrite() = %v, want %v", err, ErrUnaligned)
	}
}

func TestCheckByteWrite(t *testing.T) {
	dev := testDevice{blockSize: 512, blockCount: 4, writeSize: 16}
	if err := CheckByteWrite(dev, 16, 32); err != nil {
		t.Fatalf("CheckByteWrite() = %v, want nil", err)
	}
	if err := CheckByteWrite(dev, 2048, 16); !errors.Is(err, ErrOutOfRange) {
		t.Fatalf("CheckByteWrite() = %v, want %v", err, ErrOutOfRange)
	}
	if err := CheckByteWrite(dev, 8, 16); !errors.Is(err, ErrUnaligned) {
		t.Fatalf("CheckByteWrite() = %v, want %v", err, ErrUnaligned)
	}
}
