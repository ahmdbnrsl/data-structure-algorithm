package main

import "fmt"
import "math"

func maxSubArray(array []float64) float64{
    res := array[0]
    maxEnding := array[0]
    
    for i := 1; i < len(array); i++ {
        maxEnding = math.Max(maxEnding + array[i], array[i])
        res = math.Max(res, maxEnding)
    }
    
    return res
}

func main() {
    fmt.Println(maxSubArray([]float64{1, -2, -2, 3}))
}