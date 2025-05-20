package main 

import "fmt"

func prissonCellsAfterNDays(cells []int, countOfDays int) {
    for i := 0; i < countOfDays; i++ {
        var left int
        for index, el := range cells {
            if index == 0 || index == len(cells) - 1 {
                cells[index] = 0
            } else if index > 0 && index < len(cells) - 1 {
                if(left == cells[index + 1]) {
                    cells[index] = 1
                } else {
                    cells[index] = 0
                }
            }
            left = el
        }
    }
}

func main() {
    cells := []int{1, 1, 0, 0, 1, 0, 0, 1}
    N := 1
    fmt.Println(cells)
    prissonCellsAfterNDays(cells, N)
    fmt.Println(cells)
    
}