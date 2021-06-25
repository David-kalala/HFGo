package main

import "fmt"

func main() {
	var sum float64
	tab := [3]float64{71.8, 56.2, 89.5}
	for _, i := range tab {
		sum += i
	}
	diviseur := float64(len(tab))
	fmt.Printf("Average:%.2f \n", sum/diviseur)
}
