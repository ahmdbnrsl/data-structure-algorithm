package main 

import "fmt"

func increaseArray(slice []int)int {
    countOfOperations := 0
    for i, e := range slice {
        if i > 0 && e <= slice[i - 1] {
            diff := slice[i - 1] - e + 1
            slice[i] += diff
            countOfOperations += diff
        }
    }
    return countOfOperations
}

func main() {
    fmt.Println(increaseArray([]int{3,2,1}))
}