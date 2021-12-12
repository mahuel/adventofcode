package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data10.txt")
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

	inputLines := strings.Split(input, "\n")

	checkForOpeningSymbolMap := map[string]bool{
		"(": true,
		"[": true,
		"{": true,
		"<": true,
		")": false,
		"]": false,
		"}": false,
		">": false,
	}

	findOpeningSymbolMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	pointsMap := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	scores := make([]int, 0, 1)
	for _, inputLine := range inputLines {
		openingSymbolArray := make([]string, 0, 1)
		for _, inputSymbol := range inputLine {
			if checkForOpeningSymbolMap[string(inputSymbol)] {
				// is an opening symbol
				openingSymbolArray = append(openingSymbolArray, string(inputSymbol))
			} else {
				// is a closing symbol
				openingSymbolShouldBe := findOpeningSymbolMap[string(inputSymbol)]

				index := len(openingSymbolArray) - 1
				if openingSymbolArray[index] == openingSymbolShouldBe {
					openingSymbolArray = append(openingSymbolArray[:index], openingSymbolArray[index+1:]...)
				} else {
					goto nextLine
				}
			}
		}
		if len(openingSymbolArray) > 0 {
			size := len(openingSymbolArray)
			score := int(0)
			for i := size - 1; i >= 0; i-- {
				score = score * 5
				score += int(pointsMap[openingSymbolArray[i]])
			}
			scores = append(scores, score)
		}
	nextLine:
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	fmt.Println(scores[(len(scores)-1)/2])
}
