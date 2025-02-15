package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 998244353
const inv2 = 499122177

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		a[i] = num % MOD
	}

	comb := make([][]int, k+1)
	for p := 0; p <= k; p++ {
		comb[p] = make([]int, p+1)
		comb[p][0] = 1
		if p > 0 {
			comb[p][p] = 1
		}
		for m := 1; m < p; m++ {
			comb[p][m] = (comb[p-1][m-1] + comb[p-1][m]) % MOD
		}
	}

	S := make([]int, k+1)
	S[0] = n % MOD
	for m := 1; m <= k; m++ {
		S[m] = 0
	}

	for _, num := range a {
		pow := make([]int, k+1)
		pow[0] = 1
		if k >= 1 {
			pow[1] = num % MOD
		}
		for m := 2; m <= k; m++ {
			pow[m] = (pow[m-1] * pow[1]) % MOD
		}
		for m := 1; m <= k; m++ {
			S[m] = (S[m] + pow[m]) % MOD
		}
	}

	for p := 1; p <= k; p++ {
		sum := 0
		sumSp := S[p]
		for m := 0; m <= p; m++ {
			c := comb[p][m]
			sm := S[m]
			spm := S[p-m]
			s := (sm * spm) % MOD
			s = (s - sumSp + MOD) % MOD
			term := (c * s) % MOD
			sum = (sum + term) % MOD
		}
		res := (sum * inv2) % MOD
		fmt.Println(res)
	}
}
