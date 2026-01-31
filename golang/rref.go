package main 

import (
    "fmt"
)

type RREF struct {
    AugmentedMatrix [][]float64
}

func countZeroPrefix(row []float64)int {
    var count int
    for _, e := range row {
        if e == 0 {
            count += 1
        } else {
            return count
        }
    }
    return count
}

func (rref *RREF) NewMatrix(matrix [][]float64) {
    var initialLength int = len(matrix[0])
    for _, e := range matrix {
        if len(e) != initialLength {
            return
        }
    }
    rref.AugmentedMatrix = matrix
}

func (rref *RREF) GenerateRREF() {
    // Row Interchange
    var temp [][]float64
    var currentZero int
    for i, e := range rref.AugmentedMatrix {
        if i > 0 {
            var countZero int = countZeroPrefix(e)
            if countZero >= currentZero {
                temp = append(temp, e)
                currentZero = countZero
            } else {
                temp = append([][]float64{e}, temp...)
                currentZero = countZero
            }
        } else {
            temp = append(temp, e)
            currentZero = countZeroPrefix(e)
        }
    }
    
    rref.AugmentedMatrix = temp
    
    
    // RREFMatrix
    for i := 0; i < len(rref.AugmentedMatrix); i++ {
        for j := i; j < len(rref.AugmentedMatrix); j++ {
            if i > 0 {
                if rref.AugmentedMatrix[j][i-1] != 0 {
                    var newRow []float64
                    var multiplicate float64
                    
                    if rref.AugmentedMatrix[i-1][i-1] != 0 {
                        multiplicate = -(rref.AugmentedMatrix[j][i-1]/rref.AugmentedMatrix[i-1][i-1])
                        
                    }
                    for idx, z := range rref.AugmentedMatrix[i-1] {
                        newRow = append(newRow, rref.AugmentedMatrix[j][idx] + multiplicate * z)
                        
                    }
                    rref.AugmentedMatrix[j] = newRow
                }
            }
            
            if j == i && rref.AugmentedMatrix[j][i] != 1 && rref.AugmentedMatrix[j][i] != 0{
                multiplicate := 1 / rref.AugmentedMatrix[j][i]
                var newRow []float64
                for _, k := range rref.AugmentedMatrix[j] {
                    newRow = append(newRow, k * multiplicate)
                }
                rref.AugmentedMatrix[j] = newRow
            }
        }
    }
    
    // Clean Above Leading One
    for i := len(rref.AugmentedMatrix) - 2; i >= 0; i-- {
        for j := i; j >= 0; j-- {
            if rref.AugmentedMatrix[j][i+1] != 0 && rref.AugmentedMatrix[i+1][i+1] != 0 {
                multiplicate := -(rref.AugmentedMatrix[j][i+1]/rref.AugmentedMatrix[i+1][i+1])
                
                var newRow []float64
                for idx, z := range rref.AugmentedMatrix[j] {
                    newRow = append(newRow, z + multiplicate * rref.AugmentedMatrix[i+1][idx])
                }
                
                rref.AugmentedMatrix[j] = newRow
            }
        }
    }
}

func main() {
    var rref RREF
    // Unique Solution
    rref.NewMatrix([][]float64{
        []float64{1, 1, 1, 6},
        []float64{1, -1, 1, 2},
        []float64{1, 1, -1, 0},
    })
    rref.GenerateRREF()
    fmt.Println(rref.AugmentedMatrix)
    
    // Infinity Solution (Free Variable)
    rref.NewMatrix([][]float64{
        []float64{1, 2, 1, 4},
        []float64{1, 1, -1, 1},
        []float64{2, 4, 2, 8},
    })
    rref.GenerateRREF()
    fmt.Println(rref.AugmentedMatrix)
    
    // No Solution
    rref.NewMatrix([][]float64{
        []float64{1, 1, 1, 3},
        []float64{3, -1, 2, 4},
        []float64{2, 2, 2, 7},
    })
    rref.GenerateRREF()
    fmt.Println(rref.AugmentedMatrix)
    
    // Invers matriks
    rref.NewMatrix([][]float64{
        []float64{1, 2, 1, 0},
        []float64{0, 1, 0, 1},
    })
    rref.GenerateRREF()
    fmt.Println(rref.AugmentedMatrix)
    
    rref.NewMatrix([][]float64{
        []float64{1, -1, -1, 0, 0},
        []float64{1, 1, 0, -1, 2},
        []float64{1, -1, 1, 0, 0},
        []float64{1, 1, 0, 1, 0},
    })
    rref.GenerateRREF()
    fmt.Println(rref.AugmentedMatrix)
}