// Package pager splits result sets into pages.
package pager

import "github.com/sveit-dev/sandbox-lane-a/protected"

// PageSize returns the configured page size, or 10 if the config cannot be read.
func PageSize() int {
	c, err := protected.Load()
	if err != nil || c.DefaultPageSize <= 0 {
		return 10
	}
	return c.DefaultPageSize
}

// Bounds returns the start and end indexes of page (1-based) for total items.
func Bounds(total, size, page int) (start, end int) {
	if size <= 0 || page < 1 {
		return 0, 0
	}
	start = (page - 1) * size
	if start > total {
		start = total
	}
	end = start + size
	if end > total {
		end = total
	}
	return start, end
}

// Count returns how many pages are needed to show total items, size per page.
func Count(total, size int) int {
	if size <= 0 {
		return 0
	}
	return total / size
}
