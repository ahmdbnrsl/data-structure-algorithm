package main 

import "fmt"
import "strings"
import "slices"

func split(str, splitBy string) []string {
    result := []string{}
    if splitBy == "" {
        for i := 0; i < len(str); i++ {
            result = append(result, string(str[i]))
        }
    } else {
        var index int
        for i := 0; i < len(str); i++ {
            if len(result) == 0 && splitBy != string(str[i]) {
                result = append(result, string(str[i]))
                index = 0
                continue
            } else if len(result) > 0 && i < len(str) - 1 {
                if splitBy == string(str[i]) {
                    result = append(result, "")
                    index += 1
                    continue
                }
                result[index] = result[index] + string(str[i])
                continue
            }
        }
    }
    return result
}

func validate(str, target string, listLetter []string, countIndex []int) bool {
    for _, el := range split(target, "") {
        if !strings.Contains(str, el) {
            return false
        }
    }
    
    duplicate := []string{}
    isIndex := []int{}
    
    for _, el := range split(str, "") {
        if strings.Contains(strings.Join(listLetter, ""), el) {
            if countIndex[findIndex[string](listLetter, el)] > 1 {
                indexDuplicate := findIndex[string](duplicate, el)
                if indexDuplicate != -1 {
                    isIndex[indexDuplicate] += 1
                    continue
                }
                duplicate = append(duplicate, el)
                isIndex = append(isIndex, 1)
            }
        }
    }
    
    listTrue := []bool{}
    listTrue2 := []bool{}
    
    for _, el := range duplicate {
        if countIndex[findIndex[string](listLetter, el)] > 1 {
            newIndex := findIndex[string](duplicate, el)
            validateIndex := findIndex[string](listLetter, el)
        
            if isIndex[newIndex] == countIndex[validateIndex] {
                listTrue = append(listTrue, true)
            }
        }
    }
    
    for _, el := range countIndex {
        if el > 1 {
            listTrue2 = append(listTrue2, true)
        }
    }
    
    if len(listTrue2) == len(listTrue) {
        return true
    }
    
    return false
}

func contains(str, target string) bool {
    for _, el := range split(target, "") {
        if !strings.Contains(str, el) {
            return false
        }
    }
    return true
}

func findIndex[T int | string](slice []T, target any) int {
    for i, el := range slice {
        if el == target {
            return i
        }
    }
    return -1
}

func minimum_window_substring(str, target string) string {
    if !contains(str, target) {
        return "empty"
    }
    
    temp := []int{}
    for i, el := range split(str, "") {
        if strings.Contains(target, el) {
            temp = append(temp, i)
        }
    }
    
    listLetter := []string{}
    countIndex := []int{}
    
    for _, el := range split(target, "") {
        indexDuplicate := findIndex[string](listLetter, el)
        if indexDuplicate != -1 {
            countIndex[indexDuplicate] += 1
            continue
        }
        listLetter = append(listLetter, el)
        countIndex = append(countIndex, 1)
    }
    
    listIndex := []int{}
    listStr := []string{}
    
    for _, el := range temp {
        isString := ""
        tempTarget := ""
        for i, e := range split(str, "") {
            if i >= el {
                if len(tempTarget) < len(target) {
                    isString = isString + e
                }
                if strings.Contains(target, e) {
                    if !strings.Contains(tempTarget, e) {
                        tempTarget = tempTarget + e
                    }
                }
            }
        }
        
        if validate(isString, target, listLetter, countIndex) {
            listStr = append(listStr, isString)
            listIndex = append(listIndex, len(isString))
        }
    }
    
    if len(listIndex) > 0 {
        indexResult := findIndex[int](listIndex, slices.Min(listIndex))
        return listStr[indexResult]
    }
    return ""
}

func main() {
    str := "AAABBC"
    target := "AABC"
    fmt.Println(minimum_window_substring("ADOBECODEBANCAC", "ABC"))
    fmt.Println(minimum_window_substring(str, target))
}