package main

// 994. Rotting Oranges
// https://leetcode.com/problems/rotting-oranges/

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "пример 1",
			grid:     [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expected: 4,
		},
		{
			name:     "недостижимый апельсин",
			grid:     [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			expected: -1,
		},
		{
			name:     "нет свежих апельсинов",
			grid:     [][]int{{0, 2}},
			expected: 0,
		},
		{
			name:     "все клетки пустые",
			grid:     [][]int{{0, 0}, {0, 0}},
			expected: 0,
		},
		{
			name:     "один свежий апельсин рядом с гнилым",
			grid:     [][]int{{2, 1}},
			expected: 1,
		},
		{
			name:     "один свежий апельсин без гнилых",
			grid:     [][]int{{1}},
			expected: -1,
		},
		{
			name:     "цепочка апельсинов",
			grid:     [][]int{{2, 1, 1, 1, 1}},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := orangesRotting(tt.grid)
			if result != tt.expected {
				t.Errorf("orangesRotting(%v) = %d, ожидалось %d",
					tt.grid, result, tt.expected)
			}
		})
	}
}
