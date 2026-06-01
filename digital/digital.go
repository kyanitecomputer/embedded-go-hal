package digital

import "context"

// State identifies a digital pin state.
type State bool

const (
	// Low identifies a low pin state.
	Low State = false

	// High identifies a high pin state.
	High State = true
)

// Edge identifies a digital input transition.
type Edge uint8

const (
	// RisingEdge identifies a low-to-high input transition.
	RisingEdge Edge = iota + 1

	// FallingEdge identifies a high-to-low input transition.
	FallingEdge

	// BothEdges identifies either input transition.
	BothEdges
)

// InputPin reads a digital input pin.
type InputPin interface {
	// IsHigh reports whether the pin is currently high.
	IsHigh() (bool, error)
}

// OutputPin drives a digital output pin.
type OutputPin interface {
	// SetHigh drives the pin high.
	SetHigh() error

	// SetLow drives the pin low.
	SetLow() error
}

// StatefulOutputPin reports the state requested for a digital output pin.
type StatefulOutputPin interface {
	OutputPin

	// IsSetHigh reports whether the output is configured high.
	IsSetHigh() (bool, error)
}

// TogglePin toggles a digital output pin.
type TogglePin interface {
	// Toggle switches the output between low and high.
	Toggle() error
}

// WaitPin waits for digital input transitions.
type WaitPin interface {
	// WaitForEdge blocks until the requested edge occurs or ctx is canceled.
	WaitForEdge(ctx context.Context, edge Edge) error
}

// IsLow reports whether pin is currently low.
func IsLow(pin InputPin) (bool, error) {
	high, err := pin.IsHigh()
	if err != nil {
		return false, err
	}
	return !high, nil
}

// SetState drives pin to state.
func SetState(pin OutputPin, state State) error {
	if state == High {
		return pin.SetHigh()
	}
	return pin.SetLow()
}

// IsSetLow reports whether pin is configured low.
func IsSetLow(pin StatefulOutputPin) (bool, error) {
	high, err := pin.IsSetHigh()
	if err != nil {
		return false, err
	}
	return !high, nil
}
