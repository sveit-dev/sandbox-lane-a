// Package slug turns titles into URL-safe identifiers.
package slug

import (
	"strings"
	"unicode"
)

// Make lowercases s, keeps letters and digits, and joins words with single hyphens.
func Make(s string) string {
	var b strings.Builder
	pending := false
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if pending && b.Len() > 0 {
				b.WriteByte('-')
			}
			b.WriteRune(r)
			pending = false
			continue
		}
		pending = true
	}
	return b.String()
}
