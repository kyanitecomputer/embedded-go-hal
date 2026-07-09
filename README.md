# embedded-go-hal

Small, dependency-free Go interfaces for embedded and bare-metal hardware
drivers. Part of the [Kyanite](https://github.com/kyanitecomputer) stack.

> **Status:** experimental — expect breaking changes.

## Overview

`src.kyanite.computer/hal` defines the hardware-abstraction interfaces that
Kyanite drivers implement and firmware consumes. The packages use only the Go
standard library — no external dependencies — so they can be shared across
TamaGo bare-metal targets and host tooling alike.

## Import

```sh
go get src.kyanite.computer/hal
```

## Packages

| Package | Purpose |
| ------- | ------- |
| `adc` | Analog-to-digital converter channels and reference voltage reporting |
| `block` | Logical block devices compatible with `io.ReaderAt` / `io.WriterAt` |
| `can` | CAN 2.0 frame identifiers, validation, receive, and transmit |
| `clock` | Clock sources, configurable sources, and clock gates |
| `dac` | Digital-to-analog converter channels and reference voltage reporting |
| `digital` | GPIO-style input, output, toggle, state, and edge-wait |
| `display` | Framebuffer, display mode, flush, rectangular flush, scanout control |
| `dma` | DMA buffers, cache maintenance, and DMA allocator |
| `flash` | Raw NOR/NAND geometry and NOR access via `io.ReaderAt` / `io.WriterAt` |
| `i2c` | Addressed I2C transaction interface |
| `interrupt` | IRQ identifiers, handler registration, control, and waiting |
| `mdio` | Ethernet PHY management bus |
| `netdev` | L2 Ethernet frame device |
| `pwm` | PWM channel control |
| `reg` | MMIO register access primitives |
| `reset` | Reset lines and reset sequences |
| `rng` | Random source compatible with `io.Reader` |
| `serial` | Byte-oriented serial port compatible with `io` |
| `spi` | SPI bus, device, and transaction |
| `timer` | Delays, context-aware delays, counters, and alarms |
| `usb` | USB device/host controller, endpoint, and setup packet |
| `watchdog` | Watchdog start/feed/stop |

## Roadmap

Interface packages under consideration: PCIe, SD/MMC/eMMC, I3C, 1-Wire, RTC, fan
controllers/tachometers, mailbox/IPI, crypto accelerators, TRNG health-test,
pinctrl/pinmux, power domains/regulators, DMA engine channels, IOMMU/MPU, and
common sensor conventions.

## Contributing

See the org-wide [CONTRIBUTING guide](https://github.com/kyanitecomputer/.github/blob/main/CONTRIBUTING.md).
Contributions are dual-licensed.

## Security

See the org-wide [SECURITY policy](https://github.com/kyanitecomputer/.github/blob/main/SECURITY.md).

## License

Dual-licensed under either of Apache-2.0 ([LICENSE-APACHE](LICENSE-APACHE)) or
MIT ([LICENSE-MIT](LICENSE-MIT)) at your option.
