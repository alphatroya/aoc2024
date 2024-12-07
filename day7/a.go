package main

func aSumOfCorrectResults(data []data) int {
	sum := 0
	for _, d := range data {
		if calc(d.result, d.components[0], d.components[1:]) {
			sum += d.result
		}
	}
	return sum
}

func calc(result int, acc int, row []int) bool {
	if len(row) == 0 {
		return result == acc
	}
	return calc(result, acc+row[0], row[1:]) || calc(result, acc*row[0], row[1:])
}
