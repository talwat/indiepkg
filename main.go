package main

import "fmt"

func main() {
	err := fmt.Errorf("%s", "fefe")
	errorLog(err, 4, "An error occurred")
	//parseInput()
}
