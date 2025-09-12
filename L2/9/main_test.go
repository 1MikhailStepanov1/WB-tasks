package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
			wantErr:  false,
		},
		{
			name:     "single letter",
			input:    "a",
			expected: "a",
			wantErr:  false,
		},
		{
			name:     "multiple letters",
			input:    "abc",
			expected: "abc",
			wantErr:  false,
		},

		{
			name:     "simple unpack",
			input:    "a3",
			expected: "aaa",
			wantErr:  false,
		},
		{
			name:     "multiple digits",
			input:    "a2b3c4",
			expected: "aabbbcccc",
			wantErr:  false,
		},
		{
			name:     "digit 1",
			input:    "a1",
			expected: "a",
			wantErr:  false,
		},
		{
			name:     "digit 0 should not appear",
			input:    "a0",
			expected: "",
			wantErr:  true,
		},

		{
			name:     "escaped backslash",
			input:    "\\\\",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "mixed escaped and normal",
			input:    "a\\3b2",
			expected: "a3bb",
			wantErr:  false,
		},
		{
			name:     "digit at start",
			input:    "3a",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "multiple digits at start",
			input:    "42abc",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "only digits",
			input:    "123",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "lone backslash at end",
			input:    "abc\\",
			expected: "",
			wantErr:  true,
		},

		{
			name:     "unicode characters",
			input:    "а2б3",
			expected: "ааббб",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpack(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.expected {
				t.Errorf("unpack() = %v, want %v", got, tt.expected)
			}
		})
	}
}
