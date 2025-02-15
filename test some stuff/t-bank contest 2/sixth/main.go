package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func isCollinear(p1, p2, p3 Point) bool {
	return (p2.x-p1.x)*(p3.y-p1.y) == (p2.y-p1.y)*(p3.x-p1.x)
}

func allCollinear(points []Point) bool {
	if len(points) <= 2 {
		return true
	}
	p0, p1 := points[0], points[1]
	for i := 2; i < len(points); i++ {
		if !isCollinear(p0, p1, points[i]) {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n int
	fmt.Scan(&n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		points[i] = Point{x, y}
	}

	if allCollinear(points) {
		fmt.Println(0)
	} else {
		fmt.Println(n / 3)
	}
}
