package main

// 344. Reverse String
// https://leetcode.com/problems/reverse-string/

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		s        []byte
		expected []byte
	}{
		{
			name:     "нечётная длина",
			s:        []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:     "чётная длина",
			s:        []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name:     "один символ",
			s:        []byte{'a'},
			expected: []byte{'a'},
		},
		{
			name:     "два символа",
			s:        []byte{'a', 'b'},
			expected: []byte{'b', 'a'},
		},
		{
			name:     "все одинаковые символы",
			s:        []byte{'a', 'a', 'a'},
			expected: []byte{'a', 'a', 'a'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
			if !reflect.DeepEqual(tt.s, tt.expected) {
				t.Errorf("reverseString() = %v, ожидалось %v",
					tt.s, tt.expected)
			}
		})
	}
}
