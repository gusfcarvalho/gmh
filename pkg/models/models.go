// Models
// This code is licensed under the Apache License, Version 2.0
package models

// Node represents a node in the graph
type Node struct {
	ID        string
	Value     interface{}
	Label     string
	Neighbors []*Node
	Children  []*Node // Reference to the parent node
}

// Graph represents a graph
type Graph struct {
	Nodes []*Node
}
