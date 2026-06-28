package interrupt

import "context"

// IRQ identifies one interrupt line.
type IRQ uint32

// Handler handles an interrupt.
type Handler func()

// Registrar installs interrupt handlers.
type Registrar interface {
	// Register associates handler with irq.
	Register(irq IRQ, handler Handler) error
}

// Controller controls interrupt delivery.
type Controller interface {
	// Enable enables delivery of irq.
	Enable(irq IRQ) error

	// Disable disables delivery of irq.
	Disable(irq IRQ) error

	// Acknowledge acknowledges irq after it has been handled.
	Acknowledge(irq IRQ) error
}

// Waiter waits for interrupt delivery.
type Waiter interface {
	// Wait blocks until irq is delivered or ctx is canceled.
	Wait(ctx context.Context, irq IRQ) error
}
