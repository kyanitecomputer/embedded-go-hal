package netdev

// Link describes an Ethernet link mode.
type Link struct {
	SpeedMbps  int
	FullDuplex bool
}

// Ethernet provides L2 Ethernet frame access.
//
// Device-specific initialization such as PHY negotiation, firmware loading, or
// DMA ring setup must complete before an Ethernet value is used by a network
// stack.
type Ethernet interface {
	// HardwareAddr6 returns the device's 6-byte MAC address.
	HardwareAddr6() ([6]byte, error)

	// SendOffsetEthFrame transmits a complete Ethernet frame at the required offset.
	SendOffsetEthFrame(offsetTxEthFrame []byte) error

	// SetEthRecvHandler registers a callback for received Ethernet frames.
	SetEthRecvHandler(handler func(rxEthFrame []byte))

	// EthPoll services poll-driven devices and returns one received frame if present.
	EthPoll(buf []byte) (ethFrameOff, ethernetBytes int, err error)

	// MaxFrameSizeAndOffset returns the maximum device frame size and required frame offset.
	MaxFrameSizeAndOffset() (maxFrameSize int, frameOff int)
}
