package main 

import "fmt"
import "time"

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

type PSG struct{
    A, P, E []int64
}

func(psg *PSG) addNewArray(array []int64) {
    var M int64 = 1e9+7
    var inv2 int64 = binaryPow(2, M-2, M)
    
    psg.A = append([]int64{}, array...)
    psg.P = []int64{}
    
    var currentP int64 = 0
    
    for i, e := range psg.A {
        psg.P = append(psg.P, ((e % M)  * (binaryPow(2, int64(i), M) % M)  % M) + currentP)
        currentP = psg.P[i]
        
        if i == 0 && len(psg.E) == 0 {
            psg.E = []int64{((2 % M) * (inv2 % M)) % M}
            continue
        } else if i >= len(psg.E) {
            psg.E = append(psg.E, ((psg.E[i - 1] % M) * (inv2 % M) % M))
        }
    }
}

func(psg *PSG) prefixSum(l, r int64) int64 {
    var M int64 = 1e9+7
    var left int64 = l
    if l != 0 {
        left = psg.P[l - 1]
    }
    
    return (((psg.P[r] - left + M) % M) * (psg.E[l] % M)) % M
}

func main() {
    psg := PSG{}
    psg.addNewArray([]int64{10, 20, 30, 40, 50})
    
    start := time.Now()
    
    prefsum := psg.prefixSum(1, 3)
    fmt.Println(prefsum)
    prefsum = psg.prefixSum(0, 4)
    fmt.Println(prefsum)
    prefsum = psg.prefixSum(2, 4)
    fmt.Println(prefsum)
    
    elapsed := time.Since(start)
    fmt.Printf(" %.6f second\n", elapsed.Seconds())
}