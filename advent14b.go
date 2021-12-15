package main

import (
	"fmt"
	"io/ioutil"
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

	polymerInsertionMap := map[string]string{}

	currentPolymerPairCount := map[string]int{}
	newPolymerPairCount := map[string]int{}
	emptyPolymerPairCount := map[string]int{}

	// start maps
	for _, line := range polymerInsertions {
		insertion := strings.Split(line, " -> ")
		polymerInsertionMap[insertion[0]] = insertion[1]
		emptyPolymerPairCount[insertion[0]] = 0
		currentPolymerPairCount[insertion[0]] = 0
		newPolymerPairCount[insertion[0]] = 0
	}

	letterCount := map[string]int{}
	for _, letter := range polymerTemplate {
		if _, ok := letterCount[string(letter)]; ok {
			letterCount[string(letter)] += 1
		} else {
			letterCount[string(letter)] = 1
		}
	}

	// count template
	for i := 0; i < len(polymerTemplate)-1; i++ {

		polymerPair := polymerTemplate[i : i+2]
		currentPolymerPairCount[polymerPair] += 1
	}

	// steps
	for i := 0; i < 40; i++ {
		for polymerPair, count := range currentPolymerPairCount {
			newLetter := polymerInsertionMap[polymerPair]

			newPolymerPairCount[string(polymerPair[0])+newLetter] += count
			newPolymerPairCount[newLetter+string(polymerPair[1])] += count

			letterCount[newLetter] += count
		}

		currentPolymerPairCount = newPolymerPairCount
		newPolymerPairCount = copyMap(emptyPolymerPairCount)
	}

	fmt.Println(letterCount)

	highest := 0
	lowest := -1

	for _, count := range letterCount {
		if count > highest {
			highest = count
		}

		if lowest == -1 || lowest > count {
			lowest = count
		}
	}

	fmt.Println(highest - lowest)
}

func copyMap(source map[string]int) (destination map[string]int) {
	destination = map[string]int{}
	for key, value := range source {
		destination[key] = value
	}

	return destination
}
