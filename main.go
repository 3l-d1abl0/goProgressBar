package main

type ProgressBar struct {
	total        int64   //total number of elements
	current      int64   //current elements
	percentage   float32 //cyrretn percentage
	char         string  //character/Symbol for ProgressBar
	totalBarSize int64   //no of character to depict 100%
}

func getNewBar(totalElements int64, currentElements int64, symbol string) ProgressBar {

	bar := ProgressBar{
		total:        totalElements,
		current:      currentElements,
		char:         symbol,
		percentage:   float32(currentElements / totalElements),
		totalBarSize: 50}

	return bar
}
func main() {

	bar := getNewBar(100, 0, "#")

}
