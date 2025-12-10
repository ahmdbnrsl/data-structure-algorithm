package main 

import (
    "math"
    "fmt"
)

func gammaApprox(z float64) float64 {
    if z < 0 {
        return -1
    } else if z - math.Trunc(z) == 0 {
        power := z - 1
        multiplicate := 1.0
        
        exponentialSign := -1.0
        multiplicateSign := 1.0
        
        var result float64
        
        for power >= 0 {
            result += (multiplicate * multiplicateSign) * math.Pow(0, power) * (math.Exp(-0) * exponentialSign)
            
            multiplicate *= power
            power -= 1
            
            exponentialSign = -exponentialSign
            multiplicateSign = - multiplicateSign
        }
        
        fmt.Println(-result)
        return -result
    } else {
        var frontCalc float64 = 1
        
        for z > 1 {
            z -= 1
            frontCalc *= z
        }
        
        var g float64 = 5
        var Ag float64 = 1.000000000190015
        
        c := []float64{
            76.18009172947146,
            -86.50532032941677,
            24.01409824083091,
            -1.231739572450155,
            0.001208650973866179,
            -0.000005395239384953,
        }
        
        for i, e := range c {
            Ag += e / (z + float64(i) + 1)
        }
        
        result := frontCalc * (math.Sqrt(2*math.Pi)/z) * math.Pow(z + g + 0.5, z + 0.5) * math.Exp(-(z + g + 0.5)) * Ag
        fmt.Println(result)
        return result
    }
}

func main() {
    gammaApprox(0.3)
}