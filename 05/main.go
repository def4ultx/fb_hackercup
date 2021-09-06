package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		runtest(i)
	}
}

func runtest(tc int) {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

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
	sort.Ints(result)
	fmt.Println(result)

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Println(len(result)-i, k, i, len(result))
		if len(result)-i > k {
			break
		}
		answer += result[i]
	}
	answer += ores[0]

	fmt.Printf("Case #%d: %d\n", tc+1, answer)
	return
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
