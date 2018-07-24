package main

import "fmt"

/*
LeetCode Problem 3: Longest Substring Without Repeating Character
 */

func solve(str string) int{
	// decalre a map
	lastOccured := make(map[byte]int)
	start := 0
	maxLen := 0

	for i, ch := range []byte(str){
		if lastIdx, ok  := lastOccured[ch]; ok && lastIdx >= start{
			start = lastOccured[ch] + 1
		}
		if maxLen < i - start + 1{
			maxLen = i - start + 1
		}
		// update lastOccured
		lastOccured[ch] = i
	}
	return maxLen
}

// solve with international character
func solveInt(str string) int{
	// decalre a map
	lastOccured := make(map[rune]int)
	start := 0
	maxLen := 0

	for i, ch := range []rune(str){
		if lastIdx, ok  := lastOccured[ch]; ok && lastIdx >= start{
			start = lastOccured[ch] + 1
		}
		if maxLen < i - start + 1{
			maxLen = i - start + 1
		}
		// update lastOccured
		lastOccured[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(solve("abcabb"))
	fmt.Println(solveInt("abcabb"))
	fmt.Println(solve("吃葡萄不吐葡萄皮"))
	fmt.Println(solveInt("吃葡萄不吐葡萄皮"))
	fmt.Println((solveInt("黑化肥发灰会挥发灰化肥挥发会发黑")))
}
