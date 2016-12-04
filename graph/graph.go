package graph

import (
	"fmt"
	"sync"
)

// Interface defines the behavior of a Graph data structure
type Interface interface {
	Size() int
	Insert(Node)
	Remove(Node)
	HasNode(Node) bool
	DFS(Node, SearchFunc) (interface{}, error)
	BFS(Node, SearchFunc) (interface{}, error)
}

// Node defines behavior of a Graph Node, with a value and a set of neighbors.
type Node interface {
	Value() interface{}
	Neighbors() map[Node]struct{}
	HasNeighbor(Node) bool
	AddNeighbor(Node)
	RemoveNeighbor(Node)
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
	nodes map[Node]struct{}
}

// IntNode implements Node for int values
type IntNode struct {
	lock      sync.Mutex
	value     int
	neighbors map[Node]struct{}
}

// MissingNodeError describes the case when a Graph does not contain a Node that
// has been referenced.
type MissingNodeError struct {
	graph Interface
	node  Node
}

func (err *MissingNodeError) Error() string {
	return fmt.Sprintf("Graph does not contain Node\ngraph: %v\nnode: %v", err.graph, err.node)
}

// NotFoundError describes the state when a Graph search completes without
// completing the objective, which is indicated by the SearchFunc returning
// done=true.
type NotFoundError struct {
	msg string
}

func (err NotFoundError) Error() string {
	return err.msg
}

// NewIntNode creates and returns a new *IntNode
func NewIntNode(value int) *IntNode {
	return &IntNode{
		value:     value,
		neighbors: map[Node]struct{}{},
	}
}

// String returns a string representation of an IntNode
func (n *IntNode) String() string {
	return fmt.Sprintf("%v[%p]", n.Value(), n)
}

// Value returns the value of an IntNode
func (n *IntNode) Value() interface{} {
	return n.value
}

// Neighbors returns the set of n's neighboring IntNodes
func (n *IntNode) Neighbors() map[Node]struct{} {
	return n.neighbors
}

// AddNeighbor adds an edge from n to node
func (n *IntNode) AddNeighbor(node Node) {
	if n.HasNeighbor(node) {
		return
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	n.neighbors[node] = struct{}{}
}

// RemoveNeighbor removes an edge from n to node, if it exists
func (n *IntNode) RemoveNeighbor(node Node) {
	if !n.HasNeighbor(node) {
		return
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	delete(n.neighbors, node)
}

// HasNeighbor returns true if node is n's neighbor
func (n *IntNode) HasNeighbor(node Node) bool {
	n.lock.Lock()
	defer n.lock.Unlock()
	_, ok := n.neighbors[node]
	return ok
}

// NewIntGraph creates and returns a new *IntGraph
func NewIntGraph() *IntGraph {
	return &IntGraph{nodes: map[Node]struct{}{}}
}

// String returns a string representation of the graph as an adjecency list.
func (g *IntGraph) String() string {
	var str string
	for node := range g.nodes {
		str += fmt.Sprintf("%v", node) + "->{ "
		for n := range node.Neighbors() {
			str += fmt.Sprintf("%v", n) + " "
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
func (g *IntGraph) HasNode(node Node) bool {
	g.lock.Lock()
	defer g.lock.Unlock()
	_, ok := g.nodes[node]
	return ok
}

// Insert adds node to the graph
func (g *IntGraph) Insert(node Node) {
	if g.HasNode(node) {
		return
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes[node] = struct{}{}
}

// Remove removes node from the graph
func (g *IntGraph) Remove(node Node) {
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
func (g *IntGraph) DFS(node Node, sf SearchFunc) (interface{}, error) {
	if !g.HasNode(node) {
		return nil, &MissingNodeError{g, node}
	}
	visited := map[Node]struct{}{}
	return g.dfs(node, sf, visited)
}

func (g *IntGraph) dfs(node Node, sf SearchFunc, visited map[Node]struct{}) (interface{}, error) {
	if value, done := sf(node); done {
		return value, nil
	}
	visited[node] = struct{}{}
	for nbr := range node.Neighbors() {
		if _, ok := visited[nbr]; !ok {
			return g.dfs(nbr, sf, visited)
		}
	}
	return nil, NotFoundError{"Search exhausted graph: objective not found"}
}

// BFS executes a breadth-first search, applying the SearchFunc to each IntNode
// visited to yield values and determine whether or not to continue.
func (g *IntGraph) BFS(node Node, sf SearchFunc) (interface{}, error) {
	if !g.HasNode(node) {
		return nil, &MissingNodeError{g, node}
	}
	return nil, NotFoundError{"Search exhausted graph: objective not found"}
}
