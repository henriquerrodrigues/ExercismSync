package bookstore

import (
	"sort"
	"strconv"
	"strings"
)

func Cost(books []int) int {
	counts := make(map[int]int)
	for _, book := range books {
		counts[book]++
	}

	var quantities []int
	for _, q := range counts {
		quantities = append(quantities, q)
	}

	memo := make(map[string]int)

	prices := map[int]int{
		1: 800,
		2: 1520,
		3: 2160,
		4: 2560,
		5: 3000,
	}

	var solve func([]int) int
	solve = func(q []int) int {
		sort.Sort(sort.Reverse(sort.IntSlice(q)))
		
		for len(q) > 0 && q[len(q)-1] == 0 {
			q = q[:len(q)-1]
		}

		if len(q) == 0 {
			return 0
		}

		keyBuilder := strings.Builder{}
		for _, v := range q {
			keyBuilder.WriteString(strconv.Itoa(v))
			keyBuilder.WriteByte(',')
		}
		key := keyBuilder.String()

		if val, found := memo[key]; found {
			return val
		}

		minPrice := -1

		maxSetSize := 5
		if len(q) < 5 {
			maxSetSize = len(q)
		}

		for size := 1; size <= maxSetSize; size++ {
			nextQ := make([]int, len(q))
			copy(nextQ, q)

			for i := 0; i < size; i++ {
				nextQ[i]--
			}

			currentPrice := prices[size] + solve(nextQ)

			if minPrice == -1 || currentPrice < minPrice {
				minPrice = currentPrice
			}
		}

		memo[key] = minPrice
		return minPrice
	}

	return solve(quantities)
}