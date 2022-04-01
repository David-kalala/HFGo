package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {

	number, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)

	}
	var sum float64
	for _, i := range number {
		sum += i
		fmt.Println(i)
	}
	div := float64(len(number))
	fmt.Printf("Average: %.2f\n", sum/div)
}
