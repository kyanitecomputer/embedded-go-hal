package can

import (
	"errors"
	"testing"
)

func TestIDConstructors(t *testing.T) {
	if _, err := StandardID(0x7ff); err != nil {
		t.Fatalf("StandardID() = %v, want nil", err)
	}
	if _, err := StandardID(0x800); !errors.Is(err, ErrInvalidFrame) {
		t.Fatalf("StandardID() = %v, want %v", err, ErrInvalidFrame)
	}
	if _, err := ExtendedID(0x1fffffff); err != nil {
		t.Fatalf("ExtendedID() = %v, want nil", err)
	}
	if _, err := ExtendedID(0x20000000); !errors.Is(err, ErrInvalidFrame) {
		t.Fatalf("ExtendedID() = %v, want %v", err, ErrInvalidFrame)
	}
}

func TestFrameValidate(t *testing.T) {
	if err := (Frame{ID: ID{Value: 0x123}, Data: make([]byte, 8)}).Validate(); err != nil {
		t.Fatalf("Validate() = %v, want nil", err)
	}
	if err := (Frame{ID: ID{Value: 0x800}}).Validate(); !errors.Is(err, ErrInvalidFrame) {
		t.Fatalf("Validate() = %v, want %v", err, ErrInvalidFrame)
	}
	if err := (Frame{ID: ID{Value: 0x123}, Data: make([]byte, 9)}).Validate(); !errors.Is(err, ErrInvalidFrame) {
		t.Fatalf("Validate() = %v, want %v", err, ErrInvalidFrame)
	}
}
