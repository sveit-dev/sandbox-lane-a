package pager

import "testing"

func TestBounds(t *testing.T) {
	tests := []struct {
		name               string
		total, size, page  int
		wantStart, wantEnd int
	}{
		{"first page", 45, 20, 1, 0, 20},
		{"last partial page", 45, 20, 3, 40, 45},
		{"past the end", 45, 20, 4, 45, 45},
		{"zero size", 45, 0, 1, 0, 0},
		{"page zero", 45, 20, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, e := Bounds(tt.total, tt.size, tt.page)
			if s != tt.wantStart || e != tt.wantEnd {
				t.Errorf("Bounds(%d, %d, %d) = (%d, %d), want (%d, %d)",
					tt.total, tt.size, tt.page, s, e, tt.wantStart, tt.wantEnd)
			}
		})
	}
}

func TestPageSize(t *testing.T) {
	if got := PageSize(); got != 20 {
		t.Errorf("PageSize() = %d, want 20", got)
	}
}
