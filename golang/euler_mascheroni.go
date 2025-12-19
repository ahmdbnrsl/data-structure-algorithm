package main 

import "fmt"
import "math"

func main() {
    eulerMascheroni := math.Log(1e9) + 1/(2*1e9) - 1/(12*1e9*1e9) - math.Log(1e9)
    fmt.Println(eulerMascheroni, 1e+100000)
}