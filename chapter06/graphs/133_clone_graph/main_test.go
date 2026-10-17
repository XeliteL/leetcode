package main

// 133. Clone Graph
// https://leetcode.com/problems/clone-graph/

import "testing"

// buildGraph строит граф из списка смежности (1-индексированного)
func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList)+1)
	for i := 1; i <= len(adjList); i++ {
		nodes[i] = &Node{Val: i}
	}
	for i, neighbors := range adjList {
		for _, n := range neighbors {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[n])
		}
	}
	return nodes[1]
}

// graphToAdjList преобразует граф обратно в список смежности
func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	visited := map[int]*Node{}
	queue := []*Node{node}
	visited[node.Val] = node

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, n := range curr.Neighbors {
			if _, ok := visited[n.Val]; !ok {
				visited[n.Val] = n
				queue = append(queue, n)
			}
		}
	}

	result := make([][]int, len(visited))
	for _, n := range visited {
		neighbors := []int{}
		for _, nb := range n.Neighbors {
			neighbors = append(neighbors, nb.Val)
		}
		result[n.Val-1] = neighbors
	}
	return result
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "четыре узла",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "один узел без соседей",
			adjList: [][]int{{}},
		},
		{
			name:    "пустой граф",
			adjList: [][]int{},
		},
		{
			name:    "два узла связанных между собой",
			adjList: [][]int{{2}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			cloned := cloneGraph(original)

			// проверяем что клон не nil если оригинал не nil
			if original == nil && cloned != nil {
				t.Errorf("ожидался nil, получили не nil")
				return
			}
			if original != nil && cloned == nil {
				t.Errorf("ожидался не nil, получили nil")
				return
			}
			if original == nil {
				return
			}

			// проверяем что клон — отдельный граф (не те же указатели)
			if original == cloned {
				t.Errorf("клон указывает на тот же узел что и оригинал")
			}

			// проверяем структуру через список смежности
			origAdj := graphToAdjList(original)
			cloneAdj := graphToAdjList(cloned)

			if len(origAdj) != len(cloneAdj) {
				t.Errorf("разное количество узлов: %d vs %d", len(origAdj), len(cloneAdj))
			}
		})
	}
}
