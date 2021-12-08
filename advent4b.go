package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data4.txt")
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

	sections := strings.Split(input, "\n\n")

	editedTables := make([][5][5]string, 0, 1)
	selectedNumbers := make([]string, 0, 1)

	for j, group := range sections {
		if j == 0 {
			selectedNumbers = strings.Split(group, ",")
			continue
		}

		var table [5][5]string
		lines := strings.Split(string(group), "\n")
		for i, line := range lines {
			numbers := strings.Split(line, " ")

			var cleanedNumbers []string
			for _, number := range numbers {
				if number != "" {
					cleanedNumbers = append(cleanedNumbers, number)
				}
			}

			for j, number := range cleanedNumbers {
				table[i][j] = number
			}
		}
		editedTables = append(editedTables, table)
	}

	for i, number := range selectedNumbers {
		winningIndices := make([]int, 0, 1)
		for k, table := range editedTables {
			for j := 0; j < 25; j++ {
				if table[j/5][j%5] == number {
					table[j/5][j%5] = "b"
				}
			}
			editedTables[k] = table
			if i > 4 {
				if checkForBingo(editedTables[k]) {
					winningIndices = append(winningIndices, k)
					if len(editedTables) == 1 {
						fmt.Println("BINGO!!!")
						fmt.Println("Winning number ", number)
						fmt.Println(editedTables[k])
						return
					} else {
						fmt.Println("BINGO!!!", len(editedTables))
					}
				}
			}
		}
		if len(winningIndices) > 0 {
			for l := len(winningIndices) - 1; l >= 0; l-- {
				winningIndex := winningIndices[l]
				editedTables = append(editedTables[:winningIndex], editedTables[winningIndex+1:]...)
			}
		}
	}
}

func checkForBingo(table [5][5]string) (result bool) {
	hasPassedHorizontal := true
	hasPassedVertical := true

	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if table[j][i] != "b" {
				hasPassedHorizontal = false
			}

			if table[i][j] != "b" {
				hasPassedVertical = false
			}

			if i == 4 {
				if hasPassedHorizontal || hasPassedVertical {
					return true
				} else {
					hasPassedHorizontal = true
					hasPassedVertical = true
				}
			}
		}
	}
	return false
}
