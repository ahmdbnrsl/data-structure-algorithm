package main 

import (
    "fmt"
    "errors"
    "time"
    "strconv"
)

type Matrix2x2 struct {
    L, R, T, B []int64
}

func(m *Matrix2x2) new2x2Matrix(arr [4]int64) {
    m.L = []int64{arr[0], arr[2]}
    m.R = []int64{arr[1], arr[3]}
    m.T = []int64{arr[0], arr[1]}
    m.B = []int64{arr[2], arr[3]}
}

func subDotProduct(row, col []int64) int64 {
    if len(row) != 2 || len(col) != 2 {
        errors.New("Invalid row and column shape")
    }
    var M int64 = 1e9+7
    a := ((row[0] % M) * (col[0] % M)) % M
    b := ((row[1] % M) * (col[1] % M)) % M
    return ((a % M) + (b % M)) % M
}

func dotProduct(m, n Matrix2x2) [4]int64 {
    sdp := subDotProduct
    return [4]int64{
        sdp(m.T, n.L),
        sdp(m.T, n.R),
        sdp(m.B, n.L),
        sdp(m.B, n.R),
    }
}

func(m *Matrix2x2) binaryPow(B int64) {
    result := Matrix2x2{}
    result.new2x2Matrix([4]int64{
        1, 0, 0, 1,
    })
    for B > 0 {
        if B % 2 == 1 {
            result.new2x2Matrix(dotProduct(result, *m))
        }
        m.new2x2Matrix(dotProduct(*m, *m))
        B = B / 2
    }
    m.new2x2Matrix([4]int64{
        result.T[0], result.T[1],
        result.B[0], result.B[1],
    })
}

func(m *Matrix2x2) Print() {
    str := "|" + strconv.FormatInt(m.T[0], 10) + ", " + strconv.FormatInt(m.T[1], 10) + "|\n|" + strconv.FormatInt(m.B[0], 10) + ", " + strconv.FormatInt(m.B[1], 10) + "|"
    fmt.Println(str)
}

func main() {
    var matrix Matrix2x2
    matrix.new2x2Matrix([4]int64{2, 1, 1, 1})
    
    start := time.Now()
    
    matrix.binaryPow(1e18)
    matrix.Print()
    
    elapsed := time.Since(start)
    fmt.Printf(" %.6f second\n", elapsed.Seconds())
}