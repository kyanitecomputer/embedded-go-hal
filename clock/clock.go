package clock

// Source reports a clock frequency.
type Source interface {
	// Frequency returns the clock rate in hertz.
	Frequency() uint64
}

// ConfigurableSource changes a clock frequency.
type ConfigurableSource interface {
	Source

	// SetFrequency changes the clock rate in hertz.
	SetFrequency(hz uint64) error
}

// Gate controls whether a clock is supplied to a peripheral.
type Gate interface {
	// Enable supplies the clock.
	Enable() error

	// Disable stops supplying the clock.
	Disable() error
}

// StatefulGate reports whether a clock gate is enabled.
type StatefulGate interface {
	Gate

	// Enabled reports whether the clock is supplied.
	Enabled() (bool, error)
}
