package main 

import "fmt"
import "strings"
import "slices"

func findIndex(slice []string, target string)int {
    for i, item := range slice {
        if target == item {
            return i
        }
    }
    return -1
}

func bracketMatcher(params string) bool {
    
    splitedBracket := strings.Split(params, "")
    copySplitedBracket := make([]string, len(splitedBracket))
    
    openingBracket := []string{"{", "[", "("}
    closingBracket := []string{"}", "]", ")"}
    
    copy(copySplitedBracket, splitedBracket)
    
    if len(splitedBracket) % 2 != 0 {
        fmt.Println("0")
        return false
    }
    
    mid := len(splitedBracket) / 2
    splited1 := copySplitedBracket[:mid]
    splited2 := copySplitedBracket[mid:]
    
    for i, s := range splited1 {
        if(slices.Contains(closingBracket, s)) {
            break
        }
        indexBracket := findIndex(openingBracket, s)
        splited1[i] = closingBracket[indexBracket]
    }
    slices.Reverse(splited1)
    if slices.Equal(splited1, splited2) {
        return true
    }
    
    for index, item := range splitedBracket {
        if index == 0 {
            if slices.Contains(closingBracket, item) {
                return false
            }
            indexBracket := findIndex(openingBracket, item)
            if !slices.Contains(openingBracket, splitedBracket[index + 1]) && splitedBracket[index + 1] != closingBracket[indexBracket] {
                return false
            }
        } else {
            if len(splitedBracket) > 2 {
                if index == 1 {
                    if !slices.Contains(openingBracket, splitedBracket[index + 1]) {
                        return false
                    }
                } else if index > 1 && index < len(splitedBracket) - 1 {
                    if index % 2 == 0 {
                        if !slices.Contains(closingBracket, splitedBracket[index + 1]) || !slices.Contains(closingBracket, splitedBracket[index - 1]) {
                            return false
                        }
                        indexBracket := findIndex(openingBracket, item)
                        if splitedBracket[index + 1] != closingBracket[indexBracket] {
                            return false
                        }
                    } else {
                        if !slices.Contains(openingBracket, splitedBracket[index - 1]) || !slices.Contains(openingBracket, splitedBracket[index + 1]) {
                            return false
                        }
                        indexBracket := findIndex(closingBracket, item)
                        if splitedBracket[index - 1] != openingBracket[indexBracket] {
                            return false
                        }
                    }
                } else {
                    if !slices.Contains(openingBracket, splitedBracket[index - 1]) {
                        return false
                    }
                    indexBracket := findIndex(closingBracket, item)
                    if splitedBracket[index - 1] != openingBracket[indexBracket] {
                        return false
                    }
                }
            }
        }
    }
    return true
}

func main() {
    fmt.Println(bracketMatcher("{}[]()()")) // true
    fmt.Println(bracketMatcher("(({[{[()]}]}))")) // true
    fmt.Println(bracketMatcher("{{}}[]()()")) // false
    fmt.Println(bracketMatcher("{({([}])])")) // false
    fmt.Println(bracketMatcher(")()((()(")) // false
}