package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"time"

	"jamieburns.me/quicksort/cities"
	"jamieburns.me/quicksort/quicksort"
)

const noOfItemsToSort = 1000000
const maxSortValue =    1000000

func main() {

	var memStats runtime.MemStats

	csv := fmt.Sprintf("%v,%v",noOfItemsToSort, maxSortValue)

	unsortedData := make([]int, 0, noOfItemsToSort)

	for i := 0; i < noOfItemsToSort; i++ {
		unsortedData = append(unsortedData, rand.IntN(maxSortValue))
	}


	{
		unsortedInts := make([]int, noOfItemsToSort)
		copy(unsortedInts, unsortedData)

		fmt.Printf("Before Sort with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

		start := time.Now()
		quicksort.Sort(unsortedInts)
		elapsed := time.Since(start)

		fmt.Printf("After Sort. Elapsed time is %v\n", elapsed)
		csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	}

	runtime.GC()
	runtime.ReadMemStats(&memStats)
	fmt.Printf("HeapAlloc: %v, NumGC: %v\n", memStats.HeapAlloc, memStats.NumGC)

	time.Sleep(1*time.Second)

	// {
	// 	unsortedInts := make([]int, noOfItemsToSort)
	// 	copy(unsortedInts, unsortedData)

	// 	fmt.Printf("Before Sort2 with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

	// 	start := time.Now()
	// 	quicksort.Sort2(unsortedInts)
	// 	elapsed := time.Since(start)

	// 	fmt.Printf("After Sort2. Elapsed time is %v\n", elapsed)
	// 	csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	// }

	// runtime.GC()
	// runtime.ReadMemStats(&memStats)
	// fmt.Printf("HeapAlloc: %v, NumGC: %v\n", memStats.HeapAlloc, memStats.NumGC)

	// time.Sleep(1*time.Second)

	// {
	// 	unsortedInts := make([]int, noOfItemsToSort)
	// 	copy(unsortedInts, unsortedData)

	// 	fmt.Printf("Before AsyncSort with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)
	// 	//fmt.Printf("intsForAsyncSort: %v\n", intsForAsyncSort)

	// 	start := time.Now()
	// 	quicksort.AsyncSort(unsortedInts)
	// 	elapsed := time.Since(start)

	// 	fmt.Printf("After AsyncSort. Elapsed time is %v\n", elapsed)
	// 	csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	// }

	// runtime.GC()
	// runtime.ReadMemStats(&memStats)
	// fmt.Printf("HeapAlloc: %v, NumGC: %v\n", memStats.HeapAlloc, memStats.NumGC)

	// time.Sleep(1*time.Second)

	// {
	// 	unsortedInts := make([]int, noOfItemsToSort)
	// 	copy(unsortedInts, unsortedData)

	// 	fmt.Printf("Before AsyncSortB with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

	// 	start := time.Now()
	// 	quicksort.AsyncSortB(unsortedInts)
	// 	elapsed := time.Since(start)

	// 	fmt.Printf("After AsyncSortB. Elapsed time is %v\n", elapsed)
	// 	csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	// }

	// runtime.GC()
	// runtime.ReadMemStats(&memStats)
	// fmt.Printf("HeapAlloc: %v, NumGC: %v\n", memStats.HeapAlloc, memStats.NumGC)

	// time.Sleep(1*time.Second)

	// {
	// 	unsortedInts := make([]int, noOfItemsToSort)
	// 	copy(unsortedInts, unsortedData)

	// 	fmt.Printf("Before AsyncSort2 with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

	// 	start := time.Now()
	// 	quicksort.AsyncSort2(unsortedInts)
	// 	elapsed := time.Since(start)

	// 	fmt.Printf("After AsyncSort2. Elapsed time is %v\n", elapsed)
	// 	csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	// }

	// runtime.GC()
	// runtime.ReadMemStats(&memStats)
	// fmt.Printf("HeapAlloc: %v, NumGC: %v\n", memStats.HeapAlloc, memStats.NumGC)

	// time.Sleep(1*time.Second)

	// {
	// 	unsortedInts := make([]int, noOfItemsToSort)
	// 	copy(unsortedInts, unsortedData)

	// 	fmt.Printf("Before AsyncSort2B with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

	// 	start := time.Now()
	// 	quicksort.AsyncSort2B(unsortedInts)
	// 	elapsed := time.Since(start)

	// 	fmt.Printf("After AsyncSort2B. Elapsed time is %v\n", elapsed)
	// 	csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	// }

	// runtime.GC()
	// runtime.ReadMemStats(&memStats)
	// fmt.Printf("HeapAlloc: %v, NumGC: %v\n", memStats.HeapAlloc, memStats.NumGC)

	// time.Sleep(1*time.Second)

	fmt.Println(csv)

	{
		unsortedCities := cities.GenerateRandomCities()

		fmt.Println("\nBefore Sort with 100 global cities in random order")
		fmt.Printf("Unsorted cities\n%v\n", unsortedCities)

		start := time.Now()
		sortedCities := quicksort.Sort(unsortedCities)
		quicksort.Sort(unsortedCities)
		elapsed := time.Since(start)

		fmt.Printf("Sorted cities\n%v\n", sortedCities)
		fmt.Printf("After Sort. Elapsed time is %v\n", elapsed)
	}
}
