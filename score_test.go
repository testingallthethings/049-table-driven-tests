package bowling_test

import (
	"bowling"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGutterBallGame(t *testing.T) {
	assert.Equal(t, 0, bowling.Score("-- -- -- -- -- -- -- -- -- --"))
}

func TestFirstPinNumericPoints(t *testing.T) {
	tests := []struct {
		name          string
		game          string
		expectedScore int
	}{
		{"Single pin first frame", "1- -- -- -- -- -- -- -- -- --", 1},
		{"Single pin first two frames", "1- 1- -- -- -- -- -- -- -- --", 2},
		{"Single pin first three frames", "1- 1- 1- -- -- -- -- -- -- --", 3},
		{"First ball simple scores", "-- 1- 2- 3- 4- 5- 6- 7- 8- 9-", 45},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedScore, bowling.Score(test.game))
		})
	}
}

func TestSecondBallNumericPoints(t *testing.T) {
	tests := []struct {
		name          string
		game          string
		expectedScore int
	}{
		{"Single pin first frame", "-1 -- -- -- -- -- -- -- -- --", 1},
		{"Single pin first two frames", "-1 -1 -- -- -- -- -- -- -- --", 2},
		{"Single pin first three frames", "-1 -1 -1 -- -- -- -- -- -- --", 3},
		// TODO add second ball simple score tests
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedScore, bowling.Score(test.game))

		})
	}

	// TODO start to cover spares and strikes
}
