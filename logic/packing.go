package logic

import "sort"

// FindPacks calculates the minimum number of packs needed to fulfill order
// based on the parameters like order size and available pack sizes ( provided via API ).
// It returns the total number of items in the packs and a map of pack sizes to their counts.
func FindPacks(order int, packSizes []int) (int, map[int]int) {
    sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
    maxPack := max(packSizes)
    maxOrder := order + maxPack

    dp, parent := buildDPTable(order, packSizes, maxOrder)
    minTotal, _ := findBestSolution(order, maxOrder, dp)
    if minTotal == -1 {
        return 0, nil
    }
    packs := backtrackPacks(minTotal, parent)
    return minTotal, packs
}

// max returns the largest integer in a slice.
func max(nums []int) int {
    maxVal := 0
    for _, n := range nums {
        if n > maxVal {
            maxVal = n
        }
    }
    return maxVal
}

// buildDPTable populates DP and its parent tables.
func buildDPTable(order int, packSizes []int, maxOrder int) ([]int, [][2]int) {
    const INF = int(1e9)
    dp := make([]int, maxOrder+1)
    parent := make([][2]int, maxOrder+1)
    for i := range dp {
        dp[i] = INF
    }
    dp[0] = 0

    for t := 0; t <= maxOrder; t++ {
        if dp[t] == INF {
            continue
        }
        for _, p := range packSizes {
            newT := t + p
            if newT <= maxOrder && dp[newT] > dp[t]+1 {
                dp[newT] = dp[t] + 1
                parent[newT] = [2]int{t, p}
            }
        }
    }
    return dp, parent
}

// findBestSolution returns the minimal total and minimal packs.
func findBestSolution(order, maxOrder int, dp []int) (int, int) {
    const INF = int(1e9)
    minTotal := -1
    minPacks := INF
    for t := order; t <= maxOrder; t++ {
        if dp[t] < INF {
            if minTotal == -1 || t < minTotal || (t == minTotal && dp[t] < minPacks) {
                minTotal = t
                minPacks = dp[t]
            }
        }
    }
    return minTotal, minPacks
}

// backtrackPacks reconstructs the pack breakdown from the parent table.
func backtrackPacks(minTotal int, parent [][2]int) map[int]int {
    packs := make(map[int]int)
    curr := minTotal
    for curr > 0 {
        prev, usedPack := parent[curr][0], parent[curr][1]
        packs[usedPack]++
        curr = prev
    }
    return packs
}