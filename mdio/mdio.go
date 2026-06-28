package mdio

// Bus provides IEEE 802.3 Clause 22 and Clause 45 MDIO register access.
//
// Implementations use devAddr zero for Clause 22 transactions and non-zero
// devAddr values for Clause 45 transactions.
type Bus interface {
	// Read reads a 16-bit PHY register.
	Read(phyAddr, devAddr uint8, regAddr uint16) (uint16, error)

	// Write writes a 16-bit PHY register.
	Write(phyAddr, devAddr uint8, regAddr, value uint16) error
}
