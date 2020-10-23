package helper

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"sort"
	"time"
)

// Initialize sets current node's initial value y and round r.
func Initialize() (float64, int) {

	// initialize y
	rand.Seed(time.Now().UnixNano())
	y := rand.Float64()

	//initialize r
	r := 1

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

// CheckState checks the array of states and determines if they have reached consensus
func CheckState(values []float64) bool {

	sort.Float64s(values)

	if len(values) == 0 {
		return false
	}

	if values[len(values)-1]-values[0] > 0.001 {
		return false
	}

	return true
}

// NodeCrash simulates a % chance that a node fails
func NodeCrash() {
	//20% chance to crash
	num := rand.Intn(100)
	if num <= 20 {
		fmt.Println("NODE CRASHED!")
		os.Exit(0)
	}
}

// DisplayMap displays the connected nodes in a readable format
func DisplayMap(nodes map[string]net.Conn) {
	fmt.Println("---Map---")
	for key, value := range nodes {
		fmt.Println("ID:", key, ", net.Conn:", value)
	}
	fmt.Println("---------")
}

// PrintRoundInfo prints y value, round #, and states received to the console
func PrintRoundInfo(y float64, r int, states []float64) {
	fmt.Printf("****Round: %v****\n", r)
	fmt.Printf("***y: %v\n", y)
	fmt.Println(states)
}

// PrintServerRoundInfo round # and received values
func PrintServerRoundInfo(r int, values []float64) {
	fmt.Printf("*****Round: %v*****\n", r)
	fmt.Println(values)
}
