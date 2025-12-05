package main 

import (
    "fmt"
    _"slices"
    "time"
)

func generateDictBase(length int64) []int64 {
    var base []int64 = []int64{1}
    var M int64 = 1e9+7
    var i int64 = 1
    
    for i <= length {
        base = append(
            base,
            ((base[len(base) - 1] % M) * (2 % M)) % M,
        )
        
        i++
    }
    
    return base
}

func decimalToBinary(decimal int64) []int64 {
    var bin []int64
    for decimal > 0 {
        bin = append([]int64{decimal % 2}, bin...)
        decimal = decimal / 2
    }
    
    return bin
}

func binaryToDecimal(biner, dictBase []int64) int64 {
    var result int64
    var lastIndex int = len(biner) - 1
    
    for i, e := range biner {
        if e == 1 {
            result += dictBase[lastIndex - i]
        }
    }
    
    return result
}

func main() {
    dictBase := generateDictBase(1e3)
    
    start := time.Now()
    
    biner := decimalToBinary(7345)
    decimal := binaryToDecimal(biner, dictBase)
    
    fmt.Println(biner)
    fmt.Println(decimal)
    
    elapsed := time.Since(start)
    fmt.Printf(" %.6f second\n", elapsed.Seconds())
}