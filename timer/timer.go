package timer

import (
	"context"
	"time"
)

// Delay blocks for a duration.
type Delay interface {
	// Delay blocks until d has elapsed.
	Delay(d time.Duration) error
}

// ContextDelay blocks for a duration or until a context is canceled.
type ContextDelay interface {
	// DelayContext blocks until d has elapsed or ctx is canceled.
	DelayContext(ctx context.Context, d time.Duration) error
}

// Counter reports monotonically increasing hardware time.
type Counter interface {
	// Now returns the current counter value.
	Now() uint64

	// Frequency returns counter ticks per second.
	Frequency() uint64
}

// Alarm schedules one-shot hardware wakeups.
type Alarm interface {
	// Set schedules an alarm at counter tick.
	Set(tick uint64) error

	// Stop disables the alarm.
	Stop() error
}
