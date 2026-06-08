package i2c

// Operation describes one I2C transaction segment.
//
// Exactly one of Read or Write should be non-empty. Adjacent operations of the
// same direction may be combined by an implementation. Direction changes within
// one transaction use a repeated start and the transaction ends with a stop.
type Operation struct {
	Read  []byte
	Write []byte
}

// Bus performs I2C transactions with addressed devices.
type Bus interface {
	// Transaction executes operations against addr.
	Transaction(addr uint16, ops []Operation) error
}

// Read reads len(buf) bytes from addr.
func Read(bus Bus, addr uint16, buf []byte) error {
	return bus.Transaction(addr, []Operation{{Read: buf}})
}

// Write writes buf to addr.
func Write(bus Bus, addr uint16, buf []byte) error {
	return bus.Transaction(addr, []Operation{{Write: buf}})
}

// WriteRead writes w to addr, then reads into r with a repeated start.
func WriteRead(bus Bus, addr uint16, w, r []byte) error {
	return bus.Transaction(addr, []Operation{{Write: w}, {Read: r}})
}
