package main

import "math/rand"

func main() {

	const arraySize = 5

	var arrayInt [arraySize]int

	for element := 0; element < arraySize; element++ {

		arrayInt[element] = rand.Intn(arraySize)
	}
	println("Random Array")
	for i := 0; i < len(arrayInt); i++ {
		println(arrayInt[i])
	}
	println("---")
	var sortedArray = SelectionSort(arrayInt[:])

	println("Sorted Array")
	for i := 0; i < len(sortedArray); i++ {
		println(sortedArray[i])
	}

}

func findSmallestIndex(sortArray []int) int {

	var smallest = sortArray[0]
	var smallest_index = 0

	for i := 0; i < len(sortArray); i++ {
		if sortArray[i] < smallest {
			smallest = sortArray[i]
			smallest_index = i
		}
	}

	return smallest_index
}

func removeAtIndex(array []int, index int) []int {
	//moves the last position to target index
	array[index] = array[len(array)-1]
	//returns slice without last position
	return array[:len(array)-1]
}

func SelectionSort(sortArray []int) []int {
	var sortedArray []int
	var length = len(sortArray)
	for i := 0; i < length; i++ {
		var smallest_index = findSmallestIndex(sortArray)
		sortedArray = append(sortedArray, sortArray[smallest_index])
		sortArray = removeAtIndex(sortArray, smallest_index)
	}

	return sortedArray
}
