package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Day3() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	var line string

	for filescanner.Scan() {
		line = filescanner.Text()
	}

	sum1 := processInstructions(line, false)
	fmt.Println("Part 1:", sum1)

	sum2 := processInstructions(line, true)
	fmt.Println("Part 2:", sum2)
}

func processInstructions(line string, handleConditionals bool) int {
	var sum int
	enabled := true

	mulIndices := strings.Split(line, "mul")

	for i, v := range mulIndices {
		if i == 0 {
			continue
		}

		if handleConditionals {
			beforeMul := strings.Join(mulIndices[:i], "mul")
			lastDo := strings.LastIndex(beforeMul, "do()")
			lastDont := strings.LastIndex(beforeMul, "don't()")

			if lastDo > lastDont {
				enabled = true
			} else if lastDont > lastDo {
				enabled = false
			}
		}

		var acc string
		var nums []string
		valid := false

		for ind, c := range v {
			if ind == 0 && c != '(' {
				break
			}
			if c == ',' {
				nums = append(nums, acc)
				acc = ""
				continue
			}
			if c == ')' {
				nums = append(nums, acc)
				valid = true
				break
			}
			if !unicode.IsDigit(c) && ind != 0 && c != ',' {
				break
			}
			if unicode.IsDigit(c) {
				acc += string(c)
			}
		}

		if valid && len(nums) == 2 {
			num1, err1 := strconv.Atoi(nums[0])
			num2, err2 := strconv.Atoi(nums[1])
			if err1 == nil && err2 == nil && (!handleConditionals || enabled) {
				sum += num1 * num2
			}
		}
	}
	return sum
}

func main() {
	Day3()
}
