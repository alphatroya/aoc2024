package main

func bSumOfCorrectResults(data []data) int {
	sum := 0
	for _, d := range data {
		if bcalc(d.result, d.components[0], d.components[1:]) {
			sum += d.result
		}
	}
	return sum
}

func bcalc(result int, acc int, row []int) bool {
	if len(row) == 0 {
		return result == acc
	}
	return bcalc(result, acc+row[0], row[1:]) || bcalc(result, acc*row[0], row[1:]) || bcalc(result, concat(acc, row[0]), row[1:])
}

func concat(a, b int) int {
	power := 1
	for ba := b; ba > 0; ba /= 10 {
		power *= 10
	}
	return a*power + b
}
