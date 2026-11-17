package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard ASCII string",
			input:    "hello world!",
			expected: "!dlrow olleh",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "even length string",
			input:    "123456",
			expected: "654321",
		},
		{
			name:     "string with spaces and punctuation",
			input:    "Go is awesome!",
			expected: "!emosewa si oG",
		},
		{
			name:     "multi-byte Unicode characters",
			input:    "hello, 世界",
			expected: "界世 ,olleh",
		},
		{
			name:     "string with emojis",
			input:    "Go 🚀",
			expected: "🚀 oG",
		},
	}

	for _, tt := range tests {
		t.Run("reverseString_one/"+tt.name, func(t *testing.T) {
			got := reverseString_one(tt.input)
			if got != tt.expected {
				t.Errorf("reverseString_one(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})

		t.Run("reverseString_two/"+tt.name, func(t *testing.T) {
			got := reverseString_two(tt.input)
			if got != tt.expected {
				t.Errorf("reverseString_two(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkReverseStringOne(b *testing.B) {
	input := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseString_one(input)
	}
}

func BenchmarkReverseStringTwo(b *testing.B) {
	input := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseString_two(input)
	}
}
