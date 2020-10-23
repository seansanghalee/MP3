package helper

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Initialize sets current node's initial value y and round r.
func Initialize() (float64, int) {

	// initialize y
	rand.Seed(time.Now().UnixNano())
	y := rand.Float64()
	fmt.Printf("y: %v\n", y)

	//initialize r
	r := 1
	fmt.Printf("Round: %v\n", r)

	return y, r
}

// Average calculates the average of elements in array
func Average(values []float64) float64 {
	var sum float64

	for _, v := range values {
		sum = sum + v
	}

	return sum / float64(len(values))
}

// CheckState
func CheckState(values []float64) bool {
	fmt.Println("Checking state")

	sort.Float64s(values)
	if values[len(values)-1]-values[0] > 0.001 {
		return false
	}

	return true
}

func NodeCrash() bool {
	//2% chance to crash
	num := rand.Intn(100)
	if num <= 2 {
		return true
	}
	return false
}
