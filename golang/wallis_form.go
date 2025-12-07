package main 

import (
    "fmt"
    "math"
)

func calcNumerator(m float64) float64 {
    if m == 0 {
        return 1
    }
    
    var nominator float64 = 1
    m -= 1
    for m > 0 {
        nominator *= m
        m -= 2
    }
    return nominator
}

func isEven(number float64) bool {
    if number == 0 {
        return true
    }
    return int64(number) % 2 == 0
}

func wallis(m, n float64) float64 {
    
    var denominator float64 = 1
    var reduc float64 = 0
    var K float64 = 1
    
    for m + n - reduc > 0 {
        denominator *= m + n - reduc
        reduc += 2
    }
    
    var numerator float64 = calcNumerator(m) * calcNumerator(n)
    fmt.Println(numerator, "/", denominator, "× (π/2)")
    
    if isEven(m) && isEven(n) {
        K = math.Pi / 2
    }
    return (numerator / denominator) * K
}

func main() {
    fmt.Println("Welcome To Wallis' Form")
    isWallis := wallis(4, 4)
    fmt.Println(isWallis)
}