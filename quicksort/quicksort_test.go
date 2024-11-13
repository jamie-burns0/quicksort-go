package quicksort

import (
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

	t.Run("Test 7 - sort 1,000,000 items", func(t *testing.T) {

		t.Skip("Skipping large dataset test")

		const size = 1000000

		list := make([]int, 0, size)

		for i := 0; i < size; i++ {
			list = append(list, rand.IntN(10000))
			//list = append( list, i )
		}
		t.Log("Before Sort with 1000000 items between 0-10000.")
		start := time.Now()
		Sort(list)
		elapsed := time.Since(start)
		t.Logf("After Sort. Elapsed time is %v", elapsed)
		t.Fail()
	})
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
