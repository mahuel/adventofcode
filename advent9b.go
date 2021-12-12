package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	basinSizes := make([]int, 0, 1)

	for j, line := range data {
		for i, value := range line {
			tableOfMap[i][j], _ = strconv.Atoi(string(value))
		}
	}

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

			basinSize := startBasinCheck(i, j, tableOfMap)

			basinSizes = append(basinSizes, basinSize)
		}
	}
	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}

func startBasinCheck(currentIValue int, currentJValue int, tableMap [100][100]int) (newBasinCount int) {
	var checkedValues = make(map[string]int)
	checkedValues[buildStringOfPosition(currentIValue, currentJValue)] = 1
	checkAllDirections(currentIValue, currentJValue, tableMap, checkedValues)
	return len(checkedValues)
}

func checkAllDirections(currentIValue int, currentJValue int, tableMap [100][100]int, checkedValues map[string]int) {
	checkTop(currentIValue, currentJValue, tableMap, checkedValues)
	checkBottom(currentIValue, currentJValue, tableMap, checkedValues)
	checkLeft(currentIValue, currentJValue, tableMap, checkedValues)
	checkRight(currentIValue, currentJValue, tableMap, checkedValues)
}

func checkNewPosition(currentIValue int, currentJValue int, tableMap [100][100]int, checkedValues map[string]int) {
	if tableMap[currentIValue][currentJValue] == 9 {
		return
	}

	if _, ok := checkedValues[buildStringOfPosition(currentIValue, currentJValue)]; ok {
		// value already in basin
		return
	} else {
		checkedValues[buildStringOfPosition(currentIValue, currentJValue)] = 1
		checkAllDirections(currentIValue, currentJValue, tableMap, checkedValues)
	}
}

func checkTop(currentIValue int, currentJValue int, tableMap [100][100]int, checkedValues map[string]int) {
	if currentJValue == 0 {
		return
	}

	currentJValue--

	checkNewPosition(currentIValue, currentJValue, tableMap, checkedValues)
}

func checkBottom(currentIValue int, currentJValue int, tableMap [100][100]int, checkedValues map[string]int) {
	if currentJValue == 99 {
		return
	}

	currentJValue++

	checkNewPosition(currentIValue, currentJValue, tableMap, checkedValues)
}

func checkLeft(currentIValue int, currentJValue int, tableMap [100][100]int, checkedValues map[string]int) {
	if currentIValue == 0 {
		return
	}

	currentIValue--

	checkNewPosition(currentIValue, currentJValue, tableMap, checkedValues)
}

func checkRight(currentIValue int, currentJValue int, tableMap [100][100]int, checkedValues map[string]int) {
	if currentIValue == 99 {
		return
	}

	currentIValue++

	checkNewPosition(currentIValue, currentJValue, tableMap, checkedValues)
}

func buildStringOfPosition(iValue int, jValue int) (_ string) {
	return strconv.Itoa(iValue) + "," + strconv.Itoa(jValue)
}
