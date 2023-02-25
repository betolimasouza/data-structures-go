package main

import "math/rand"

func main() {
	const arraySize = 100

	var arrayInt [arraySize]int

	for element := 0; element < arraySize; element++ {
		arrayInt[element] = rand.Intn(arraySize)
	}
	println("Random Array:")
	for i := 0; i < len(arrayInt); i++ {
		println(arrayInt[i])
	}

	// var result = sum(arrayInt[:])

	// println("Array sum: ", result)

	//////////
	println("Maximum Number in Array:", maximumNumber(arrayInt[:]))
}

func sum(array []int) int {
	var result = 0

	if len(array) == 1 {
		result += array[0]
	} else {
		result += array[0]
		result += sum(array[1:])
	}

	return result
}

func maximumNumber(array []int) int {
	var first = 0
	var last = len(array) - 1

	if len(array) == 2 {
		if array[0] > array[1] {
			return array[0]
		} else {
			return array[1]
		}
	}

}
