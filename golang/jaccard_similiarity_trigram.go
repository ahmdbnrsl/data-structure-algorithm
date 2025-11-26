package main 

import (
    "fmt"
    "slices"
)

func makeTrigram(str string) []string {
    var trigrams []string
    for i := 0; i < len(str) - 2; i++ {
        trigrams = append(trigrams, str[i:i+3])
    }
    
    return trigrams
}

func unique(slice []string) []string {
    var result []string
    for _, e := range slice {
        if !slices.Contains(result, e) {
            result = append(result, e)
        }
    }
    return result
}

func jaccardSimiliarityTrigram(str1, str2 string) float64  {
    trigram_1, trigram_2 := unique(makeTrigram(str1)), unique(makeTrigram(str2))
    
    intersection := 0
    
    for _, e := range trigram_2 {
        if slices.Contains(trigram_1, e) {
            intersection += 1
        }
    }
    
    unionSet := append(trigram_1, trigram_2...)
    union := unique(unionSet)
    
    return float64(intersection) / float64(len(union))
}

func main() {
    fmt.Println(jaccardSimiliarityTrigram("machinelearning", "mathematics"))
}