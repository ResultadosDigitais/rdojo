package bowling

type Game struct {
	rolls        [11]Frame
	currentFrame *Frame
	round        int
	ball         int
}

func NewGame() *Game {
	return &Game{currentFrame: &Frame{}}
}

func (g *Game) Roll(pinCount int) {
	//	g.rolls = append(g.rolls[:], pinCount)
	/*if ball == 0 && pinCount == 10 {
		// strike!
	}
	if round > 1 {
		g[round] = append(g.rolls[round].Attempts, pinCount)

	}*/

}

func (g *Game) Score() (int, error) {
	if len(g.rolls) < 20 {
		return 0, ErrIncompletedGame
	}
	var score int

	for _, roll := range g.rolls {
		score += roll.Score()
	}

	return score, nil
}
