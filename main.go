package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type ProgressBar struct {
	total          int64   //total number of elements
	current        int64   //current elements
	percentage     float64 //cyrretn percentage
	char           string  //character/Symbol for ProgressBar
	totalBarSize   int64   //total no of character to depict 100%
	currentBarSize int64   //current number of Characters to represent bars
}

func getNewBar(totalElements int64, currentElements int64, symbol string, totalBarSize int64) ProgressBar {

	CurrentBar := int64(math.Ceil(float64(currentElements/totalElements) * 50))
	bar := ProgressBar{
		total:          totalElements,
		current:        currentElements,
		char:           symbol,
		percentage:     float64(currentElements/totalElements) * 100,
		totalBarSize:   totalBarSize,
		currentBarSize: CurrentBar}

	return bar
}

func (bar *ProgressBar) display(current int64) {

	//update the current status
	bar.current = current

	//calculate the new current percentage
	bar.percentage = (float64(bar.current) * 100.0) / float64(bar.total)

	//calculate the number of Symbols to display
	bar.currentBarSize = int64(bar.percentage * float64(bar.totalBarSize) / 100)

	//Generate the ProgressBar string
	currentBar := strings.Repeat(bar.char, int(bar.currentBarSize))

	fmtString := fmt.Sprintf("\r[%%-%ds] %%3.2f%%%% %%8d/%%d", bar.totalBarSize)
	//fmt.Printf("\r[%-50s] %3.2f%% %8d/%d", currentBar, bar.percentage, bar.current, bar.total)
	fmt.Printf(fmtString, currentBar, bar.percentage, bar.current, bar.total)
}

func (bar *ProgressBar) end() {
	fmt.Println()
}
func main() {

	var N int64 = 75
	bar := getNewBar(N, 0, "#", 70)

	for i := 0; i <= int(N); i++ {

		time.Sleep(100 * time.Millisecond)
		bar.display(int64(i))
	}

	bar.end()

}
