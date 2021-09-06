
	results := make([]int, len(tunnels))
	for i := x + 1; i < len(tunnels); i++ {
		// if i == y {
		// 	continue
		// }

		sum := maxSum(x, i, ores, tunnels, seen)

		fmt.Println("loop after maxSum")
		fmt.Println(y, x, i)
		fmt.Println(sum)

		result := sum
		results = append(results, result)
	}

	var max int
	for _, v := range results {
		if v > max {
			max = v
		}
	}
	seen[y][x] = max + ores[y]

	fmt.Println("xy result", y, x, max, ores[y])
	return max + ores[y]