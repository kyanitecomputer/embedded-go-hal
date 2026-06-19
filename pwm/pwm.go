package pwm

import "time"

// DutyCycle controls one PWM output by raw duty-cycle value.
type DutyCycle interface {
	// MaxDutyCycle returns the value representing 100% duty cycle.
	MaxDutyCycle() uint32

	// SetDutyCycle sets the duty cycle to duty divided by MaxDutyCycle.
	SetDutyCycle(duty uint32) error
}

// Channel controls one PWM output.
type Channel interface {
	// Set configures the PWM period and high-time duty duration.
	Set(period, duty time.Duration) error

	// Enable starts PWM output.
	Enable() error

	// Disable stops PWM output.
	Disable() error
}

// StatefulChannel reports PWM output state.
type StatefulChannel interface {
	Channel

	// Enabled reports whether PWM output is active.
	Enabled() (bool, error)
}

// SetFullyOff sets ch to 0% duty cycle.
func SetFullyOff(ch DutyCycle) error {
	return ch.SetDutyCycle(0)
}

// SetFullyOn sets ch to 100% duty cycle.
func SetFullyOn(ch DutyCycle) error {
	return ch.SetDutyCycle(ch.MaxDutyCycle())
}

// SetFraction sets ch to num divided by denom duty cycle.
func SetFraction(ch DutyCycle, num, denom uint32) error {
	if denom == 0 || num > denom {
		return ErrInvalidDutyCycle
	}
	duty := uint64(num) * uint64(ch.MaxDutyCycle()) / uint64(denom)
	return ch.SetDutyCycle(uint32(duty))
}

// SetPercent sets ch to percent divided by 100 duty cycle.
func SetPercent(ch DutyCycle, percent uint8) error {
	if percent > 100 {
		return ErrInvalidDutyCycle
	}
	return SetFraction(ch, uint32(percent), 100)
}
