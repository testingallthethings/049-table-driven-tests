package bowling

import (
	"strconv"
	"strings"
)

func Score(game string) int {
	score := 0

	frames := strings.Split(game, " ")

	for _, frame := range frames {
		balls := strings.Split(frame, "")
		if balls[0] != "-" {
			points, _ := strconv.Atoi(balls[0])
			score += points
		}

		if balls[1] != "-" {
			score++
		}
	}

	return score
}
