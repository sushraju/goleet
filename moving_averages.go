package main

import "fmt"

// MovingAverage struct definition
// windowSize, an array of values and position of insertion
type MovingAverage struct {
	windowSize     int
	values         []int
	insertPosition int
}

// AddValue adds the value to the values list and increment the position until the window gets filled
// After that drop the first and append to the end.
func (mva *MovingAverage) AddValue(value int) {
	if mva.insertPosition < mva.windowSize {
		mva.values[mva.insertPosition] = value
		mva.insertPosition = mva.insertPosition + 1
	} else {
		mva.Pop(0)
		mva.values[mva.insertPosition-1] = value
	}
}

// GetAverage gets the current average in the moving window
func (mva MovingAverage) GetAverage() float64 {
	var sumValues int
	var movingAverage float64

	for _, num := range mva.values {
		sumValues = sumValues + num
	}

	if sumValues > 0 && mva.insertPosition > 0 {
		movingAverage = float64(sumValues) / float64(mva.insertPosition)
	}

	return movingAverage
}

//Pop an element from the list
func (mva MovingAverage) Pop(index int) []int {
	return append(mva.values[:index], mva.values[index+1:]...)
}

// New instance of MovingAverage
func New(windowSize int) *MovingAverage {
	return &MovingAverage{
		windowSize:     windowSize,
		values:         make([]int, windowSize),
		insertPosition: 0,
	}
}

func main() {

	// Driver code
	ma := New(6)
	valuesArray := [...]int{23, 45, 34, 56, 78, 98, 9, 90, 45, 18, 99, 85, 91, 101, 97, 110, 100, 92, 81, 77, 73, 69, 85}

	for _, value := range valuesArray {
		ma.AddValue(value)
		fmt.Print("\n Adding value ", value, " and here's the moving window and average: ", ma.values)
		fmt.Printf(" %.2f", ma.GetAverage())
	}
}
