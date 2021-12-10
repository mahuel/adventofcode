package main

import (
	"fmt"
	"io/ioutil"
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

	numberOfOneFourSevensEights := 0
	for _, line := range data {
		segments := strings.Split(line, " | ")
		uniqueValues := strings.Split(segments[0], " ")
		sort.Slice(uniqueValues, func(i, j int) bool {
			return len(uniqueValues[i]) < len(uniqueValues[j])
		})

		outputValues := strings.Split(segments[1], " ")

		for _, outputValue := range outputValues {
			if len(outputValue) == 2 {
				numberOfOneFourSevensEights++
			}

			if len(outputValue) == 3 {
				numberOfOneFourSevensEights++
			}

			if len(outputValue) == 4 {
				numberOfOneFourSevensEights++
			}

			if len(outputValue) == 7 {
				numberOfOneFourSevensEights++
			}
		}
	}

	fmt.Println(numberOfOneFourSevensEights)
}
