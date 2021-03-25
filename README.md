## Minimax algorithm implemented in Go.

Install:
```bash
go get -u https://github.com/SnowballSH/gominimax
```

Usage:
```go
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

tree.FriendlyMinimax(99) // 5
```
