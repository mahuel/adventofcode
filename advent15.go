package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readFile() (data string, err error) {
	b, err := ioutil.ReadFile("data15test.txt")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

const gridWidth = 50
const gridHeight = 50

var cavernGrid = []int{}

func main() {
	input, err := readFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Dijkstra's Algorithm

	inputLines := strings.Split(input, "\n")

	for _, inputLine := range inputLines {
		for _, value := range inputLine {
			number, _ := strconv.Atoi(string(value))
			cavernGrid = append(cavernGrid, number)
		}
	}

	nodeGrid := []Node{}

	for j := 0; j < gridHeight; j++ {
		for i := 0; i < gridWidth; i++ {

			node := Node{
				distance: math.MaxInt,
				previous: [2]int{-1, -1},
				position: [2]int{i, j},
			}

			if i == 0 && j == 0 {
				node.distance = 0
			}

			nodeGrid = append(nodeGrid, node)
		}
	}

	calculateDistance(0, 0, nodeGrid)

	positionIValue := gridWidth - 1
	positionJValue := gridHeight - 1
	arrayIndex := getArrayIndex(positionIValue, positionJValue)
	riskValue := 0
	for arrayIndex != 0 {
		riskValue += cavernGrid[arrayIndex]
		previous := nodeGrid[arrayIndex].previous
		positionIValue = previous[0]
		positionJValue = previous[1]
		arrayIndex = getArrayIndex(positionIValue, positionJValue)
	}
	fmt.Println(riskValue)

}

func calculateDistance(positionIValue int, positionJValue int, nodeGrid []Node) {
	distance := nodeGrid[getArrayIndex(positionIValue, positionJValue)].distance
	//down
	if positionJValue < gridHeight-1 {
		downIndex := getArrayIndex(positionIValue, positionJValue+1)
		node := nodeGrid[downIndex]
		tempDistance := cavernGrid[downIndex] + distance
		if tempDistance < node.distance {
			node.distance = tempDistance
			node.previous = [2]int{positionIValue, positionJValue}
			nodeGrid[downIndex] = node
			calculateDistance(positionIValue, positionJValue+1, nodeGrid)
		}
	}

	//right
	if positionIValue < gridWidth-1 {
		rightIndex := getArrayIndex(positionIValue+1, positionJValue)
		node := nodeGrid[rightIndex]
		tempDistance := cavernGrid[rightIndex] + distance
		if tempDistance < node.distance {
			node.distance = tempDistance
			node.previous = [2]int{positionIValue, positionJValue}
			nodeGrid[rightIndex] = node
			calculateDistance(positionIValue+1, positionJValue, nodeGrid)
		}
	}
}

func getArrayIndex(positionIValue int, positionJValue int) (index int) {
	return (positionJValue * gridWidth) + positionIValue
}

type Node struct {
	// position
	position [2]int

	distance int
	previous [2]int
}
