package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"jamieburns.me/quicksort/cities"
	"jamieburns.me/quicksort/quicksort"
)

const noOfItemsToSort = 1000000
const maxSortValue = 100000

func main() {

	unsortedInts := make([]int, 0, noOfItemsToSort)

	for i := 0; i < noOfItemsToSort; i++ {
		unsortedInts = append(unsortedInts, rand.IntN(maxSortValue))
	}

	fmt.Printf("Before Sort with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

	start := time.Now()
	quicksort.Sort(unsortedInts)
	elapsed := time.Since(start)

	fmt.Printf("After Sort. Elapsed time is %v\n", elapsed)


	intsForAsyncSort := make([]int, 0, noOfItemsToSort)

	for i := 0; i < noOfItemsToSort; i++ {
		intsForAsyncSort = append(intsForAsyncSort, rand.IntN(maxSortValue))
	}

	fmt.Printf("Before AsyncSort with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

	start = time.Now()
	quicksort.AsyncSort(intsForAsyncSort)
	elapsed = time.Since(start)

	fmt.Printf("After AsyncSort. Elapsed time is %v\n", elapsed)


	unsortedCities := cities.GenerateRandomCities()

	fmt.Println("\nBefore Sort with 100 global cities in random order")
	fmt.Printf("Unsorted cities\n%v\n", unsortedCities)

	start = time.Now()
	sortedCities := quicksort.Sort(unsortedCities)
	elapsed = time.Since(start)

	fmt.Printf("Sorted cities\n%v\n", sortedCities)
	fmt.Printf("After Sort. Elapsed time is %v\n", elapsed)
}