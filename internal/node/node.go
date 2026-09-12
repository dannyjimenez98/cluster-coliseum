// Package node includes core node behavior for the program
package node

type Node struct {
	id string
	health int
}

// NewNode creates a new node at default full health 
func NewNode(id string) *Node {
	return &Node{
		id: id,
		health: 100,
	}
}
