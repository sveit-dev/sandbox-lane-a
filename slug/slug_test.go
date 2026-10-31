package slug

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"empty", "", ""},
		{"simple", "Hello World", "hello-world"},
		{"punctuation", "  Go, fast & safe!  ", "go-fast-safe"},
		{"digits", "Release 2 notes", "release-2-notes"},
		{"repeated separators", "a---b__c", "a-b-c"},
		{"trimmed punctuation and spaces", "  Hello,   World!  ", "hello-world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make(tt.in); got != tt.want {
				t.Errorf("Make(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
