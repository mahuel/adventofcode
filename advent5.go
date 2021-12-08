package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data5.txt")
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

	vents := strings.Split(input, "\n")

	var seaMap [1000][1000]int

	for _, vent := range vents {
		startEnd := strings.Split(vent, " -> ")
		start := strings.Split(startEnd[0], ",")
		end := strings.Split(startEnd[1], ",")
		if (start[0] == end[0]) || (start[1] == end[1]) {
			x1, _ := strconv.Atoi(start[0])
			x2, _ := strconv.Atoi(end[0])

			y1, _ := strconv.Atoi(start[1])
			y2, _ := strconv.Atoi(end[1])

			ventAlongY := true
			ventAlongX := true

			if x1 == x2 {
				ventAlongX = false
			}

			if y1 == y2 {
				ventAlongY = false
			}

			i := x1
			j := y1
			for {
				seaMap[i][j]++

				if i == x2 && j == y2 {
					break
				} else {
					if ventAlongX {
						if x2 > x1 {
							i++
						} else {
							i--
						}
					}

					if ventAlongY {
						if y2 > y1 {
							j++
						} else {
							j--
						}
					}
				}
			}
		}
	}

	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if seaMap[i][j] > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}
