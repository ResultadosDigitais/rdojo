package bowling

import (
	"testing"
)

func TestBowlingScore(test *testing.T) {
	tests := []struct {
		Name          string
		Roll          []int
		ExpectedScore int
		ExpectedError error
	}{
		{
			Name:          "Incomplete game returns error",
			Roll:          []int{1},
			ExpectedError: ErrIncompletedGame,
		},
		{
			Name:          "All zero game, output is zero",
			Roll:          []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			ExpectedError: nil,
			ExpectedScore: 0,
		},
		{
			Name:          "With no spares or strikes, output is sum",
			Roll:          []int{4, 0, 4, 0, 1, 0, 0, 6, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0},
			ExpectedError: nil,
			ExpectedScore: 23,
		},
		{
			Name:          "Spares double the next play",
			Roll:          []int{4, 6, 4, 0, 1, 0, 0, 6, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0},
			ExpectedError: nil,
			ExpectedScore: 33,
		},
	}

	for _, testCase := range tests {
		test.Run(testCase.Name, func(t *testing.T) {
			game := &Game{}
			for _, item := range testCase.Roll {
				game.Roll(item)
			}
			score, err := game.Score()
			if testCase.ExpectedError != err {
				t.Errorf("Unexpected error: %v, expected: %v", err, testCase.ExpectedError)
			}
			if score != testCase.ExpectedScore {
				t.Errorf("Unexpected score: %v, expected: %v", score, testCase.ExpectedScore)
			}
		})
	}
}

func TestCurrentRoll(test *testing.T) {
	game := NewGame()
	if game.currentFrame == nil {
		test.Error("Expected game current frame not be nil")
	}
}

// func TestRoll(test *testing.T) {
// 	tests := []struct {
// 		Name string
// 	}{
// 		{
// 			Name:           "Is not a spare and not a strike",
// 			Roll:           Roll{Attempts: [2]int{1, 2}},
// 			ExpetcedScore:  3,
// 			ExpectedSpare:  false,
// 			ExpectedStrike: false,
// 		},
// 		{
// 			Name:           "Is spare",
// 			Roll:           Roll{Attempts: [2]int{4, 6}},
// 			ExpetcedScore:  10,
// 			ExpectedSpare:  true,
// 			ExpectedStrike: false,
// 		},
// 		{
// 			Name:           "Is strike",
// 			Roll:           Roll{Attempts: [2]int{10, 0}},
// 			ExpetcedScore:  10,
// 			ExpectedSpare:  false,
// 			ExpectedStrike: true,
// 		},
// 	}

// 	for _, testCase := range tests {
// 		test.Run(testCase.Name, func(t *testing.T) {
// 			if tesCase.Roll.Score != testCase.ExpetcedScore {
// 				t.Errorf("Error! Got: %v, expected: %v", tesCase.Roll.Score, testCase.ExpectedError)
// 			}
// 			if tesCase.Roll.IsSpare != testCase.ExpectedSpare {
// 				t.Errorf("Error! Got: %v, expected: %v", tesCase.Roll.IsSpare, testCase.ExpectedSpare)
// 			}
// 			if tesCase.Roll.IsStrike != testCase.ExpectedStrike {
// 				t.Errorf("Error! Got: %v, expected: %v", tesCase.Roll.IsStrike, testCase.ExpectedStrike)
// 			}
// 		})
// 	}
// }
