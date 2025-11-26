package main 

import "fmt"

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
    
    for currentNumber < int64(1e9) / currentNumber {
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

func uniqueFactor(dictPrime []int64, number int64) []int64{
    var result []int64
    
    currentIdx := 0
    tempPrime := dictPrime[currentIdx]
    
    for number > 1 {
        if number % tempPrime == 0 {
            number = number / tempPrime
            result = append(result, tempPrime)
        } else {
            currentIdx += 1
            tempPrime = dictPrime[currentIdx]
        }
    }
    
    return result
}

func main() {
    dictPrime := generateDictPrime()
    factor := uniqueFactor(dictPrime, int64(1e9))
    fmt.Println(factor)
}