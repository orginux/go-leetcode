package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	minPrice := prices[0]
	var currDiff, maxDiff int

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		currDiff = price - minPrice
		if currDiff > maxDiff {
			maxDiff = currDiff
		}
	}
	return maxDiff
}
