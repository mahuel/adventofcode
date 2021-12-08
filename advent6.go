package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data6.txt")
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

	fishes := make([]int, 0, 1)

	for _, fish := range data {
		value, _ := strconv.Atoi(fish)
		fishes = append(fishes, value)
	}

	numberOfDays := 256 //80
	resetValue := 6
	newFishDays := 8

	for i := 0; i < numberOfDays; i++ {
		fishBorn := 0
		for j, fish := range fishes {
			fishes[j] = fish - 1
			if fishes[j] < 0 {
				fishBorn++
				fishes[j] = resetValue
			}
		}
		for j := 0; j < fishBorn; j++ {
			fishes = append(fishes, newFishDays)
		}
	}

	fmt.Println(len(fishes))
}
