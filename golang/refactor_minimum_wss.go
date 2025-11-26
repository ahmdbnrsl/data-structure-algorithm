package main

import (
	"fmt"
	"math"
)

func minWindow(s, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	targetFreq := make(map[rune]int)
	for _, c := range t {
		targetFreq[c]++
	}

	left := 0
	minLen := math.MaxInt32
	start := 0
	matchCount := 0

	windowFreq := make(map[rune]int)
	sRunes := []rune(s)

	for right, char := range sRunes {
		windowFreq[char]++

		if targetFreq[char] > 0 && windowFreq[char] <= targetFreq[char] {
			matchCount++
		}

		for matchCount == len(t) {
			// Update minimum window
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			// Move left to shrink window
			leftChar := sRunes[left]
			windowFreq[leftChar]--
			if targetFreq[leftChar] > 0 && windowFreq[leftChar] < targetFreq[leftChar] {
				matchCount--
			}
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return string(sRunes[start : start+minLen])
}

func main() {
	fmt.Println(minWindow("AAABBCA", "AABCA")) // Output: "BANC"
}