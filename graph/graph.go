package graph

import (
	"fmt"
	"sort"
	"sync"
)

// Interface defines the behavior of a Graph data structure
type Interface interface {
	Size() int
	Insert(Node)
	Remove(Node)
}

// Node defines behavior of a Graph Node, with a value and a set of neighbors.
type Node interface {
	Value() interface{}
	Neighbors() map[Node]struct{}
	HasNeighbor(Node) bool
	AddNeighbor(Node) error
	RemoveNeighbor(Node) error
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

// BST defines the behavior of a binary search tree data structure
type BST interface {
	Size() int
	Root() BSTNode
	Insert(BSTNode)
	Remove(BSTNode)
	Search() (interface{}, error)
}

// BSTNode defines behavior of a node in a BST
type BSTNode interface {
	Value() interface{}
	LessThan(BSTNode) bool
	Left() BSTNode
	Right() BSTNode
	setLeft(BSTNode)
	setRight(BSTNode)
}

// LessFunc is used to compare two Nodes, e.g. in Binary Tree Search, returning
// true if Node a is less than Node b
type LessFunc func(a, b BSTNode) bool

// IntBST implements a Binary Search Tree of integers
type IntBST struct {
	lock sync.Mutex
	root BSTNode
	size int
	lf   LessFunc
}

// IntBSTNode implements a Node for Binary Search Tree of integers
type IntBSTNode struct {
	lock  sync.Mutex
	value int
	left  BSTNode
	right BSTNode
}

// MissingNodeError describes the case when a Graph does not contain a Node that
// has been referenced.
type MissingNodeError struct {
	graph Interface
	node  Node
}

func (err MissingNodeError) Error() string {
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
func (n *IntNode) AddNeighbor(node Node) error {
	if n.HasNeighbor(node) {
		return nil
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	n.neighbors[node] = struct{}{}
	return nil
}

// RemoveNeighbor removes an edge from n to node, if it exists
func (n *IntNode) RemoveNeighbor(node Node) error {
	if !n.HasNeighbor(node) {
		return nil
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	delete(n.neighbors, node)
	return nil
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
	/* TODO Determine if losing this to the Interface change is problematic
	for n := range g.nodes {
		n.RemoveNeighbor(node)
	}
	*/
	return
}

// DFS executes a depth-first search, applying the SearchFunc to each IntNode
// visited to yield values and determine whether or not to continue.
func (g *IntGraph) DFS(node Node, sf SearchFunc) (interface{}, error) {
	if !g.HasNode(node) {
		return nil, MissingNodeError{g, node}
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
		return nil, MissingNodeError{g, node}
	}
	visited := map[Node]struct{}{}
	queue := []Node{node}
	done := false
	for len(queue) > 0 || done {
		curr := queue[0]
		if value, done := sf(curr); done {
			return value, nil
		}
		visited[curr] = struct{}{}
		queue = queue[1:]
		for n := range curr.Neighbors() {
			if _, ok := visited[n]; !ok {
				queue = append(queue, n)
			}
		}
	}
	return nil, NotFoundError{"Search exhausted graph: objective not found"}
}

// RouteExists (4.1) returns true if a route from start to finish exists
func (g *IntGraph) RouteExists(start Node, finish Node) bool {
	if !g.HasNode(start) || !g.HasNode(finish) {
		return false
	}
	if _, err := g.BFS(start, func(node Node) (interface{}, bool) {
		if node == finish {
			return nil, true
		}
		return nil, false
	}); err == nil {
		return true
	}
	return false
}

// NewIntBSTNode returns a new *IntBSTNode
func NewIntBSTNode(value int) *IntBSTNode {
	return &IntBSTNode{value: value}
}

// Value returns the int value of the node
func (n *IntBSTNode) Value() interface{} {
	return n.value
}

// Left returns the left child of n
func (n *IntBSTNode) Left() BSTNode {
	return n.left
}
func (n *IntBSTNode) setLeft(l BSTNode) {
	n.left = l
}

// Right returns the right child of n
func (n *IntBSTNode) Right() BSTNode {
	return n.right
}
func (n *IntBSTNode) setRight(r BSTNode) {
	n.right = r
}

// LessThan returns true if n's value is less than node's value. NOTE that this
// function relies on a type assertion that will panic if node's value is not
// assertable to (int). Can we do better?
func (n *IntBSTNode) LessThan(node BSTNode) bool {
	return n.value < node.Value().(int)
}

// NewIntBST (4.2) takes a slice of integers, returning a minimal binary search
// tree of those numbers.
func NewIntBST(nums []int) *IntBST {
	sort.Ints(nums)
	var r BSTNode
	var m int
	if len(nums) > 0 {
		m = len(nums) / 2
		r = NewIntBSTNode(nums[m])
	}
	t := &IntBST{
		root: r,
		lf: func(a, b BSTNode) bool {
			return a.Value().(int) < b.Value().(int)
		},
		size: len(nums),
	}
	if len(nums[0:m]) > 0 {
		r.setLeft(NewIntBST(nums[0:m]).Root())
	}
	if len(nums[m+1:]) > 0 {
		r.setRight(NewIntBST(nums[m+1:]).Root())
	}
	return t
}

// Root returns the root Node of the IntBST
func (bst *IntBST) Root() BSTNode {
	return bst.root
}

// Size returns number of Nodes in the BST
func (bst *IntBST) Size() int {
	return bst.size
}

// Insert adds node to the BST
func (bst *IntBST) Insert(node BSTNode) {
	if bst.root == nil {
		bst.root = node
		bst.size = 1
		return
	}
	curr := bst.root
	for curr != nil {
		if curr.LessThan(node) {
			curr = curr.Right()
		} else {
			curr = curr.Left()
		}
	}
	curr = node
}

// Remove removes node from the BST
func (bst *IntBST) Remove(node BSTNode) {
	// TODO
	return
}

// Search uses binary tree search to
func (bst *IntBST) Search() (interface{}, error) {
	// TODO
	return nil, nil
}

// PreOrderTraverse applies function f from smallest to largest node in BST
func (bst *IntBST) PreOrderTraverse(f func(BSTNode)) {
	curr := bst.root
	preOrderTraverse(curr, f)
}

func preOrderTraverse(node BSTNode, f func(BSTNode)) {
	if node == nil {
		return
	}
	preOrderTraverse(node.Left(), f)
	f(node)
	preOrderTraverse(node.Right(), f)
}

// ToSlice converts an IntBST to a slice of ints
func (bst *IntBST) ToSlice() []int {
	s := []int{}
	bst.PreOrderTraverse(func(node BSTNode) {
		s = append(s, node.Value().(int))
	})
	return s
}
