package main 

import "fmt"

func Contains(slice []int, item int) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

func twoSum(arr []int, target int) []int {
    var result []int
    
    for index, el := range arr {
        for idx, e := range arr {
            if el + e == target && index != idx && !Contains(result, index) {
                result = append(result, index)
            }
        }
    }
    
    return result
}

func main() {
    array := []int{
        1, 5, 2, 9, 4, 8, 5,
    }
    
    fmt.Println("Index: ", twoSum(array, 6))
}