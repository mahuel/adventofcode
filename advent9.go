package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data9.txt")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	input, err := readFile()
	if err != nil {
		return
	}

	data := strings.Split(input, "\n")

	tableOfMap := [100][100]int{}

	for j, line := range data {
		for i, value := range line {
			tableOfMap[i][j], _ = strconv.Atoi(string(value))
		}
	}
	riskLevel := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			value := tableOfMap[i][j]

			// check left
			if i > 0 {
				if value >= tableOfMap[i-1][j] {
					continue
				}
			}

			// check right
			if i < 99 {
				if value >= tableOfMap[i+1][j] {
					continue
				}
			}

			// check top
			if j > 0 {
				if value >= tableOfMap[i][j-1] {
					continue
				}
			}

			// check bottom
			if j < 99 {
				if value >= tableOfMap[i][j+1] {
					continue
				}
			}

			riskLevel += value + 1

			if value == 9 {
				fmt.Println(i, j)
			}
		}
	}

	fmt.Println(riskLevel)

}
