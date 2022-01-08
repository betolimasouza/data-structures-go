package main

import "fmt"

func main() {

	var square int
	var cube int
	var error error

	square, cube, error = powerSeries(3)

	if error == nil {
		fmt.Println("Square:", square, "Cube:", cube)
	}

}

func powerSeries(a int) (int, int, error) {

	var square = a * a

	var cube = a * a * a

	return square, cube, nil

}
