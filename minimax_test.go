package gominimax

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMinimax(t *testing.T) {
	tree := CreateNode([]*Node{ // max 3
		CreateNode([]*Node{ // min 3
			CreateNode([]*Node{ // max 5
				CreateRootNode([]float64{ // min 5
					5,
					6,
				}),
				CreateRootNode([]float64{ // min 4
					7,
					4,
					5,
				}),
			}),
			CreateNode([]*Node{ // max 3
				CreateRootNode([]float64{ // min 3
					3,
				}),
			}),
		}),
	})

	tree.Minimax(0)
	assert.Equal(t, float64(3), *tree.Value)

	tree = CreateNode([]*Node{ // max -7
		CreateNode([]*Node{ // min -10
			CreateNode([]*Node{ // max 10
				CreateRootNode([]float64{ // min 10
					10,
					math.Inf(1),
				}),
				CreateRootNode([]float64{ // min 5
					5,
				}),
			}),
			CreateNode([]*Node{ // max -10
				CreateRootNode([]float64{ // min -10
					-10,
				}),
			}),
		}),
		CreateNode([]*Node{ // min -7
			CreateNode([]*Node{ // max 5
				CreateRootNode([]float64{ // min 5
					7,
					5,
				}),
				CreateRootNode([]float64{ // min -inf
					math.Inf(-1),
				}),
			}),
			CreateNode([]*Node{ // max -7
				CreateRootNode([]float64{ // min -7
					-7,
					-5,
				}),
			}),
		}),
	})

	tree.Minimax(0)
	assert.Equal(t, float64(-7), *tree.Value)
}
