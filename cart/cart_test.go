package cart

import (
	"errors"
	"testing"

	"github.com/sveit-dev/sandbox-lane-a/protected"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name      string
		adds      []int
		wantCount int
		wantErr   error
	}{
		{"single", []int{3}, 3, nil},
		{"several", []int{1, 2, 3}, 6, nil},
		{"exactly the limit", []int{protected.MaxItems}, protected.MaxItems, nil},
		{"over the limit", []int{protected.MaxItems, 1}, protected.MaxItems, ErrLimit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			var err error
			for _, q := range tt.adds {
				err = c.Add("item", q)
			}
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if c.Count() != tt.wantCount {
				t.Fatalf("Count() = %d, want %d", c.Count(), tt.wantCount)
			}
		})
	}
}

func TestAddRejectsNonPositive(t *testing.T) {
	for _, q := range []int{0, -1} {
		if err := New().Add("item", q); err == nil {
			t.Errorf("Add(%d) returned nil error", q)
		}
	}
}
