package can

import "errors"

// ErrInvalidFrame reports an invalid CAN frame.
var ErrInvalidFrame = errors.New("can: invalid frame")

// ID identifies a CAN frame.
type ID struct {
	Value    uint32
	Extended bool
}

// StandardID returns an 11-bit CAN identifier.
func StandardID(id uint16) (ID, error) {
	if id > 0x7ff {
		return ID{}, ErrInvalidFrame
	}
	return ID{Value: uint32(id)}, nil
}

// ExtendedID returns a 29-bit CAN identifier.
func ExtendedID(id uint32) (ID, error) {
	if id > 0x1fffffff {
		return ID{}, ErrInvalidFrame
	}
	return ID{Value: id, Extended: true}, nil
}

// Frame is one CAN frame.
type Frame struct {
	ID     ID
	Data   []byte
	Remote bool
}

// Validate reports whether f is a valid CAN 2.0 frame.
func (f Frame) Validate() error {
	if f.ID.Extended {
		if f.ID.Value > 0x1fffffff {
			return ErrInvalidFrame
		}
	} else if f.ID.Value > 0x7ff {
		return ErrInvalidFrame
	}
	if len(f.Data) > 8 {
		return ErrInvalidFrame
	}
	return nil
}

// Receiver receives CAN frames.
type Receiver interface {
	// Receive receives one CAN frame.
	Receive() (Frame, error)
}

// Transmitter transmits CAN frames.
type Transmitter interface {
	// Transmit transmits one CAN frame.
	Transmit(Frame) error
}

// Device receives and transmits CAN frames.
type Device interface {
	Receiver
	Transmitter
}
