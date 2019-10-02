package person

import (
	"math/rand"
	"time"
)

// waitRandomTime waits a random amout of time (ms) and returns the duration
func waitRandomTime(min, max int) int {

	// New seed
	rand.Seed(time.Now().UnixNano())

	// Amout of time
	ms := rand.Intn(max-min) + min

	// Wait
	time.Sleep(time.Duration(ms) * time.Millisecond)

	return ms
}
