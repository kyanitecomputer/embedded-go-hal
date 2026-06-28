package usb

// Direction identifies a USB endpoint direction.
type Direction uint8

const (
	// OUT identifies host-to-device transfers.
	OUT Direction = iota

	// IN identifies device-to-host transfers.
	IN
)

// TransferType identifies a USB endpoint transfer type.
type TransferType uint8

const (
	// Control identifies control transfers.
	Control TransferType = iota

	// Isochronous identifies isochronous transfers.
	Isochronous

	// Bulk identifies bulk transfers.
	Bulk

	// Interrupt identifies interrupt transfers.
	Interrupt
)

// EndpointAddress identifies a USB endpoint number and direction.
type EndpointAddress struct {
	Number    uint8
	Direction Direction
}

// EndpointDescriptor describes a USB endpoint.
type EndpointDescriptor struct {
	Address       EndpointAddress
	TransferType  TransferType
	MaxPacketSize uint16
	Interval      uint8
}

// SetupPacket is a USB control transfer setup packet.
type SetupPacket struct {
	RequestType uint8
	Request     uint8
	Value       uint16
	Index       uint16
	Length      uint16
}

// Endpoint performs USB endpoint I/O.
type Endpoint interface {
	// ReadPacket reads one packet from the endpoint.
	ReadPacket(buf []byte) (int, error)

	// WritePacket writes one packet to the endpoint.
	WritePacket(buf []byte) (int, error)

	// Stall stalls or unstalls the endpoint.
	Stall(stalled bool) error
}

// DeviceController controls a USB device controller.
type DeviceController interface {
	// Enable starts the USB device controller.
	Enable() error

	// Disable stops the USB device controller.
	Disable() error

	// ConfigureEndpoint configures an endpoint.
	ConfigureEndpoint(desc EndpointDescriptor) (Endpoint, error)

	// SetAddress sets the USB device address assigned by the host.
	SetAddress(addr uint8) error
}

// SetupHandler handles USB control setup packets.
type SetupHandler interface {
	// HandleSetup handles one control setup packet.
	HandleSetup(setup SetupPacket) error
}

// HostController controls a USB host controller.
type HostController interface {
	// Enable starts the USB host controller.
	Enable() error

	// Disable stops the USB host controller.
	Disable() error

	// ResetPort resets the root port.
	ResetPort(port int) error
}
