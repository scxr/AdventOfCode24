package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	var left []int
	var right []int
	for filescanner.Scan() {
		lineNums := filescanner.Text()
		numbers := strings.Split(lineNums, " ")
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Print(err)
			fmt.Println("Error converting string to int")
			return
		}
		num2, err := strconv.Atoi(numbers[3])
		if err != nil {
			fmt.Print(err)
			fmt.Println("Error converting string to int")
			return
		}
		left = append(left, num1)
		right = append(right, num2)
	}

	slices.Sort(left)
	slices.Sort(right)
	var sum int
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}
	}
	fmt.Print(sum)
}

func Day1Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	var left []int
	var right []int
	for filescanner.Scan() {
		lineNums := filescanner.Text()
		numbers := strings.Split(lineNums, " ")
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Print(err)
			fmt.Println("Error converting string to int")
			return
		}
		num2, err := strconv.Atoi(numbers[3])
		if err != nil {
			fmt.Print(err)
			fmt.Println("Error converting string to int")
			return
		}
		left = append(left, num1)
		right = append(right, num2)
	}
	var sum int
	for i := 0; i < len(left); i++ {
		lookingFor := left[i]
		cnt := 0
		for j := 0; j < len(right); j++ {
			if right[j] == lookingFor {
				cnt++
			}
		}
		sum += cnt * lookingFor
	}
	fmt.Print(sum)

}

func main() {
	// Day1()
	Day1Part2()
}
