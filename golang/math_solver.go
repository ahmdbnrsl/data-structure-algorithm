package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"strconv"
)

func strToFloat(str string) float64 {
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        return 0.0
    }
    return f
}

func getQuestion() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	fmt.Println("\nRESULT\n")
	var result float64
	for i, line := range lines {
		if i == 2 {
		    arr := strings.Split(line, " ")
		    
		    one, two := strToFloat(arr[2]), strToFloat(arr[4])
		    
		    switch arr[3] {
		    case "÷":
		        result = one / two
		        break;
		    case "×":
		        result = one * two
		        break;
		    case "+":
		        result = one + two
		        break;
		     case "-":
		        result = one - two
		        break;
		    default:
		        result = 0.1
		    }
		}
	}
	
	fmt.Printf("Result : %.0f\n", result)
	str := strconv.FormatFloat(result, 'f', -1, 64)
    fmt.Println("Angka:", str)

    bin := "/data/data/com.termux/files/usr/bin/termux-clipboard-set"
    cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader(str)
    
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Gagal copy ke clipboard:", err)
        fmt.Println("Output:", string(out))
        return
    }

    fmt.Println("✅ Berhasil disalin ke clipboard!")

    getQuestion()
}

func main() {
    fmt.Println("Enter multiple lines of text (press Ctrl+D or Ctrl+Z and Enter to finish):")
	getQuestion()
}