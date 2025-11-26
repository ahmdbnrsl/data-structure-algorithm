package main

import (
    "fmt"
    "math"
    "sort"
)

type Jobs struct {
    Start, End, Profits float64
}

func jobScheduling(start, end, profits []float64) float64{
    
    var jobs []Jobs
    for i := 0; i < len(end); i++ {
        jobs = append(jobs, Jobs{
            Start: start[i],
            End: end[i],
            Profits: profits[i],
        })
    }
    
    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i].End < jobs[j].End
    })
    
    var prevEnd       float64 = jobs[0].End
    var currentEnd    float64 = jobs[0].End
    var maxProfit     float64 = jobs[0].Profits
    var currentProfit float64 = jobs[0].Profits
    
    for i := 1; i < len(jobs); i++{
        if currentEnd > jobs[i].Start {
            if (jobs[i].Start >= prevEnd) {
                maxProfit = math.Max(maxProfit, jobs[i].Profits + currentProfit)
                currentEnd = jobs[i].End
            } else {
                maxProfit = math.Max(maxProfit, jobs[i].Profits)
            }
        } else {
            maxProfit    += jobs[i].Profits
            currentEnd    = jobs[i].End
            if i > 0 {
                prevEnd       = jobs[i - 2].End
                currentProfit = jobs[i - 2].Profits
            }
        }
    }

    return maxProfit
    
}

func main() {
    startTime := []float64{1, 2, 3, 4,  6,  7,  9 }
    endTime   := []float64{3, 5, 6, 7,  9,  10, 11}
    profits   := []float64{5, 6, 5, 11, 12, 14, 20}
    
    js        := jobScheduling(startTime, endTime, profits)
    fmt.Println(js)
    /*
        expected result : 37
    */
}