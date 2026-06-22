package dma

// Buffer describes memory that a device can access through DMA.
type Buffer struct {
	// Phys is the device-visible physical address of Bytes.
	Phys uint64

	// Bytes is the CPU-visible memory backing the DMA buffer.
	Bytes []byte
}

// Cache maintains CPU cache coherency for DMA buffers.
type Cache interface {
	// Clean writes dirty cache lines back to memory before device reads.
	Clean(start, size uintptr)

	// Invalidate discards cache lines before CPU reads data written by a device.
	Invalidate(start, size uintptr)
}

// Allocator reserves and releases DMA-capable buffers.
type Allocator interface {
	// Reserve reserves a DMA buffer of size bytes with the requested alignment.
	Reserve(size, align int) (Buffer, error)

	// Release releases a DMA buffer returned by Reserve.
	Release(Buffer) error
}
