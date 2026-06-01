package reset

// Line controls one reset signal.
type Line interface {
	// Assert holds the reset signal active.
	Assert() error

	// Deassert releases the reset signal.
	Deassert() error
}

// Resetter performs a complete reset pulse or sequence.
type Resetter interface {
	// Reset resets the target hardware.
	Reset() error
}

// StatefulLine reports reset signal state.
type StatefulLine interface {
	Line

	// Asserted reports whether the reset signal is active.
	Asserted() (bool, error)
}
