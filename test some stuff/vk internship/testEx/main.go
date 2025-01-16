package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Item struct {
	point    Point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func findShortestPath(maze [][]int, start, end Point) []Point {
	rows := len(maze)
	cols := len(maze[0])
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}
	dist[start.x][start.y] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{point: start, priority: 0})

	prev := make(map[Point]Point)

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*Item)
		current := currentItem.point

		if current == end {
			return reconstructPath(prev, start, end)
		}

		for _, d := range directions {
			neighbor := Point{current.x + d.x, current.y + d.y}
			if isValid(maze, neighbor, rows, cols) {
				newDist := dist[current.x][current.y] + maze[neighbor.x][neighbor.y]
				if newDist < dist[neighbor.x][neighbor.y] {
					dist[neighbor.x][neighbor.y] = newDist
					prev[neighbor] = current
					heap.Push(pq, &Item{point: neighbor, priority: newDist})
				}
			}
		}
	}
	return nil
}

func isValid(maze [][]int, p Point, rows, cols int) bool {
	return p.x >= 0 && p.x < rows && p.y >= 0 && p.y < cols && maze[p.x][p.y] != 0
}

func reconstructPath(prev map[Point]Point, start, end Point) []Point {
	path := []Point{}
	for at := end; at != start; at = prev[at] {
		path = append([]Point{at}, path...)
	}
	return append([]Point{start}, path...)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dimensions := strings.Split(scanner.Text(), " ")
	rows, _ := strconv.Atoi(dimensions[0])
	cols, _ := strconv.Atoi(dimensions[1])

	maze := make([][]int, rows)
	for i := 0; i < rows; i++ {
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Недостаточно строк для лабиринта")
			os.Exit(1)
		}
		line := strings.Split(scanner.Text(), " ")
		maze[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			maze[i][j], _ = strconv.Atoi(line[j])
		}
	}

	scanner.Scan()
	coords := strings.Split(scanner.Text(), " ")
	if len(coords) != 4 {
		fmt.Fprintln(os.Stderr, "Неверный формат координат")
		os.Exit(1)
	}
	start := Point{x: myAtoi(coords[0]), y: myAtoi(coords[1])}
	end := Point{x: myAtoi(coords[2]), y: myAtoi(coords[3])}

	path := findShortestPath(maze, start, end)
	if path == nil {
		fmt.Fprintln(os.Stderr, "Путь не найден")
		os.Exit(1)
	}

	for _, p := range path {
		fmt.Printf("%d %d\n", p.x, p.y)
	}
	fmt.Println(".")
}

func myAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
