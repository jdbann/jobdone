package person

import "math/rand"

func randomStep(m Person, mapWidth, mapHeight int) Person {
	m.x = constrain(0, mapWidth-1, m.x+randomDistance())
	m.y = constrain(0, mapHeight-1, m.y+randomDistance())
	return m
}

var randomDistance = func() func() int {
	var stepDistances = []int{-1, 0, 0, 0, 1}

	return func() int {
		return stepDistances[rand.Intn(5)]
	}
}()

func constrain(min, max, val int) int {
	if val < min {
		return min
	}

	if val > max {
		return max
	}

	return val
}
