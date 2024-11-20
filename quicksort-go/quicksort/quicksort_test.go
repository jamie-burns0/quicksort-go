package quicksort

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"testing"
	"time"
)

func TestSort(t *testing.T) {

	t.Run("Test 1 - with many unsorted items in list", func(t *testing.T) {

		list := []int{2, 1, 5, 4, 3}
		expectedList := []int{1, 2, 3, 4, 5}

		Sort(list)

		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Expected list to be %v. Got %v", expectedList, list)
		}
	})

	t.Run("Test 2 - with one item in list", func(t *testing.T) {

		list := []int{8}
		expectedList := []int{8}

		Sort(list)

		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Expected list to be %v. Got %v", expectedList, list)
		}
	})

	t.Run("Test 3 - with no items in list", func(t *testing.T) {

		list := []int{}
		expectedList := []int{}

		Sort(list)

		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Expected list to be %v. Got %v", expectedList, list)
		}
	})

	t.Run("Test 4 - with sorted list", func(t *testing.T) {

		list := []int{1, 2, 3, 4, 5}
		expectedList := []int{1, 2, 3, 4, 5}

		Sort(list)

		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Expected list to be %v. Got %v", expectedList, list)
		}
	})

	t.Run("Test 5 - with all same item", func(t *testing.T) {

		list := []int{5, 5, 5, 5, 5}
		expectedList := []int{5, 5, 5, 5, 5}

		Sort(list)

		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Expected list to be %v. Got %v", expectedList, list)
		}
	})

	t.Run("Test 6 - with reverse sorted list", func(t *testing.T) {

		list := []int{5, 4, 3, 2, 1}
		expectedList := []int{1, 2, 3, 4, 5}

		Sort(list)

		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Expected list to be %v. Got %v", expectedList, list)
		}
	})
}

func TestPerformance(t *testing.T) {

	t.Run("Test 7 - sort 1,000,000 items between 0-1,000,000", func(t *testing.T) {

		//t.Skip("Skipping performance test")

		const noOfItemsToSort = 1000000

		csv := "\nitems,max_rand_value,Sort_ms,Sort2_ms,AsyncSort_ms,AsyncSortB_ms,AsyncSort2_ms,AsyncSort2B_ms\n"

		for maxSortValue := noOfItemsToSort; maxSortValue > 1; maxSortValue /= 10 {
			csv += _testPerformanceImpl(t, noOfItemsToSort, maxSortValue)
		}

		t.Log(csv)

		t.Fail()
	})
}

func _testPerformanceImpl(t *testing.T, noOfItemsToSort int, maxSortValue int) string {

	csv := fmt.Sprintf("%v,%v",noOfItemsToSort, maxSortValue)

	unsortedData := make([]int, 0, noOfItemsToSort)

	for i := 0; i < noOfItemsToSort; i++ {
		unsortedData = append(unsortedData, rand.IntN(maxSortValue))
	}

	{
		unsortedInts := make([]int, noOfItemsToSort)
		copy(unsortedInts, unsortedData)

		t.Logf("Before Sort with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

		start := time.Now()
		Sort(unsortedInts)
		elapsed := time.Since(start)

		t.Logf("After Sort. Elapsed time is %v\n", elapsed)
		csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	}

	{
		unsortedInts := make([]int, noOfItemsToSort)
		copy(unsortedInts, unsortedData)

		t.Logf("Before Sort2 with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

		start := time.Now()
		Sort2(unsortedInts)
		elapsed := time.Since(start)

		t.Logf("After Sort2. Elapsed time is %v\n", elapsed)
		csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	}

	{
		unsortedInts := make([]int, noOfItemsToSort)
		copy(unsortedInts, unsortedData)

		t.Logf("Before AsyncSort with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

		start := time.Now()
		AsyncSort(unsortedInts)
		elapsed := time.Since(start)

		t.Logf("After AsyncSort. Elapsed time is %v\n", elapsed)
		csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	}

	{
		unsortedInts := make([]int, noOfItemsToSort)
		copy(unsortedInts, unsortedData)

		t.Logf("Before AsyncSortB with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

		start := time.Now()
		AsyncSortB(unsortedInts)
		elapsed := time.Since(start)

		t.Logf("After AsyncSortB. Elapsed time is %v\n", elapsed)
		csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	}

	{
		unsortedInts := make([]int, noOfItemsToSort)
		copy(unsortedInts, unsortedData)

		t.Logf("Before AsyncSort2 with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

		start := time.Now()
		AsyncSort2(unsortedInts)
		elapsed := time.Since(start)

		t.Logf("After AsyncSort2. Elapsed time is %v\n", elapsed)
		csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	}

	{
		unsortedInts := make([]int, noOfItemsToSort)
		copy(unsortedInts, unsortedData)

		t.Logf("Before AsyncSort2B with %v items between 0-%v\n", noOfItemsToSort, maxSortValue)

		start := time.Now()
		AsyncSort2B(unsortedInts)
		elapsed := time.Since(start)

		t.Logf("After Sort. Elapsed time is %v\n", elapsed)
		csv += fmt.Sprintf(",%v", elapsed.Milliseconds())
	}

	csv += fmt.Sprintf("\n")

	return csv
}

func TestSwap(t *testing.T) {

	s := []int{3, 4, 5, 6}

	leftIndex := 0
	rightIndex := 2

	expectedLeftValue := 5
	expectedRightValue := 3

	swap(leftIndex, rightIndex, s)

	if s[leftIndex] != expectedLeftValue {
		t.Errorf("Expected s[%v] to be %v. Got %v", leftIndex, expectedLeftValue, s[leftIndex])
	}

	if s[rightIndex] != expectedRightValue {
		t.Errorf("Expected s[%v] to be %v. Got %v", rightIndex, expectedRightValue, s[rightIndex])
	}
}

func TestPartition(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {

		s := []int{2, 4, 3, 7, 6, 1, 8, 9, 5}
		pivotValue := 6

		//  {2, 4, 3, 7, 6, 1, 8, 9, 5}
		// L             ^             R
		//  {2, 4, 3, 5, 6, 1, 8, 9, 7}
		//            L              R
		//  {2, 4, 3, 5, 1, 6, 8, 9, 7}
		//               L  R
		//  {2, 4, 3, 5, 1, 6, 8, 9, 7}
		//               R  L

		expectedIndex := 4
		expectedSlice := []int{2, 4, 3, 5, 1, 6, 8, 9, 7}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 2", func(t *testing.T) {

		s := []int{9, 7, 8, 9, 6, 1, 4, 3, 5}
		pivotValue := 6

		//  {9, 7, 8, 9, 6, 1, 4, 3, 5}
		// L             ^             R
		//  {5, 7, 8, 9, 6, 1, 4, 3, 9}
		//   L                       R
		//  {5, 3, 8, 9, 6, 1, 4, 7, 9}
		//      L                 R
		//  {5, 3, 4, 9, 6, 1, 8, 7, 9}
		//         L           R
		//  {5, 3, 4, 1, 6, 9, 8, 7, 9}
		//            L     R
		//  {5, 3, 4, 1, 6, 9, 8, 7, 9}
		//               L
		//               R

		expectedIndex := 4
		expectedSlice := []int{5, 3, 4, 1, 6, 9, 8, 7, 9}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 3", func(t *testing.T) {

		s := []int{9, 6, 1}
		pivotValue := 6

		//  {9, 6, 1}
		// L    ^    R
		//  {1, 6, 9}
		//   L     R
		//  {1, 6, 9}
		//      L
		//      R

		expectedIndex := 1
		expectedSlice := []int{1, 6, 9}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 4", func(t *testing.T) {

		s := []int{6}
		pivotValue := 6

		//  {6}
		// L ^ R
		//  {6}
		//   L
		//   R

		expectedIndex := 0
		expectedSlice := []int{6}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 5", func(t *testing.T) {

		s := []int{6, 7}
		pivotValue := 6

		//  {6,7}
		// L ^   R
		//  {6,7}
		//   L
		//   R

		expectedIndex := 0
		expectedSlice := []int{6, 7}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 6", func(t *testing.T) {

		s := []int{7, 6}
		pivotValue := 6

		//  {7,6}
		// L   ^ R
		//  {6,7}
		//   L R
		//  {6,7}
		//   R L

		expectedIndex := 0
		expectedSlice := []int{6, 7}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 7 - everything is the pivot value", func(t *testing.T) {

		s := []int{6, 6, 6, 6, 6}
		pivotValue := 6

		//  {6,6,6,6,6}
		// L ^ ^ ^ ^ ^ R
		//  {6,6,6,6,6}
		//   L       R
		//  {6,6,6,6,6}
		//     L   R
		//  {6,6,6,6,6}
		//       L
		//       R

		expectedIndex := 2
		expectedSlice := []int{6, 6, 6, 6, 6}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 8.1 - everything else is greater than the pivot value", func(t *testing.T) {

		s := []int{6, 7, 8, 9}
		pivotValue := 6

		//  {6,7,8,9}
		// L ^       R
		//  {6,7,8,9}
		//   L
		//   R

		expectedIndex := 0
		expectedSlice := []int{6, 7, 8, 9}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 8.2 - everything else is greater than the pivot value", func(t *testing.T) {

		s := []int{7, 8, 9, 6}
		pivotValue := 6

		//  {7,8,9,6}
		// L       ^ R
		//  {6,8,9,7}
		//   L     R
		//  {6,8,9,7}
		//   R L

		expectedIndex := 0
		expectedSlice := []int{6, 8, 9, 7}

		i := partition(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})
}

func TestPartition2(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {

		s := []int{2, 4, 3, 7, 6, 1, 8, 9, 5}
		pivotValue := 6

		//  {2, 4, 3, 7, 6, 1, 8, 9, 5}
		// L             ^             R
		//  {2, 4, 3, 5, 6, 1, 8, 9, 7}
		//            L              R
		//  {2, 4, 3, 5, 1, 6, 8, 9, 7}
		//               L  R
		//  {2, 4, 3, 5, 1, 6, 8, 9, 7}
		//               R  L

		expectedIndex := 4
		expectedSlice := []int{2, 4, 3, 5, 1, 6, 8, 9, 7}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 2", func(t *testing.T) {

		s := []int{9, 7, 8, 9, 6, 1, 4, 3, 5}
		pivotValue := 6

		//  {9, 7, 8, 9, 6, 1, 4, 3, 5}
		// L             ^             R
		//  {5, 7, 8, 9, 6, 1, 4, 3, 9}
		//   L                       R
		//  {5, 3, 8, 9, 6, 1, 4, 7, 9}
		//      L                 R
		//  {5, 3, 4, 9, 6, 1, 8, 7, 9}
		//         L           R
		//  {5, 3, 4, 1, 6, 9, 8, 7, 9}
		//            L     R
		//  {5, 3, 4, 1, 6, 9, 8, 7, 9}
		//               L
		//               R

		expectedIndex := 4
		expectedSlice := []int{5, 3, 4, 1, 6, 9, 8, 7, 9}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 3", func(t *testing.T) {

		s := []int{9, 6, 1}
		pivotValue := 6

		//  {9, 6, 1}
		// L    ^    R
		//  {1, 6, 9}
		//   L     R
		//  {1, 6, 9}
		//      L
		//      R

		expectedIndex := 1
		expectedSlice := []int{1, 6, 9}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 4", func(t *testing.T) {

		s := []int{6}
		pivotValue := 6

		//  {6}
		// L ^ R
		//  {6}
		//   L
		//   R

		expectedIndex := 0
		expectedSlice := []int{6}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 5", func(t *testing.T) {

		s := []int{6, 7}
		pivotValue := 6

		//  {6,7}
		// L ^   R
		//  {6,7}
		//   L
		//   R

		expectedIndex := 0
		expectedSlice := []int{6, 7}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 6", func(t *testing.T) {

		s := []int{7, 6}
		pivotValue := 6

		//  {7,6}
		// L   ^ R
		//  {6,7}
		//   L R
		//  {6,7}
		//   R L

		expectedIndex := 0
		expectedSlice := []int{6, 7}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 7 - everything is the pivot value", func(t *testing.T) {

		s := []int{6, 6, 6, 6, 6}
		pivotValue := 6

		//  {6,6,6,6,6}
		// L ^ ^ ^ ^ ^ R
		//  {6,6,6,6,6}
		//   L       R
		//  {6,6,6,6,6}
		//     L   R
		//  {6,6,6,6,6}
		//       L
		//       R

		expectedIndex := 2
		expectedSlice := []int{6, 6, 6, 6, 6}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 8.1 - everything else is greater than the pivot value", func(t *testing.T) {

		s := []int{6, 7, 8, 9}
		pivotValue := 6

		//  {6,7,8,9}
		// L ^       R
		//  {6,7,8,9}
		//   L
		//   R

		expectedIndex := 0
		expectedSlice := []int{6, 7, 8, 9}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 8.2 - everything else is greater than the pivot value", func(t *testing.T) {

		s := []int{7, 8, 9, 6}
		pivotValue := 6

		//  {7,8,9,6}
		// L       ^ R
		//  {6,8,9,7}
		//   L     R
		//  {6,8,9,7}
		//   R L

		expectedIndex := 0
		expectedSlice := []int{6, 8, 9, 7}

		i, _ := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 9 - inOrder is false", func(t *testing.T) {

		s := []int{9, 2, 3, 4, 6, 5, 7, 8, 1}
		pivotValue := 6

		//  {9, 2, 3, 4, 6, 5, 7, 8, 1}
		// L             ^             R
		//  {1, 2, 3, 4, 6, 5, 7, 8, 9}
		//   L                       R
		//  {1, 2, 3, 4, 5, 6, 7, 8, 9}
		//               L  R
		//  {1, 2, 3, 4, 5, 6, 7, 8, 9}
		//               R  L

		expectedIndex := 4
		expectedInOrder := false
		expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		i, inOrder := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if inOrder != expectedInOrder {
			t.Errorf("Expected inOrder to be %v. Got %v", expectedInOrder, inOrder)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 9 - inOrder is true", func(t *testing.T) {

		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		pivotValue := 5

		//  {1, 2, 3, 4, 5, 6, 7, 8, 9}
		// L             ^             R
		//  {1, 2, 3, 4, 5, 6, 7, 8, 9}
		//               L
		//               R

		expectedIndex := 4
		expectedInOrder := true
		expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		i, inOrder := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if inOrder != expectedInOrder {
			t.Errorf("Expected inOrder to be %v. Got %v", expectedInOrder, inOrder)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})

	t.Run("Test 9 - inOrder is true", func(t *testing.T) {

		s := []int{6, 6, 6, 6, 6, 6, 6, 6, 6}
		pivotValue := 6

		//  {6, 6, 6, 6, 6, 6, 6, 6, 6}
		// L             ^             R
		//  {6, 6, 6, 6, 6, 6, 6, 6, 6}
		//               L
		//               R

		expectedIndex := 4
		expectedInOrder := true
		expectedSlice := []int{6, 6, 6, 6, 6, 6, 6, 6, 6}

		i, inOrder := partition2(0, len(s)-1, pivotValue, s)

		if i != expectedIndex {
			t.Errorf("Expected index to be %v. Got %v", expectedIndex, i)
		}

		if inOrder != expectedInOrder {
			t.Errorf("Expected inOrder to be %v. Got %v", expectedInOrder, inOrder)
		}

		if !reflect.DeepEqual(s, expectedSlice) {
			t.Errorf("Expected slice to be %v. Got %v", expectedSlice, s)
		}
	})
}
