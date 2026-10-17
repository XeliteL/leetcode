package main

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "анаграмма",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "не анаграмма",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "разная длина",
			s:        "ab",
			t:        "a",
			expected: false,
		},
		{
			name:     "одинаковые строки",
			s:        "abc",
			t:        "abc",
			expected: true,
		},
		{
			name:     "одна буква — анаграмма",
			s:        "a",
			t:        "a",
			expected: true,
		},
		{
			name:     "одна буква — не анаграмма",
			s:        "a",
			t:        "b",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("isAnagram(%q, %q) = %v, ожидалось %v",
					tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
