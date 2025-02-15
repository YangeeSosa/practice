package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Box struct {
	ID       string
	X1, Y1   int
	X2, Y2   int
	Area     int
	Parent   *Box
	Children []*Box
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for testCase := 0; testCase < t; testCase++ {
		scanner.Scan()
		nm := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(nm[0])
		m, _ := strconv.Atoi(nm[1])

		matrix := make([][]rune, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			line := scanner.Text()
			if len(line) < m {
				line += strings.Repeat(" ", m-len(line))
			}
			matrix[i] = []rune(line)
		}

		boxes := findBoxes(matrix)
		findParents(boxes)
		jsonMap := buildJSON(boxes)

		jsonData, _ := json.Marshal(jsonMap)
		fmt.Println(string(jsonData))
	}
}

func findBoxes(matrix [][]rune) []*Box {
	boxes := []*Box{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == '+' && !visited[y][x] {
				j, i := x, y
				j2, i2 := findCorner(matrix, i, j)
				if j2 == -1 || i2 == -1 {
					continue
				}

				if !validateBorders(matrix, i, j, i2, j2) {
					continue
				}

				visited[i][j] = true
				visited[i][j2] = true
				visited[i2][j] = true
				visited[i2][j2] = true

				id := extractID(matrix, i, j)
				width := j2 - j + 1
				height := i2 - i + 1
				area := (width - 2) * (height - 2)

				boxes = append(boxes, &Box{
					ID:   id,
					X1:   j,
					Y1:   i,
					X2:   j2,
					Y2:   i2,
					Area: area,
				})
			}
		}
	}

	return boxes
}

func findCorner(matrix [][]rune, i, j int) (int, int) {
	w := 0
	for x := j + 1; x < len(matrix[i]); x++ {
		if matrix[i][x] == '+' {
			w = x - j
			break
		} else if matrix[i][x] != '-' {
			return -1, -1
		}
	}
	if w == 0 {
		return -1, -1
	}
	j2 := j + w

	h := 0
	for y := i + 1; y < len(matrix); y++ {
		if matrix[y][j] == '+' {
			h = y - i
			break
		} else if matrix[y][j] != '|' {
			return -1, -1
		}
	}
	if h == 0 {
		return -1, -1
	}
	i2 := i + h

	if i2 >= len(matrix) || j2 >= len(matrix[i2]) || matrix[i2][j2] != '+' {
		return -1, -1
	}

	return j2, i2
}

func validateBorders(matrix [][]rune, i, j, i2, j2 int) bool {
	for x := j + 1; x < j2; x++ {
		if matrix[i2][x] != '-' {
			return false
		}
	}

	for y := i + 1; y < i2; y++ {
		if matrix[y][j2] != '|' {
			return false
		}
	}

	for y := i + 1; y < i2; y++ {
		if matrix[y][j] != '|' {
			return false
		}
	}

	return true
}

func extractID(matrix [][]rune, i, j int) string {
	if i+1 >= len(matrix) || j+1 >= len(matrix[i+1]) {
		return ""
	}
	line := matrix[i+1]
	id := []rune{}
	for x := j + 1; x < len(line) && x <= j+3; x++ {
		c := line[x]
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			id = append(id, c)
		} else {
			break
		}
	}
	return string(id)
}

func findParents(boxes []*Box) {
	sort.Slice(boxes, func(a, b int) bool {
		return boxes[a].Area < boxes[b].Area
	})

	for _, child := range boxes {
		var bestParent *Box
		for _, parent := range boxes {
			if parent == child {
				continue
			}
			if isParent(parent, child) {
				if bestParent == nil || parent.Area < bestParent.Area {
					bestParent = parent
				}
			}
		}
		if bestParent != nil {
			child.Parent = bestParent
			bestParent.Children = append(bestParent.Children, child)
		}
	}
}

func isParent(parent, child *Box) bool {
	return child.X1 > parent.X1 && child.X2 < parent.X2 &&
		child.Y1 > parent.Y1 && child.Y2 < parent.Y2
}

func buildJSON(boxes []*Box) map[string]interface{} {
	rootMap := make(map[string]interface{})
	for _, box := range boxes {
		if box.Parent == nil {
			if len(box.Children) == 0 {
				rootMap[box.ID] = box.Area
			} else {
				rootMap[box.ID] = buildBoxJSON(box)
			}
		}
	}
	return rootMap
}

func buildBoxJSON(box *Box) map[string]interface{} {
	m := make(map[string]interface{})
	for _, child := range box.Children {
		if len(child.Children) == 0 {
			m[child.ID] = child.Area
		} else {
			m[child.ID] = buildBoxJSON(child)
		}
	}
	return m
}
