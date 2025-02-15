package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums ...int64) int64 {
	res := nums[0]
	for _, n := range nums[1:] {
		res = res / gcd(res, n) * n
	}
	return res
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	parts := strings.Fields(sc.Text())
	n, _ := strconv.Atoi(parts[0])
	x, _ := strconv.ParseInt(parts[1], 10, 64)
	y, _ := strconv.ParseInt(parts[2], 10, 64)
	z, _ := strconv.ParseInt(parts[3], 10, 64)

	sc.Scan()
	a := make([]int64, n)
	for i, s := range strings.Fields(sc.Text()) {
		a[i], _ = strconv.ParseInt(s, 10, 64)
	}

	xy := lcm(x, y)
	xz := lcm(x, z)
	yz := lcm(y, z)
	xyz := lcm(x, y, z)

	mins := [7]int64{
		math.MaxInt64,
		math.MaxInt64,
		math.MaxInt64,
		math.MaxInt64,
		math.MaxInt64,
		math.MaxInt64,
		math.MaxInt64,
	}

	for _, num := range a {
		costs := [7]int64{
			(x - num%x) % x,
			(y - num%y) % y,
			(z - num%z) % z,
			(xy - num%xy) % xy,
			(xz - num%xz) % xz,
			(yz - num%yz) % yz,
			(xyz - num%xyz) % xyz,
		}

		for i := range mins {
			if costs[i] < mins[i] {
				mins[i] = costs[i]
			}
		}
	}

	options := []int64{
		mins[0] + mins[1] + mins[2],
		mins[3] + mins[2],
		mins[4] + mins[1],
		mins[5] + mins[0],
		mins[6],
	}

	min := options[0]
	for _, opt := range options[1:] {
		if opt < min {
			min = opt
		}
	}

	fmt.Println(min)
}
