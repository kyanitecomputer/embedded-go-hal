package dac

// Sample is one digital-to-analog output value.
type Sample struct {
	// Value is the raw output value.
	Value uint32

	// Bits is the number of valid low-order bits in Value.
	Bits uint8
}

// Channel writes one DAC output channel.
type Channel interface {
	// Write updates the analog output value.
	Write(sample Sample) error
}

// Reference reports DAC reference information.
type Reference interface {
	// ReferenceMillivolts returns the DAC reference voltage in millivolts.
	ReferenceMillivolts() uint32
}
