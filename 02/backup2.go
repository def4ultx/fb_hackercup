package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		runtest(i)
	}
}

func runtest(tc int) {
	var str string
	fmt.Scanf("%s", &str)

	var n int
	fmt.Scanf("%d", &n)

	var (
		mapping      = make(map[rune][]rune)
		costs        = make(map[string]int)
		destinations = make(map[rune]struct{})
	)
	for i := 0; i < n; i++ {
		var replacement string
		fmt.Scanf("%s", &replacement)

		src, _ := utf8.DecodeRune([]byte{replacement[0]})
		dst, _ := utf8.DecodeRune([]byte{replacement[1]})
		if _, ok := mapping[src]; ok {
			mapping[src] = append(mapping[src], dst)
		} else {
			mapping[src] = []rune{dst}
		}
		costs[replacement] = 1
		destinations[dst] = struct{}{}
	}

	// check if all same character, no need to calculate
	isSameRune := true
	char, _ := utf8.DecodeRune([]byte{str[0]})
	for _, v := range str {
		if v != char {
			isSameRune = false
			break
		}
	}
	if isSameRune {
		fmt.Printf("Case #%d: 0\n", tc+1)
		return
	}

	for _, v := range str {
		destinations[v] = struct{}{}
	}

	var result []int
	for dst := range destinations {
		var total int
		for _, v := range str {
			if v == dst {
				continue
			}

			cost := findCost(v, dst, mapping, costs, false, map[rune]struct{}{})
			if cost == -1 {
				total = cost
				break
			} else {
				total += cost
			}
		}

		if total != -1 {
			fmt.Println(string(dst), total)
			result = append(result, total)
		}
	}

	min := math.MaxInt32
	for _, v := range result {
		if v < min {
			min = v
		}
	}

	if min == math.MaxInt32 {
		min = -1
	}

	fmt.Printf("Case #%d: %d\n", tc+1, min)
	return
}

type Pair struct {
	Character rune
	Cost      int
}

func findCost(src, dst rune, mapping map[rune][]rune, costs map[string]int, canMap bool, cyclic map[rune]struct{}) int {

	key := string(src) + string(dst)
	cost, ok := costs[key]
	if ok {
		return cost
	}

	runes, ok := mapping[src]
	if !ok {

		// fmt.Println("return not found")
		costs[key] = -1
		return -1
	}

	_, ok = cyclic[src]
	if ok {
		// fmt.Println("return not found 3, cyclic")
		costs[key] = -1
		return -1
	}

	newCyclic := make(map[rune]struct{})
	newCyclic[src] = struct{}{}
	for k, v := range cyclic {
		newCyclic[k] = v
	}

	// Need cyclic detection ??
	pairs := make([]Pair, 0)
	for _, r := range runes {
		c := findCost(rune(r), dst, mapping, costs, true, newCyclic)

		if c != -1 {
			pair := Pair{
				Character: rune(r),
				Cost:      c,
			}
			pairs = append(pairs, pair)
		}
	}

	// first find cannot be mapping !!
	if len(pairs) == 0 {
		// fmt.Println("return not found 2")
		return -1 // Not found
	}

	min := pairs[0].Cost
	for _, v := range pairs {
		if v.Cost < min {
			min = v.Cost
		}
	}

	if canMap {

		costs[key] = min + 1
		// fmt.Println("create mapping", key, costs[key])
	}

	return min + 1
}
