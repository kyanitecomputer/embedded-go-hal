package adc

// Sample is one analog-to-digital conversion result.
type Sample struct {
	// Value is the raw conversion value.
	Value uint32

	// Bits is the number of valid low-order bits in Value.
	Bits uint8
}

// Channel reads one ADC input channel.
type Channel interface {
	// Read reads one conversion sample.
	Read() (Sample, error)
}

// Reference reports ADC reference information.
type Reference interface {
	// ReferenceMillivolts returns the ADC reference voltage in millivolts.
	ReferenceMillivolts() uint32
}
