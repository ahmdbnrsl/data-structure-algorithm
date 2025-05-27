package main 

import "fmt"

func ISR(num int)int {
    for i := num; i > 0; i-- {
        for j := 0; j < num; j++ {
            if i * j == num && i == j {
                return i
            }
        }
    }
    return 0
}

func estimateSquareRoot(num int) float64 {
    n := ISR(num)
    if n != 0 {
        return float64(n)
    }
    count_down := 0
    count_up := 0
    for i := num; i > 0; i-- {
        if ISR(i) != 0 {
            count_down = i
            break
        }
    }
    temp := num + 1
    for count_up == 0 {
        if ISR(temp) != 0 {
            count_up = temp
            break
        }
        temp++
    }
    if count_up - num < num - count_down {
        m := float64(ISR(count_up))
        return m - (float64(count_up - num) / (m * 2.0))
    } else {
        m := float64(ISR(count_down))
        return m + (float64(num - count_down) / (m * 2.0))
    }
}

func main() {
    fmt.Println(estimateSquareRoot(18))
}