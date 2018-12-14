package bowling

type Frame struct {
	Attempts []int
}

func (g *Frame) Score() int {
	score := 0
	for _, attempt := range g.Attempts {
		score += attempt
	}

	return score
}

// func (g *Frame) IsSpare() bool {
// 	return g
// }
