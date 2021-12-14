package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data14.txt")
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

	polymerTemplate := inputSections[0]
	polymerInsertions := strings.Split(inputSections[1], "\n")

	polymerInstertionMap := map[string]string{}

	for _, line := range polymerInsertions {
		insertion := strings.Split(line, " -> ")
		polymerInstertionMap[insertion[0]] = insertion[1]
	}

	// number of steps
	for i := 0; i < 10; i++ {
		insertionList := map[int]string{}

		for j := 0; j < len(polymerTemplate)-1; j++ {
			letterToBeInserted := polymerInstertionMap[polymerTemplate[j:j+2]]
			insertionList[j+1] = letterToBeInserted
		}

		// reverse insertion list
		insertionIndices := []int{}
		for index := range insertionList {
			insertionIndices = append(insertionIndices, index)
		}

		sort.Slice(insertionIndices, func(i, j int) bool {
			return insertionIndices[i] > insertionIndices[j]
		})

		for _, index := range insertionIndices {
			polymerTemplate = polymerTemplate[:index] + insertionList[index] + polymerTemplate[index:]
		}
	}

	letterCount := map[string]int{}
	for _, letter := range polymerTemplate {
		if _, ok := letterCount[string(letter)]; ok {
			letterCount[string(letter)] += 1
		} else {
			letterCount[string(letter)] = 1
		}
	}

	fmt.Println(letterCount)
}
