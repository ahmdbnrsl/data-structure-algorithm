package main 

import (
    "fmt"
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

func binaryToDecimal(biner []int64) int64 {
    dictBase := generateDictBase(1e3)
    var result int64
    var lastIndex int = len(biner) - 1
    
    for i, e := range biner {
        if e == 1 {
            result += dictBase[lastIndex - i]
        }
    }
    
    return result
}

func binaryPow(A, B, M int64) int64 {
    var result int64 = 1
    A = A % M
    
    for B > 0 {
        if B % 2 == 1 {
            result = (result * A) % M
        }
        
        B = B / 2
        A = (A * A) % M
    }
    
    return result
}

func inversBinary(biner []int64) []int64{
    var result []int64
    for _, e := range biner {
        result = append(result, binaryPow(2, e, 1e9+7) - 2 * e)
    }
    return result
}

func additionBinary(biner []int64) {
    var additive int64 = 1
    var lastIndex int = len(biner) - 1
    
    for i, _ := range biner {
        if additive == 0 {
            break
        }
        
        e := biner[lastIndex - i]
        biner[lastIndex - i] = binaryPow(2, e, 1e9+7) - 2 * e
        if biner[lastIndex - i] == 1 {
            additive = 0
        }
    }
}

func LSB(biner []int64) []int64 {
    invBiner := inversBinary(biner)
    additionBinary(invBiner)
    var result []int64
    for i, e := range biner {
        result = append(result, (invBiner[i] + e) / 2)
    }
    return result
}

func query(tree []int64, x int64) int64 {
    var indices []int64 = []int64{x}
    for x > 0 {
        lsb := binaryToDecimal(LSB(decimalToBinary(x)))
        indices = append(indices, x - lsb)
        x -= lsb
    }
    
    var result int64 = 0
    for _, e := range indices {
        result += tree[e]
    }
    
    return result
}

func update(tree []int64, x int64) {
    for x < int64(len(tree)) {
        tree[x] += 1
        
        lsb := binaryToDecimal(
            LSB(decimalToBinary(x)),
        )
        x = x + lsb
    }
}

func BIT(A []int64, max int64) int64 {
    tree := make([]int64, max)
    var total int64
    for i, e := range A {
        newInversion := int64(i) - query(tree, e)
        total += newInversion
        update(tree, e)
    }
    return total
}

func main() {
    
    start := time.Now()
    
    biner := decimalToBinary(3)
    lsb := LSB(biner)
    decimal := binaryToDecimal(lsb)
    
    fmt.Println(biner)
    fmt.Println(decimal)
    
    A := []int64{6, 1, 8, 2, 5, 4, 7, 3}
    var max int64 = 8
    bit := BIT(A, max+1)
    
    fmt.Println(bit)
    
    elapsed := time.Since(start)
    fmt.Printf(" %.6f second\n", elapsed.Seconds())
    
    fmt.Println(decimalToBinary(73), decimalToBinary(110))
}