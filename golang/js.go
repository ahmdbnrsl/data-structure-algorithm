package main

import (
	"fmt"
	"sort"
)

type Job struct {
	start  int
	end    int
	profit int
}

func jobScheduling(startTime, endTime, profit []int) int {
	n := len(startTime)
	jobs := make([]Job, n)

	for i := 0; i < n; i++ {
		jobs[i] = Job{startTime[i], endTime[i], profit[i]}
	}

	// Urutkan berdasarkan endTime
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	// dp[i] menyimpan maksimum profit hingga job ke-i
	dp := make([]int, n)
	dp[0] = jobs[0].profit

	for i := 1; i < n; i++ {
		// Inisialisasi dengan ambil profit sekarang
		inclProfit := jobs[i].profit

		// Cari job terakhir yang tidak bentrok (pakai binary search manual)
		l := binarySearch(jobs, i)
		if l != -1 {
			inclProfit += dp[l]
		}

		// Bandingkan dengan tidak mengambil job i
		dp[i] = max(dp[i-1], inclProfit)
	}

	return dp[n-1]
}

// Binary search untuk cari job terakhir yang tidak bentrok
func binarySearch(jobs []Job, index int) int {
	lo, hi := 0, index-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if jobs[mid].end <= jobs[index].start {
			if jobs[mid+1].end <= jobs[index].start {
				lo = mid + 1
			} else {
				return mid
			}
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	startTime := []int{1, 2, 3, 4,  6,  7,  9 }
    endTime   := []int{3, 5, 6, 7,  9,  10, 11}
    profit  := []int{5, 6, 5, 11, 12, 14, 20}

	result := jobScheduling(startTime, endTime, profit)
	fmt.Println(result) // Output: 150
}