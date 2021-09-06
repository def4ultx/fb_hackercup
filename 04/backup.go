package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		runtest(i)
	}

	// runtest(0)
}

func runtest(tc int) {
	var n int
	fmt.Scanf("%d", &n)

	ores := make([]int, n)
	for i := range ores {
		fmt.Scanf("%d", &ores[i])
	}

	tunnels := make([][]int, n)
	for i := 0; i < n; i++ {
		tunnels[i] = make([]int, n)
	}

	// fmt.Println(tunnels)

	for i := 0; i < n-1; i++ {
		var src, dst int
		fmt.Scanf("%d %d", &src, &dst)

		fmt.Println(src, dst)
		tunnels[src-1][dst-1] = 1
		tunnels[dst-1][src-1] = 1
	}

	// 2 max sum from all path from 1
	for _, ts := range tunnels {
		fmt.Println(ts)
	}
	fmt.Println("----")

	MaxSum(ores, tunnels)
	return
}

func MaxSum(ores []int, tunnels [][]int) int {
	size := len(tunnels)

	memory := make([][]int, size)
	for i := range memory {
		memory[i] = make([]int, size)
		for j := range memory[i] {
			memory[i][j] = -1
		}
	}

	// for i := 1; i < size; i++ {
	// 	sum := maxSum(0, i, ores, tunnels, memory)
	// 	fmt.Println("result at", 0, i, ":", sum)
	// }

	i := 4
	sum := maxSum(0, i, ores, tunnels, memory)
	fmt.Println("result at", 0, i, ":", sum)

	return -1
}

func maxSum(y, x int, ores []int, tunnels, memory [][]int) int {
	// if memory[y][x] != -1 {
	// 	return memory[y][x]
	// }

	size := len(tunnels)
	results := make([]int, size)
	for i := 0; i < size; i++ {
		if tunnels[x][i] == 0 {
			continue
		}
		if i == y {
			continue
		}

		fmt.Println("try maxSum", x, i)

		sum := maxSum(x, i, ores, tunnels, memory)
		fmt.Println("test", sum)

		result := sum
		results = append(results, result)
	}

	var max int
	for _, v := range results {
		if v > max {
			max = v
		}
	}

	// fmt.Println("save memory at", y, x, "with: ", max, " ores", ores[y])
	// memory[y][x] = max + ores[y]

	return max + ores[y]
}
