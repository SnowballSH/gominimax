// Go Minimax is a simple implementation of the minimax algorithm with alpha-beta pruning in Go
package gominimax

import (
	"math"
)

// Node represents a node in the tree
type Node struct {
	Value    *float64
	Children []*Node
}

// CreateNode creates a parent node
func CreateNode(nodes []*Node) *Node {
	return &Node{
		Value:    nil,
		Children: nodes,
	}
}

// CreateRootNode creates a root node containing a float64.
func CreateRootNode(value []float64) *Node {
	arr := make([]*Node, len(value))
	for i, y := range value {
		w := y
		arr[i] = &Node{
			Value:    &w,
			Children: nil,
		}
	}
	return &Node{
		Value:    nil,
		Children: arr,
	}
}

// Minimax calculates minimax for the node. This method fills n.Value from null to an float64 pointer.
// The mode argument is either 0 or 1, 0 means MAX and 1 means MIN.
func (n *Node) Minimax(mode int8) {
	if n.Value != nil {
		return
	}

	var channels = make([]chan bool, len(n.Children))
	for i, w := range n.Children {
		ch := make(chan bool)
		channels[i] = ch
		go func(x *Node, c chan bool) {
			x.Minimax(mode ^ 1)
			c <- true
		}(w, ch)
	}

	var lowest float64

	if mode == 0 {
		lowest = math.Inf(-1)
	} else {
		lowest = math.Inf(1)
	}

	for i, w := range channels {
		<-w
		res := *n.Children[i].Value // should never be nil

		if mode == 0 {
			lowest = math.Max(res, lowest)
		} else {
			lowest = math.Min(res, lowest)
		}
	}

	n.Value = &lowest
}
