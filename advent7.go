package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data7.txt")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	input, err := readFile()
	partA := false
	if err != nil {
		return
	}

	data := strings.Split(input, ",")

	crabs := make([]int, 0, 1)
	highestValue := 0
	for _, fish := range data {
		value, _ := strconv.Atoi(fish)
		if value > highestValue {
			highestValue = value
		}
		crabs = append(crabs, value)
	}

	lowestFuelUsed := -1
	for i := 0; i <= highestValue; i++ {
		fuelUsed := 0
		for _, crab := range crabs {
			difference := crab - i
			if difference < 0 {
				difference = -difference
			}

			if partA {
				fuelUsed += difference
			} else {
				fuelUsed += int((float64(difference+1) / float64(2)) * float64(difference))
			}
		}

		if lowestFuelUsed == -1 {
			lowestFuelUsed = fuelUsed
		}

		if fuelUsed < lowestFuelUsed {
			lowestFuelUsed = fuelUsed
		}
	}

	fmt.Println(lowestFuelUsed)
}
