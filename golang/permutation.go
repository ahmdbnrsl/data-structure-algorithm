package main 

import (
    "fmt"
    "math"
)

type Permutation struct {
    matrix map[int64]int64
}

func isContainItem(slice []int64, item int64) bool {
    for _, e := range slice {
        if e == item {
            return true
        }
    }
    return false
}

func(p *Permutation) createMatrix(matrix map[int64]int64) {
    var out []int64
    for _, e := range matrix {
        out = append(out, e)
    }
    
    for i, _ := range matrix {
        if !isContainItem(out, i) {
            fmt.Println("Input and Output Must contain same element")
            return
        }
    }
    
    p.matrix = matrix
}

func(p *Permutation) Sign() string {
    if len(p.matrix) == 0 {
        fmt.Println("please add permutations first")
        return "please add permutations first"
    } else {
        var cycles [][]int64
        var currentQuery []int64
        
        for e, _ := range p.matrix {
            if isContainItem(currentQuery, e) {
                continue
            }
            
            var cycle []int64
            var current int64 = 0
            
            for current != e {
                if current == 0 {
                    current = e
                }
                
                currentQuery = append(currentQuery, current)
                cycle = append(cycle, current)
                current = p.matrix[current]
            }
            
            cycles = append(cycles, cycle)
        }
        
        var k int64
        for _, e := range cycles {
            if len(e) <= 2 {
                k += 1
                continue
            } else {
                k += int64(len(e) - 1)
            }
        }
        
        if math.Pow(-1, float64(k)) < 1 {
            return "Odd"
        } else {
            return "Even"
        }
    }
}

func main() {
    fmt.Println("Welcome To Permutation")
    var permutation Permutation
    permutation.createMatrix(map[int64]int64{
        1:3,
        2:4,
        3:1,
        4:5,
        5:6,
        6:2,
    })
    fmt.Println(permutation.Sign())
}