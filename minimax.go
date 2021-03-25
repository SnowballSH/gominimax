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
func (n *Node) Minimax(mode int8, depth int, alpha, beta float64) {
	if n.Value != nil || depth <= 0 {
		return
	}

	var channels = make([]chan bool, len(n.Children))
	for i, w := range n.Children {
		ch := make(chan bool)
		channels[i] = ch
		go func(x *Node, c chan bool) {
			x.Minimax(mode^1, depth-1, alpha, beta)
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
			alpha = math.Max(alpha, lowest)
		} else {
			lowest = math.Min(res, lowest)
			beta = math.Min(beta, lowest)
		}

		if beta <= alpha {
			break
		}
	}

	n.Value = &lowest
}

// Use this method for general minimax. This method is a shortcut for n.Minimax(0, depth, math.Inf(-1), math.Inf(1)).
func (n *Node) FriendlyMinimax(depth int) {
	n.Minimax(0, depth, math.Inf(-1), math.Inf(1))
}
