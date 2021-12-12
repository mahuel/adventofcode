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
	batchCreatedFish := make([][2]int, 0, 1)

	for _, fish := range data {
		value, _ := strconv.Atoi(fish)
		fishes = append(fishes, value)
	}

	numberOfDays := 256
	resetValue := 6
	newFishDays := 8

	for i := 0; i < numberOfDays; i++ {
		fmt.Println("Day", i+1)
		fishBorn := 0
		for j, fish := range fishes {
			fishes[j] = fish - 1
			if fishes[j] < 0 {
				fishBorn++
				fishes[j] = resetValue
			}
		}
		for j, batch := range batchCreatedFish {
			batchCreatedFish[j][0] = batch[0] - 1
			if batchCreatedFish[j][0] < 0 {
				fishBorn += batchCreatedFish[j][1]
				batchCreatedFish[j][0] = resetValue
			}
		}

		batchCreatedFish = append(batchCreatedFish, [2]int{newFishDays, fishBorn})
	}

	totalFishes := len(fishes)

	for _, batch := range batchCreatedFish {
		totalFishes += batch[1]
	}

	fmt.Println(totalFishes)
}
