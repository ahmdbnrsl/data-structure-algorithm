package main 

import "fmt"
import "maps"

func findMaxValue(m map[int]int) int {
    if len(m) == 0 {
        return 0
    }

    maxValue := 0
    first := true

    for _, value := range m {
        if first {
            maxValue = value
            first = false
            continue
        }
        if value > maxValue {
            maxValue = value
        }
    }
    return maxValue
}

func closestSubsetSum(A []int, T int)int {
    possible_sum := map[int]int{0:0}
    
    for _, num := range A {
        var next_sums map[int]int = make(map[int]int)
        for _, s := range possible_sum {
            if s + num <= T {
                next_sums[s + num] = s + num
            }
        }
        maps.Copy(possible_sum, next_sums)
    }
    
    return findMaxValue(possible_sum)
}

func main() {
    fmt.Println(closestSubsetSum([]int{3, 5, 8}, 10))
}