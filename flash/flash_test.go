package flash

import (
	"errors"
	"testing"
)

func TestGeometryValidate(t *testing.T) {
	tests := []struct {
		name string
		g    Geometry
		want error
	}{
		{
			name: "valid",
			g: Geometry{
				TotalSize:      1024,
				EraseBlockSize: 256,
				PageSize:       16,
			},
		},
		{
			name: "zero total size",
			g: Geometry{
				EraseBlockSize: 256,
				PageSize:       16,
			},
			want: ErrInvalidGeometry,
		},
		{
			name: "erase block larger than device",
			g: Geometry{
				TotalSize:      128,
				EraseBlockSize: 256,
				PageSize:       16,
			},
			want: ErrInvalidGeometry,
		},
		{
			name: "page does not divide erase block",
			g: Geometry{
				TotalSize:      1024,
				EraseBlockSize: 256,
				PageSize:       24,
			},
			want: ErrInvalidGeometry,
		},
		{
			name: "erase block does not divide device",
			g: Geometry{
				TotalSize:      1000,
				EraseBlockSize: 256,
				PageSize:       16,
			},
			want: ErrInvalidGeometry,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.g.Validate()
			if !errors.Is(err, tt.want) {
				t.Fatalf("Validate() = %v, want %v", err, tt.want)
			}
		})
	}
}
