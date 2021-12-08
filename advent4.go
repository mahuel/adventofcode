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

	tables := make([][5][5]string, 0, 1)
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
		tables = append(tables, table)
		editedTables = append(editedTables, table)
	}

	for i, number := range selectedNumbers {
		for k, table := range editedTables {
			for j := 0; j < 25; j++ {
				if table[j/5][j%5] == number {
					table[j/5][j%5] = "b"
				}
			}
			editedTables[k] = table
			if i > 4 {
				if checkForBingo(editedTables[k]) {
					fmt.Println("BINGO!!!")
					fmt.Println("Winning number ", number)
					fmt.Println(tables[k])
					fmt.Println(editedTables[k])
					return
				}
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
