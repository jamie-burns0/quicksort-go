package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"jamieburns.me/quicksort/cities"
	"jamieburns.me/quicksort/quicksort"
)

const noOfItemsToSort = 1000000
const maxSortValue = 10000

func main() {

	unsortedInts := make([]int, 0, noOfItemsToSort)

	for i := 0; i < noOfItemsToSort; i++ {
		unsortedInts = append(unsortedInts, rand.IntN(maxSortValue))
	}

	fmt.Printf("Before Sort with 1000000 items between 0-%v\n", maxSortValue)

	start := time.Now()
	quicksort.Sort(unsortedInts)
	elapsed := time.Since(start)

	fmt.Printf("After Sort. Elapsed time is %v\n", elapsed)

	unsortedCities := cities.GenerateRandomCities()

	fmt.Println("\nBefore Sort with 100 global cities in random order")
	fmt.Printf("Unsorted cities\n%v\n", unsortedCities)

	start = time.Now()
	sortedCities := quicksort.Sort(unsortedCities)
	elapsed = time.Since(start)

	fmt.Printf("Sorted cities\n%v\n", sortedCities)
	fmt.Printf("After Sort. Elapsed time is %v\n", elapsed)
}