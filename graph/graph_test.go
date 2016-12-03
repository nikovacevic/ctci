package graph

import "testing"

var nodeNeighborTests = []struct {
	nodes  []int
	add    [][]int
	remove [][]int
	exp    [][]int
}{
	{
		[]int{0, 1, 2},
		[][]int{
			[]int{1, 2},
			[]int{2},
			[]int{1},
		},
		[][]int{
			[]int{2},
			[]int{},
			[]int{1},
		},
		[][]int{
			[]int{1},
			[]int{2},
			[]int{},
		},
	},
	{
		[]int{0, 1, 2, 3, 4, 5, 6},
		[][]int{
			[]int{1, 2, 3, 5},
			[]int{4},
			[]int{1, 5},
			[]int{2, 4, 5},
			[]int{3},
			[]int{1, 4, 6},
			[]int{0},
		},
		[][]int{
			[]int{2},
			[]int{},
			[]int{4},
			[]int{5},
			[]int{},
			[]int{},
			[]int{0},
		},
		[][]int{
			[]int{1, 3, 5},
			[]int{4},
			[]int{1, 5},
			[]int{2, 4},
			[]int{3},
			[]int{1, 4, 6},
			[]int{},
		},
	},
}

func TestNodeNeighor(t *testing.T) {
	for _, tt := range nodeNeighborTests {
		// Create nodes
		nodes := make([]*IntNode, len(tt.nodes))
		for i, n := range tt.nodes {
			nodes[i] = NewIntNode(n)
		}
		g := NewIntGraph()
		for _, n := range nodes {
			g.Insert(n)
		}
		// Add neighbors
		for i, add := range tt.add {
			for _, a := range add {
				nodes[i].AddNeighbor(nodes[a])
			}
		}
		for i, add := range tt.add {
			if len(nodes[i].Neighbors()) != len(add) {
				t.Errorf("Node %s neighbors: expected %v, actual %v", nodes[i], len(add), len(nodes[i].Neighbors()))
			}
			for _, a := range add {
				if !nodes[i].HasNeighbor(nodes[a]) {
					t.Errorf("Node neighbors: expected %s->%s", nodes[i], nodes[a])
				}
			}
		}
		// Remove neighbors
		for i, remove := range tt.remove {
			for _, r := range remove {
				nodes[i].RemoveNeighbor(nodes[r])
			}
		}
		for i, exp := range tt.exp {
			if len(nodes[i].Neighbors()) != len(exp) {
				t.Errorf("Node %s neighbors: expected %v, actual %v", nodes[i], len(exp), len(nodes[i].Neighbors()))
			}
			for _, e := range exp {
				if !nodes[i].HasNeighbor(nodes[e]) {
					t.Errorf("Node neighbors: expected %s->%s", nodes[i], nodes[e])
				}
			}
		}
	}
}
