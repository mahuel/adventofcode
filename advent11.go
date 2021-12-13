package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data11.txt")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

var grid = [10][10]int{}

func main() {
	input, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputLines := strings.Split(input, "\n")

	for i, inputLine := range inputLines {
		for j, value := range inputLine {
			grid[i][j], _ = strconv.Atoi(string(value))
		}
	}

	flashCount := 0
	for k := 0; k < 100; k++ {
		var flashedValues = make(map[string]bool)
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				incrementValue(i, j, flashedValues)
			}
		}
		flashCount += len(flashedValues)

		for position := range flashedValues {
			iValue, jValue := getPositionFromString(position)
			grid[iValue][jValue] = 0
		}
	}

	fmt.Println(flashCount)
}

func addFlash(flashIValue int, flashJValue int, flashed map[string]bool) {
	if _, ok := flashed[buildStringOfPosition(flashIValue, flashJValue)]; ok {
		// value already flashed
		return
	} else {
		flashed[buildStringOfPosition(flashIValue, flashJValue)] = true
	}

	// top left
	if flashJValue > 0 && flashIValue > 0 {
		incrementValue(flashIValue-1, flashJValue-1, flashed)
	}

	// top
	if flashJValue > 0 {
		incrementValue(flashIValue, flashJValue-1, flashed)
	}

	// top right
	if flashIValue < 9 && flashJValue > 0 {
		incrementValue(flashIValue+1, flashJValue-1, flashed)
	}

	// right
	if flashIValue < 9 {
		incrementValue(flashIValue+1, flashJValue, flashed)
	}

	// bottom right
	if flashIValue < 9 && flashJValue < 9 {
		incrementValue(flashIValue+1, flashJValue+1, flashed)
	}

	// bottom
	if flashJValue < 9 {
		incrementValue(flashIValue, flashJValue+1, flashed)
	}

	// bottom left
	if flashIValue > 0 && flashJValue < 9 {
		incrementValue(flashIValue-1, flashJValue+1, flashed)
	}

	// left
	if flashIValue > 0 {
		incrementValue(flashIValue-1, flashJValue, flashed)
	}
}

func incrementValue(iValue int, jValue int, flashed map[string]bool) {
	grid[iValue][jValue] += 1
	if grid[iValue][jValue] > 9 {
		addFlash(iValue, jValue, flashed)
	}
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
