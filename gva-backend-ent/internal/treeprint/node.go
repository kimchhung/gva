package treeprint

import (
	"reflect"
	"strings"
)

// Node represents a tree node.
type Node struct {
	Value   interface{}
	Nodes   []*Node
	Virtual bool
}

// New creates a new tree / root node.
func New(value interface{}) *Node {
	return &Node{
		Value: value,
	}
}

// Add adds a leaf node
func (n *Node) Add(value interface{}) *Node {
	node := New(value)
	n.Nodes = append(n.Nodes, node)
	return node
}

// AddPath adds a chain of nodes
func (n *Node) AddPath(values ...interface{}) *Node {
	if len(values) == 0 {
		return nil
	}

	current := n
	for _, value := range values {
		if node := current.Find(value); node == nil {
			previous := current
			current = New(value)
			current.Virtual = true
			previous.Nodes = append(previous.Nodes, current)
		} else {
			current = node
		}
	}
	current.Virtual = false
	return current
}

func (n *Node) AddPathString(path string) *Node {
	edges := strings.Split(path, "/")
	values := make([]interface{}, len(edges))
	for i := range edges {
		values[i] = edges[i]
	}
	return n.AddPath(values...)
}

// Find finds the child node with the target value.
// Nil if not found.
func (n *Node) Find(value interface{}) *Node {
	for _, node := range n.Nodes {
		if reflect.DeepEqual(node.Value, value) {
			return node
		}
	}
	return nil
}
