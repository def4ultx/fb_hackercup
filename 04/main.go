package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		runtest(i)
	}
}

func runtest(tc int) {
	var n int
	fmt.Scanf("%d", &n)

	ores := make([]int, n)
	for i := range ores {
		fmt.Scanf("%d", &ores[i])
	}

	tunnels := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		var src, dst int
		fmt.Scanf("%d %d", &src, &dst)

		_, ok := tunnels[src]
		if ok {
			tunnels[src] = append(tunnels[src], dst)
		} else {
			tunnels[src] = []int{dst}
		}

		_, ok = tunnels[dst]
		if ok {
			tunnels[dst] = append(tunnels[dst], src)
		} else {
			tunnels[dst] = []int{src}
		}
	}

	// for k, v := range tunnels {
	// 	fmt.Print(k, ":")
	// 	for _, t := range v {
	// 		fmt.Print(" ", t)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("---- tunnels ----")

	seen := map[int]struct{}{
		1: {},
	}

	ts, ok := tunnels[1]
	if !ok {
		fmt.Printf("Case #%d: %d\n", tc+1, ores[0])
		return
	}

	result := make([]int, 0)
	for _, v := range ts {
		sum := MaxSum(v, tunnels, ores, seen, map[int]int{})
		result = append(result, sum)
	}

	var answer int
	if len(result) == 1 {
		answer = result[0] + ores[0]
	} else if len(result) > 1 {
		m1, m2 := Max(result)
		answer = m1 + m2 + ores[0]
	} else {
		fmt.Printf("Case #%d: Error\n", tc+1)
		return
	}

	fmt.Printf("Case #%d: %d\n", tc+1, answer)
	return
}

func Max(arr []int) (int, int) {
	var max, secondMax int
	if arr[0] > arr[1] {
		max = arr[0]
		secondMax = arr[1]
	} else {
		max = arr[1]
		secondMax = arr[0]
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] > secondMax {
			if arr[i] <= max {
				secondMax = arr[i]
			} else {
				secondMax, max = max, arr[i]
			}
		}
	}
	return max, secondMax

}

func MaxSum(src int, tunnels map[int][]int, ores []int, seen map[int]struct{}, memory map[int]int) int {

	if ore, ok := memory[src]; ok {
		return ore
	}

	if _, ok := seen[src]; ok {
		return 0
	}

	ts, ok := tunnels[src]
	if !ok {
		return 0
	}

	newSeen := make(map[int]struct{})
	for k, v := range seen {
		newSeen[k] = v
	}
	newSeen[src] = struct{}{}

	result := make([]int, 0)
	for _, v := range ts {
		sum := MaxSum(v, tunnels, ores, newSeen, memory)
		result = append(result, sum)
	}

	var max int
	for _, v := range result {
		if v > max {
			max = v
		}
	}

	memory[src] = max + ores[src-1]
	return memory[src]
}
