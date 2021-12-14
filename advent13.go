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

	// first fold at x=655
	// All folds look to be at the half way point

	for index, dot := range dots {
		i, j := getPositionFromString(dot)

		if i > 655 {
			i = 1310 - i
		}

		dots[index] = buildStringOfPosition(i, j)
	}

	dots = removeDuplicates(dots)

	fmt.Println(len(dots))

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
