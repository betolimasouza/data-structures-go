package main

func main() {

	var arrayInt [1000]int

	for element := 0; element < 1000; element++ {

		arrayInt[element] = element + 1
	}

	binarySearch(arrayInt[:], 1000)
	binarySearch(arrayInt[:], 1)
	binarySearch(arrayInt[:], 1001)
}

func binarySearch(searchArray []int, target int) int {
	var low = 0
	var high = len(searchArray) - 1

	for low <= high {
		var middle = (low + high) / 2
		var guess = searchArray[middle]

		println("Middle: ", middle, "High:", high, "Low:", low, "Guess:", guess)

		if guess == target {
			println("Found target: ", guess)
			return guess
		}

		if guess < target {
			low = middle + 1
			println("Guess is lower than target", "Middle: ", middle, "High:", high, "Low:", low, "Guess:", guess)
		}

		if guess > target {
			high = middle - 1
			println("Guess is higher than target", "Middle: ", middle, "High:", high, "Low:", low, "Guess:", guess)
		}
	}
	println("Target not found in array")
	return 0
}
