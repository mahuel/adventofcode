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

	for i := 0; i < highestValue; i++ {
		fuelUsed := 0
		for _, crab := range crabs {
			if (crab - i) < 0 {
				fuelUsed += -(crab - i)
			} else {
				fuelUsed += (crab - i)
			}
		}
	}

	fmt.Println(len(fishes))
}
