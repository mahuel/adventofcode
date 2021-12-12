package main

import (
	"fmt"
	"io/ioutil"
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

	findClosingSymbolMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	pointsMap := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	score := 0

	openingSymbolArray := make([]string, 0, 1)
	for _, inputLine := range inputLines {
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
					fmt.Printf("Expected %s but got %s\n", findClosingSymbolMap[openingSymbolArray[index]], string(inputSymbol))
					score += pointsMap[string(inputSymbol)]
					goto nextLine
				}
			}
		}
	nextLine:
	}
	fmt.Println(score)
}
