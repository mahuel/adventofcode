package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data12.txt")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

var completedPaths = []string{}

func main() {
	input, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputLines := strings.Split(input, "\n")

	connectionMap := make(map[string][]string)

	for _, inputLine := range inputLines {
		sections := strings.Split(inputLine, "-")

		addValueToConnectionMap(connectionMap, sections[0], sections[1])
		addValueToConnectionMap(connectionMap, sections[1], sections[0])
	}

	exploreCave(connectionMap, "start", "", false)

	fmt.Println(len(completedPaths))
}

func addValueToConnectionMap(connectionMap map[string][]string, key string, value string) {
	connections := []string{value}
	if _, ok := connectionMap[key]; ok {
		connections = connectionMap[key]
		connections = append(connections, value)
	}

	connectionMap[key] = connections
}

func exploreCave(connectionMap map[string][]string, cave string, currentPath string, hasExploredASmallCaveTwice bool) {
	for _, connectedCave := range connectionMap[cave] {
		if connectedCave == "end" {
			completedPaths = append(completedPaths, currentPath+"->"+cave+"->end")
			continue
		}

		if connectedCave == "start" {
			continue
		}

		if connectedCave == strings.ToUpper(connectedCave) {
			exploreCave(connectionMap, connectedCave, currentPath+"->"+cave, hasExploredASmallCaveTwice)
		} else {
			if strings.Contains(currentPath, connectedCave) {
				if !hasExploredASmallCaveTwice {
					exploreCave(connectionMap, connectedCave, currentPath+"->"+cave, true)
				}
			} else {
				exploreCave(connectionMap, connectedCave, currentPath+"->"+cave, hasExploredASmallCaveTwice)
			}
		}
	}
}
