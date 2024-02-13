package pkg

import (
	"sort"
)

// Knapsack is a function to calculate the minimum number of items to make the target index
func Knapsack(items []int, capacity int) (int, int, map[int]int) {
	// Sort the items in ascending order
	sort.Ints(items)

	// Calculate the maximum capacity which may be larger than given capacity;
	maxCapacity := capacity + items[0]

	// dp array to store the minimum number of items to make the current index;
	// init dp array and usedItem array;
	dp := make([]int, maxCapacity+1)
	usedItem := make([]int, maxCapacity+1)

	dp[0] = 1
	usedItem[0] = -1
	for _, item := range items {
		if item <= maxCapacity {
			dp[item] = 1
			usedItem[item] = item
		}
	}

	// Calculate the minimum number of items to make the current index;
	for _, item := range items {
		for i := 1; i+item <= maxCapacity; i++ {
			// if the current index have been made and
			// the current item can make the target index with less than number of items or target index hasn't been made
			if (dp[i] > 0) && (dp[i+item] > dp[i]+1 || dp[i+item] == 0) {
				dp[i+item] = dp[i] + 1
				usedItem[i+item] = item
			}
		}
	}

	// Find the minimum amount we can make with the given items;
	var minItemsAmount int
	for i := capacity; i <= maxCapacity; i++ {
		if dp[i] > 0 {
			minItemsAmount = i
			break
		}
	}

	// Init usedItemCounter to count the number of items used, defaults to 0;
	usedItemsMap := make(map[int]int)
	for _, item := range items {
		usedItemsMap[item] = 0
	}

	// Traceback usedItems
	start := minItemsAmount
	for start > 0 {
		usedItemsMap[usedItem[start]]++
		start -= usedItem[start]
	}

	// Removing unused items when filling the knapsack
	for k, v := range usedItemsMap {
		if v == 0 {
			delete(usedItemsMap, k)
		}
	}

	return minItemsAmount, dp[minItemsAmount], usedItemsMap
}
