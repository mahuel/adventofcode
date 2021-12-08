package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile() (nums []string, err error) {
	b, err := ioutil.ReadFile("data3.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	return lines, nil
}

func main() {
	lines, err := readFile()
	if err != nil {
		return
	}

	for i, _ := range lines[0] {
		onesAreMajority := checkIfOnesAreMajority(lines, i)
		lines = filterArray(lines, i, onesAreMajority)
		fmt.Println("number of lines:", len(lines))
		if len(lines) < 2 {
			break
		}
	}

	fmt.Println(lines)
}

func checkIfOnesAreMajority(lines []string, index int) (onesAreMajority bool) {
	numberOfOnes := 0

	for _, line := range lines {
		if line[index:index+1] == "1" {
			numberOfOnes++
		}
	}
	return numberOfOnes >= len(lines)/2
}

func checkIfOnesAreMinority(lines []string, index int) (onesAreMinority bool) {
	numberOfOnes := 0

	for _, line := range lines {
		if line[index:index+1] == "1" {
			numberOfOnes++
		}
	}
	return numberOfOnes < len(lines)/2
}

func filterArray(lines []string, index int, onesAreMajority bool) (newArray []string) {
	truncatedLines := make([]string, 0, 1)

	for _, line := range lines {
		if line[index:index+1] == "1" && onesAreMajority {
			fmt.Println("Adding line")
			truncatedLines = append(truncatedLines, line)
		} else {
			if line[index:index+1] != "1" && !onesAreMajority {
				fmt.Println("Adding line")
				truncatedLines = append(truncatedLines, line)
			}
		}
	}

	return truncatedLines
}
