package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	var isAscending bool = true
	for i := 0; i < len(levels)-1; i++ {
		if i == 0 {
			if levels[i] > levels[i+1] {
				isAscending = true
			} else {
				isAscending = false
			}
		} else {
			if isAscending {
				if levels[i] < levels[i+1] {
					return false
				}
			} else {
				if levels[i] > levels[i+1] {
					return false
				}
			}
		}
	}

	for i := 0; i < len(levels)-1; i++ {
		if !isAscending {
			if levels[i+1]-levels[i] > 3 || levels[i+1]-levels[i] <= 0 {
				return false
			}
		} else {
			if levels[i]-levels[i+1] > 3 || levels[i]-levels[i+1] <= 0 {
				return false
			}
		}
	}

	return true
}
func isSafePt2(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		newLevels := make([]int, 0, len(levels)-1)
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)

		if isSafe(newLevels) {
			return true
		}
	}

	return false
}

func Day2(pt2 bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		fmt.Println("Error reading file")
		return
	}

	defer file.Close()

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	var safeReports int

	for filescanner.Scan() {
		line := filescanner.Text()
		parsedLine := strings.Split(line, " ")
		var tmp = make([]int, len(parsedLine))
		for i, v := range parsedLine {
			tmp[i], err = strconv.Atoi(v)
			if err != nil {
				fmt.Print(err)
				fmt.Println("Error converting string to int")
				return
			}

		}
		if pt2 {
			if isSafePt2(tmp) {
				safeReports++
			}
		} else {
			if isSafe(tmp) {
				safeReports++
			}
		}

		// if isSafe(tmp) {
		// 	safeReports++
		// }

		// safeReports += SafeReports(line)
	}
	fmt.Println(safeReports)
}

func main() {
	// Day2(false)
	Day2(true)
}
