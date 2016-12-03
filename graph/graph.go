package graph

import (
	"fmt"
	"sync"
)

// Interface defines the behavior of a Graph data structure
type Interface interface {
	Size() int
	Insert(*Node)
	Remove(*Node)
	HasNode(*Node) bool
	DFS(*Node, SearchFunc) (interface{}, error)
	BFS(*Node, SearchFunc) (interface{}, error)
}

// Node defines behavior of a Graph Node, with a value and a set of neighbors.
type Node interface {
	Value() interface{}
	Neighbors() map[*Node]struct{}
	HasNeighbor(*Node) bool
	AddNeighbor(*Node)
	RemoveNeighbor(*Node)
}

// Tree defines the behavior of a Tree data structure, which is a subset of
// the graph.Interface in that it has a defined root and is acyclic
type Tree interface {
	Interface
	Root() Node
}

// SearchFunc is applied to each Node that a graph search algorithm visits.
// It has the option to return a value, as well as to tell the search routine
// to stop by returning done=true.
type SearchFunc func(Node) (value interface{}, done bool)

// IntGraph implements a Graph of IntNodes
type IntGraph struct {
	lock  sync.Mutex
	nodes map[*IntNode]struct{}
}

// IntNode implements Node for int values
type IntNode struct {
	lock      sync.Mutex
	value     int
	neighbors map[*IntNode]struct{}
}

// NewIntNode creates and returns a new *IntNode
func NewIntNode(value int) *IntNode {
	return &IntNode{
		value:     value,
		neighbors: map[*IntNode]struct{}{},
	}
}

// String returns a string representation of an IntNode
func (n *IntNode) String() string {
	return fmt.Sprintf("%v[%p]", n.Value(), n)
}

// Value returns the value of an IntNode
func (n *IntNode) Value() int {
	return n.value
}

// Neighbors returns the set of n's neighboring IntNodes
func (n *IntNode) Neighbors() map[*IntNode]struct{} {
	return n.neighbors
}

// AddNeighbor adds an edge from n to node
func (n *IntNode) AddNeighbor(node *IntNode) {
	if n.HasNeighbor(node) {
		return
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	n.neighbors[node] = struct{}{}
}

// RemoveNeighbor removes an edge from n to node, if it exists
func (n *IntNode) RemoveNeighbor(node *IntNode) {
	if !n.HasNeighbor(node) {
		return
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	delete(n.neighbors, node)
}

// HasNeighbor returns true if node is n's neighbor
func (n *IntNode) HasNeighbor(node *IntNode) bool {
	n.lock.Lock()
	defer n.lock.Unlock()
	_, ok := n.neighbors[node]
	return ok
}

// NewIntGraph creates and returns a new *IntGraph
func NewIntGraph() *IntGraph {
	return &IntGraph{nodes: map[*IntNode]struct{}{}}
}

// String returns a string representation of the graph as an adjecency list.
func (g *IntGraph) String() string {
	var str string
	for node := range g.nodes {
		str += node.String() + "->{ "
		for n := range node.Neighbors() {
			str += n.String() + " "
		}
		str += "}\n"
	}
	return str
}

// Size returns the number of IntNodes in the IntGraph
func (g *IntGraph) Size() int {
	return len(g.nodes)
}

// HasNode returns the true if the IntGraph has the node
func (g *IntGraph) HasNode(node *IntNode) bool {
	g.lock.Lock()
	defer g.lock.Unlock()
	_, ok := g.nodes[node]
	return ok
}

// Insert adds node to the graph
func (g *IntGraph) Insert(node *IntNode) {
	if g.HasNode(node) {
		return
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes[node] = struct{}{}
}

// Remove removes node from the graph
func (g *IntGraph) Remove(node *IntNode) {
	if !g.HasNode(node) {
		return
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	delete(g.nodes, node)
	for n := range g.nodes {
		n.RemoveNeighbor(node)
	}
	return
}

// DFS executes a depth-first search, applying the SearchFunc to each IntNode
// visited to yield values and determine whether or not to continue.
func (g *IntGraph) DFS(node *IntNode, sf SearchFunc) (interface{}, error) {
	// TODO
	return nil, nil
}

// BFS executes a breadth-first search, applying the SearchFunc to each IntNode
// visited to yield values and determine whether or not to continue.
func (g *IntGraph) BFS(*Node, SearchFunc) (interface{}, error) {
	// TODO
	return nil, nil
}
