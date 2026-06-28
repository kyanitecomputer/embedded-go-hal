package watchdog

import "time"

// Starter starts a watchdog timeout.
type Starter interface {
	// Start starts or reconfigures the watchdog timeout.
	Start(timeout time.Duration) error
}

// Feeder refreshes a watchdog timeout.
type Feeder interface {
	// Feed refreshes the watchdog timeout.
	Feed() error
}

// Stopper stops a watchdog timeout.
type Stopper interface {
	// Stop stops the watchdog timeout.
	Stop() error
}

// Device is a watchdog that can be started and fed.
type Device interface {
	Starter
	Feeder
}
