package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data13.txt")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	input, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputSections := strings.Split(input, "\n\n")

	dots := strings.Split(inputSections[0], "\n")

	folds := strings.Split(inputSections[1], "\n")

	for _, fold := range folds {

		fold = strings.Replace(fold, "fold along ", "", -1)

		axis := strings.Split(fold, "=")[0]
		foldAlongValue, _ := strconv.Atoi(strings.Split(fold, "=")[1])

		// first fold at x=655
		// All folds look to be at the half way point

		for index, dot := range dots {
			i, j := getPositionFromString(dot)

			if axis == "x" {
				if i > foldAlongValue {
					i = (foldAlongValue * 2) - i
				} else {
					continue
				}
			} else {
				if j > foldAlongValue {
					j = (foldAlongValue * 2) - j
				} else {
					continue
				}
			}

			dots[index] = buildStringOfPosition(i, j)
		}

		dots = removeDuplicates(dots)
	}

	// final x value is 40
	// final y value is 6
	finalGrid := [40][6]int{}
	for _, dot := range dots {
		i, j := getPositionFromString(dot)
		finalGrid[i][j] = 1
	}

	for j := 0; j < 6; j++ {
		for i := 0; i < 40; i++ {
			if finalGrid[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	//fmt.Println(finalGrid)

}

func removeDuplicates(arr []string) []string {
	occurred := map[string]bool{}
	result := []string{}
	for e := range arr {
		if !occurred[arr[e]] {
			occurred[arr[e]] = true

			// Append to result slice.
			result = append(result, arr[e])
		}
	}

	return result
}

func buildStringOfPosition(iValue int, jValue int) (_ string) {
	return strconv.Itoa(iValue) + "," + strconv.Itoa(jValue)
}

func getPositionFromString(value string) (iValue int, jValue int) {
	values := strings.Split(value, ",")
	i, _ := strconv.Atoi(values[0])
	j, _ := strconv.Atoi(values[1])
	return i, j
}
