package main 

import "fmt"

func main() {
    slice := []int{2,7,3,1,9,0,7,2,6,8,5,5,6,9,8}
    fmt.Println(slice)
    Sort(slice, "max")
    fmt.Println(slice)
    Sort(slice, "min")
    fmt.Println(slice)
}

func Sort(slice []int, from string) {
    for i := 0; i < len(slice) - 1; i++ {
        for j := 0; j < len(slice) - 1; j++ {
            if from == "min" {
                if slice[j] > slice[j + 1] {
                    receptacle := slice[j]
                    slice[j] = slice[j + 1]
                    slice[j + 1] = receptacle
                }
            } else if from == "max" {
                if slice[j] < slice[j + 1] {
                    receptacle := slice[j]
                    slice[j] = slice[j + 1]
                    slice[j + 1] = receptacle
                }
            }
        }
    }

}