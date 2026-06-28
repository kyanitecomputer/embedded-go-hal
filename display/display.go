package display

// PixelFormat identifies framebuffer pixel encoding.
type PixelFormat uint32

const (
	// RGB565 stores pixels as 16-bit RGB values.
	RGB565 PixelFormat = iota + 1

	// XRGB8888 stores pixels as 32-bit RGB values with unused alpha bits.
	XRGB8888

	// ARGB8888 stores pixels as 32-bit alpha and RGB values.
	ARGB8888
)

// Mode describes framebuffer geometry.
type Mode struct {
	Width       int
	Height      int
	Stride      int
	PixelFormat PixelFormat
}

// Rect describes a rectangular display region.
type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

// Framebuffer provides CPU access to display memory.
type Framebuffer interface {
	// Mode returns framebuffer geometry.
	Mode() Mode

	// Bytes returns the CPU-visible framebuffer memory.
	Bytes() []byte
}

// Flusher writes framebuffer changes to the display or scanout memory.
type Flusher interface {
	// Flush writes all pending framebuffer changes.
	Flush() error
}

// RectFlusher writes framebuffer changes for a rectangular region.
type RectFlusher interface {
	// FlushRect writes framebuffer changes in r.
	FlushRect(r Rect) error
}

// Controller controls display scanout.
type Controller interface {
	// SetFramebuffer selects the framebuffer used for scanout.
	SetFramebuffer(fb Framebuffer) error

	// Enable starts display scanout.
	Enable() error

	// Disable stops display scanout.
	Disable() error
}
