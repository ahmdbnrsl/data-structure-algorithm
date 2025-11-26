package main 

import "fmt"
import "time"

func checkPrime(number int64) bool{
    var i int64 = 2
    for i <= number/i {
        if number % i == 0 {
            return false
        }
        i++
    }
    return true
}

func generateDictPrime()[]int64 {
    var result []int64 = []int64{2}
    
    var i int64 = 0
    var currentNumber int64 = 2
    
    for currentNumber <= int64(1e9) / currentNumber {
        currentNumber += 1
        isPrime := checkPrime(currentNumber)
        for !isPrime {
            currentNumber += 1
            isPrime = checkPrime(currentNumber)
        }
        
        result = append(result, int64(currentNumber))
        i++
    }
    
    return result
}

func countSliceElements(slice []int64) map[int64]int {
	counts := make(map[int64]int)
	for _, item := range slice {
		counts[item]++
	}
	return counts
}

func uniqueFactor(dictPrime []int64, number int64) ([]int64, map[int64]int) {
    var factor []int64
    lenDict := len(dictPrime)
    
    currentIdx := 0
    tempPrime := dictPrime[currentIdx]
    
    for number > 1 {
        if number % tempPrime == 0 {
            number = number / tempPrime
            factor = append(factor, tempPrime)
        } else {
            currentIdx += 1
            if currentIdx >= lenDict {
                factor = append(factor, number)
                break
            }
            tempPrime = dictPrime[currentIdx]
        }
    }
    
    return factor, countSliceElements(factor)
}

func GCD(a, b int64)int64 {
    if a < b {
        a, b = b, a
    }
    
    mod := a % b
    if mod == 0 {
        return b
    }
    
    return GCD(b, mod)
}

func GCD_LCM(dictPrime []int64, numbers ...int64)(int64, int64) {
    var M int64 = 1e9+7
    var gcd int64 = numbers[0]
    var factor map[int64]int = make(map[int64]int)
    
    for i, e := range numbers {
        if i > 0 {
            gcd = GCD(gcd, e)
        }
        _, count := uniqueFactor(dictPrime, e)
        for i, f := range count {
            if factor[i] < f {
                factor[i] = f
            }
        }
    }
    
    var lcm int64 = 1
    for i, e := range factor {
        var temp int64 = i
        for j := 1; j < e; j++ {
            temp = ((temp % M) * (i % M)) % M
        }
        lcm = ((lcm % M) * (temp % M)) % M
    }
    return gcd, lcm
}

func main() {
    dictPrime := generateDictPrime()
    start := time.Now()
    gcd, lcm := GCD_LCM(dictPrime, 5, 999999937, 999999929, 1000000000, 268435456, 244140625)
    fmt.Println(gcd, lcm)
    elapsed := time.Since(start)
    fmt.Printf("myFunc took %.6f second\n", elapsed.Seconds())
}