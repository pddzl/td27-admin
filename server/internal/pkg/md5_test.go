package pkg

import (
	"testing"
)

func TestMD5V(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		expect string
	}{
		{
			name:   "empty string",
			input:  []byte(""),
			expect: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:   "hello world",
			input:  []byte("hello world"),
			expect: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
		{
			name:   "special chars",
			input:  []byte("!@#$%^&*()"),
			expect: "05b28d17a7b6e7024b6e5d8cc43a8bf7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MD5V(tt.input)
			if result != tt.expect {
				t.Errorf("MD5V(%q) = %q, want %q", tt.input, result, tt.expect)
			}
		})
	}
}
