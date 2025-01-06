package main

import (
	"fmt"
	"sort"
)

func mostActive(customers []string) []string {
	tradeCount := make(map[string]int)
	totalTrades := len(customers)

	for _, customer := range customers {
		tradeCount[customer]++
	}
	activeCustomers := []string{}
	for customer, count := range tradeCount {
		if float64(count)/float64(totalTrades) >= 0.05 {
			activeCustomers = append(activeCustomers, customer)
		}
	}
	sort.Strings(activeCustomers)
	return activeCustomers
}

func main() {
	var n int
	fmt.Scan(&n)
	customers := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&customers[i])
	}
	active := mostActive(customers)
	for _, customer := range active {
		fmt.Println(customer)
	}
}
