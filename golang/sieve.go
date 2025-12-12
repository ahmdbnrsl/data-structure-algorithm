package main 

import (
    "fmt"
    "time"
)

func generatePrimeDict(n int64) []int64 {
    array := make([]int64, n+1)
    
    for i, _ := range array {
        if i > 1 {
            array[i] = 1
        }
        
    }
    
    for i, e := range array {
        if i < 2 || e == 0 {
            continue
        }
        var j int64 = int64(i * i)
        for j <= n {
            array[j] = 0
            j += int64(i) 
        }
    }
    
    return array
}

func prefixSum(array []int64) []int64 {
    result := make([]int64, len(array))
    for i, _ := range result {
        if i > 1 {
            result[i] = result[i-1] + array[i]
        }
    }
    return result
}

func main() {
    fmt.Println("Welcome to Sieve of Eratostenes")
    isPrime := generatePrimeDict(1e7)
    aps := prefixSum(isPrime)
    
    start := time.Now()
    fmt.Println(aps[1e3])
    elapsed := time.Since(start)
    fmt.Printf(" %.6f second\n", elapsed.Seconds())
    
    start = time.Now()
    fmt.Println(aps[1e7])
    elapsed = time.Since(start)
    fmt.Printf(" %.6f second\n", elapsed.Seconds())
    
    start = time.Now()
    fmt.Println(aps[1e6])
    elapsed = time.Since(start)
    fmt.Printf(" %.6f second\n", elapsed.Seconds())
}