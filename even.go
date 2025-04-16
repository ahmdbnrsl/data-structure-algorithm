package main 

import "fmt"

func main() {
    var even []int
    for i := 1; i <= 100; i++ {
        if i % 2 == 0 && (i % 6 == 0 || i % 9 == 0) {
            even = append(even, i)
        }
    }
    
    fmt.Println(len(even))
}
