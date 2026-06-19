package pwm

import "errors"

// ErrInvalidDutyCycle reports an invalid duty-cycle fraction.
var ErrInvalidDutyCycle = errors.New("pwm: invalid duty cycle")
