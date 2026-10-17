package main

// 200. Number of Islands
// https://leetcode.com/problems/number-of-islands/

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "один остров",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "три острова",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "нет островов",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "вся сетка — остров",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			expected: 1,
		},
		{
			name: "каждая клетка — отдельный остров",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			expected: 5,
		},
		{
			name: "один элемент — суша",
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name: "один элемент — вода",
			grid: [][]byte{
				{'0'},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numIslands(tt.grid)
			if result != tt.expected {
				t.Errorf("numIslands(%v) = %d, ожидалось %d",
					tt.grid, result, tt.expected)
			}
		})
	}
}
