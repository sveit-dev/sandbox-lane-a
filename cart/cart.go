// Package cart keeps item quantities within the limit set in protected.
package cart

import (
	"errors"

	"github.com/sveit-dev/sandbox-lane-a/protected"
)

// ErrLimit is returned when an addition would exceed protected.MaxItems.
var ErrLimit = errors.New("cart: item limit exceeded")

// Cart maps item names to quantities.
type Cart struct {
	items map[string]int
	count int
}

// New returns an empty cart.
func New() *Cart {
	return &Cart{items: map[string]int{}}
}

// Add adds qty of item. qty must be positive.
func (c *Cart) Add(item string, qty int) error {
	if qty <= 0 {
		return errors.New("cart: quantity must be positive")
	}
	if c.count+qty > protected.MaxItems {
		return ErrLimit
	}
	c.items[item] += qty
	c.count += qty
	return nil
}

// Count returns the total quantity in the cart.
func (c *Cart) Count() int {
	return c.count
}
