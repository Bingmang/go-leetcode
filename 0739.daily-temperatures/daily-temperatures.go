package _739_daily_temperatures

func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)

	stack := make([]int, n)
	top := -1

	for i, temperature := range(T) {
		for top >= 0 && T[stack[top]] < temperature {
			res[stack[top]] = i - stack[top]
			top--
		}

		top++
		stack[top] = i
	}

	return res
}
