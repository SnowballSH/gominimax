package gominimax

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMinimax(t *testing.T) {
	tree := CreateNode([]*Node{ // max 6
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
		CreateNode([]*Node{ // min 6
			CreateNode([]*Node{ // max 6
				CreateRootNode([]float64{ // min 6
					6,
				}),
				CreateRootNode([]float64{ // min 6
					6,
					9,
				}),
			}),
			CreateNode([]*Node{ // max 7
				CreateRootNode([]float64{ // min 7
					7,
				}),
			}),
		}),
		CreateNode([]*Node{ // min 5
			CreateNode([]*Node{ // max 5
				CreateRootNode([]float64{ // min 5
					5,
				}),
			}),
			CreateNode([]*Node{ // max 8
				CreateRootNode([]float64{ // min 8
					9,
					8,
				}),
				CreateRootNode([]float64{ // min 6
					6,
				}),
			}),
		}),
	})

	tree.FriendlyMinimax(99)
	assert.Equal(t, float64(6), *tree.Value)

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

	tree.FriendlyMinimax(99)
	assert.Equal(t, float64(-7), *tree.Value)

	tree = CreateNode([]*Node{ // max 5
		CreateNode([]*Node{ // min 5
			CreateRootNode([]float64{ // max 5
				3,
				5,
			}),
			CreateRootNode([]float64{ // max 6
				6,
				9,
			}),
		}),
		CreateNode([]*Node{ // min 2
			CreateRootNode([]float64{ // max 2
				1,
				2,
			}),
			CreateRootNode([]float64{ // max 0
				0,
				-1,
			}),
		}),
	})

	tree.FriendlyMinimax(99)
	assert.Equal(t, float64(5), *tree.Value)

	tree = CreateNode([]*Node{ // max 5
		CreateNode([]*Node{ // min 5
			CreateRootNode([]float64{ // max 5
				3,
				5,
			}),
			CreateRootNode([]float64{ // max 6
				6,
				9,
			}),
		}),
		CreateNode([]*Node{ // min 2
			CreateRootNode([]float64{ // max 2
				1,
				2,
			}),
			CreateRootNode([]float64{ // max 0
				0,
				-1,
			}),
		}),
	})

	tree.FriendlyMinimax(1)
	assert.Equal(t, float64(0), *tree.Value)
}
