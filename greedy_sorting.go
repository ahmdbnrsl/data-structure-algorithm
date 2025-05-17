package main 

import "fmt"

func sort_ballons(ballons [][2]int) {
    for i := 0; i < len(ballons) - 1; i++ {
        for j := 0; j < len(ballons) - 1; j++ {
            if ballons[j][1] > ballons[j + 1][1] {
                receptacle := ballons[j]
                ballons[j] = ballons[j + 1]
                ballons[j + 1] = receptacle
            }
        }
    }
}

func greedy_sorting(ballons [][2]int)int {
    sort_ballons(ballons)
    arrowPosition := 1
    temp := 0
    
    for i, el := range ballons {
        if i == 0 { temp = el[1] } else if i > 0 {
            if el[0] > temp {
                arrowPosition += 1
                temp = el[1]
            }
        }
    }
    
    return arrowPosition
}

func main() {
    params := [][2]int{
        [2]int{1, 2},
        [2]int{3, 4},
        [2]int{5, 6},
        [2]int{7, 8},
    }
    result := greedy_sorting(params)
    fmt.Println(result)
}