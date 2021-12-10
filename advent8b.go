package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data8.txt")
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
	finalValue := 0

	for _, line := range data {
		segments := strings.Split(line, " | ")
		uniqueValues := strings.Split(segments[0], " ")
		sort.Slice(uniqueValues, func(i, j int) bool {
			return len(uniqueValues[i]) < len(uniqueValues[j])
		})

		var displaySegmentCorrectionMap = make(map[string]int)

		// a=0...g=6
		numberOfLetterOccurrence := [7]int{}

		for _, uniqueValue := range uniqueValues {
			for _, letter := range uniqueValue {
				if string(letter) == "a" {
					numberOfLetterOccurrence[0]++
				}

				if string(letter) == "b" {
					numberOfLetterOccurrence[1]++
				}

				if string(letter) == "c" {
					numberOfLetterOccurrence[2]++
				}

				if string(letter) == "d" {
					numberOfLetterOccurrence[3]++
				}

				if string(letter) == "e" {
					numberOfLetterOccurrence[4]++
				}

				if string(letter) == "f" {
					numberOfLetterOccurrence[5]++
				}

				if string(letter) == "g" {
					numberOfLetterOccurrence[6]++
				}
			}
		}

		for index, number := range numberOfLetterOccurrence {
			// segment e
			if number == 4 {
				addLetterToMap(index, 10000, displaySegmentCorrectionMap)
			}

			// segment b
			if number == 6 {
				addLetterToMap(index, 10, displaySegmentCorrectionMap)
			}

			// segment f
			if number == 9 {
				addLetterToMap(index, 100000, displaySegmentCorrectionMap)
			}
		}

		for index, uniqueValue := range uniqueValues {
			if index == 0 {
				// segment c
				i := lookForIndexOfUnidentifiedLetter(0, uniqueValue, displaySegmentCorrectionMap)
				displaySegmentCorrectionMap[string(uniqueValue[i])] = 100
				continue
			}

			if index == 1 {
				// segment a
				i := lookForIndexOfUnidentifiedLetter(0, uniqueValue, displaySegmentCorrectionMap)
				displaySegmentCorrectionMap[string(uniqueValue[i])] = 1
				continue
			}

			if index == 2 {
				// segment d
				i := lookForIndexOfUnidentifiedLetter(0, uniqueValue, displaySegmentCorrectionMap)
				displaySegmentCorrectionMap[string(uniqueValue[i])] = 1000
				continue
			}

			if index == 6 {
				// segment g
				i := lookForIndexOfUnidentifiedLetter(0, uniqueValue, displaySegmentCorrectionMap)
				displaySegmentCorrectionMap[string(uniqueValue[i])] = 1000000
				continue
			}
		}

		unknownNumbers := strings.Split(segments[1], " ")
		for index, unknownNumber := range unknownNumbers {
			mapValue := 0
			for _, letter := range unknownNumber {
				mapValue += displaySegmentCorrectionMap[string(letter)]
			}

			place := math.Pow(10, float64(-(index - 3)))
			if mapValue == 1110111 {
				// zero
				finalValue += int(place * 0)
				continue
			}

			if mapValue == 100100 {
				// one
				finalValue += int(place * 1)
				continue
			}

			if mapValue == 1011101 {
				// two
				finalValue += int(place * 2)
				continue
			}

			if mapValue == 1101101 {
				// three
				finalValue += int(place * 3)
				continue
			}

			if mapValue == 101110 {
				// four
				finalValue += int(place * 4)
				continue
			}

			if mapValue == 1101011 {
				// five
				finalValue += int(place * 5)
				continue
			}

			if mapValue == 1111011 {
				// six
				finalValue += int(place * 6)
				continue
			}

			if mapValue == 100101 {
				// seven
				finalValue += int(place * 7)
				continue
			}

			if mapValue == 1111111 {
				// eight
				finalValue += int(place * 8)
				continue
			}

			if mapValue == 1101111 {
				// nine
				finalValue += int(place * 9)
				continue
			}

			fmt.Println("Something went wrong", mapValue)
			return
		}
	}
	fmt.Println(finalValue)
}

func lookForIndexOfUnidentifiedLetter(currentIndex int, word string, displaySegmentCorrectionMap map[string]int) (_ int) {
	if _, ok := displaySegmentCorrectionMap[string(word[currentIndex])]; ok {
		return lookForIndexOfUnidentifiedLetter(currentIndex+1, word, displaySegmentCorrectionMap)
	}
	return currentIndex
}

func addLetterToMap(letterIndex int, value int, displaySegmentCorrectionMap map[string]int) {
	letter := ""
	if letterIndex == 0 {
		letter = "a"
	}

	if letterIndex == 1 {
		letter = "b"
	}

	if letterIndex == 2 {
		letter = "c"
	}

	if letterIndex == 3 {
		letter = "d"
	}

	if letterIndex == 4 {
		letter = "e"
	}

	if letterIndex == 5 {
		letter = "f"
	}

	if letterIndex == 6 {
		letter = "g"
	}

	displaySegmentCorrectionMap[letter] = value
}
