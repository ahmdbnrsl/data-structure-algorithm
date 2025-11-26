package main 

import "fmt"

func main() {
    slice := []int{2,7,3,1,9,0,7,2,6,8,5,5,6,9,8}
    // fmt.Println(slice)
    // Sort(slice, "max")
    // fmt.Println(slice)
    // Sort(slice, "min")
    // fmt.Println(slice)
    slice_2 := []int{3,2,9,0,9,0,9,2,6,8,6,1,2,9,8}
    mergeSortedSlices, unpackSortedSlices := SortMultiple(slice, slice_2)
    fmt.Println(mergeSortedSlices)
    for _, i := range unpackSortedSlices {
        fmt.Println(i)
    }
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

func SortMultiple(slices ...[]int) ([]int, [][]int) {
    var tempSlices []int
    var length []int
    var packSlices [][]int
    for i := 0; i < len(slices); i++ {
        tempSlices = append(tempSlices, slices[i]...)
        length = append(length, len(slices[i]))
    }
    for i := 0; i < len(tempSlices) - 1; i++ {
        for j := 0; j < len(tempSlices) - 1; j++ {
            if tempSlices[j] > tempSlices[j + 1] {
                receptacle := tempSlices[j]
                tempSlices[j] = tempSlices[j + 1]
                tempSlices[j + 1] = receptacle
            }
        }
    }
    var tempIndex int
    for index, i := range length {
        switch index {
        case 0:
            packSlices = append(packSlices, tempSlices[:i + tempIndex])
            break
        case len(length) - 1:
            packSlices = append(packSlices, tempSlices[index + tempIndex - 1:])
            break
        default:
            packSlices = append(packSlices, tempSlices[index + tempIndex - 1:i + tempIndex])
            break
        }
        tempIndex = i
    }
    return tempSlices, packSlices
}