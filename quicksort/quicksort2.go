package quicksort

import (
	"jamieburns.me/quicksort/stack"
)

func Sort2[T ordered](list []T) []T {

	segmentStack := stack.Stack[segment]{}

	segmentStack.Push(segment{
		leftIndex:  0,
		rightIndex: len(list) - 1,
	})

	for {

		// pop our next segment off the stack
		seg, ok := segmentStack.Pop()

		if !ok {
			break
		}

		leftIndex, rightIndex := seg.leftIndex, seg.rightIndex

		if rightIndex <= leftIndex {
			continue // discard this pair of indexes
		}

		// Calculate the midpoint index and use it to get the pivot value
		midIndex := rightIndex - (rightIndex-leftIndex)/2
		pivotValue := list[midIndex]

		pivotIndex, inOrder := partition2(leftIndex, rightIndex, pivotValue, list)

		if inOrder || (rightIndex-leftIndex) <= maxGuaranteedSortedSegmentSize {
			continue // nothing more to do here - this segment of our list is guaranteed to be sorted
		}

		if pivotIndex > leftIndex { // push our left segment onto the stack for partitioning
			segmentStack.Push(segment{
				leftIndex:  leftIndex,
				rightIndex: pivotIndex,
			})
		}

		pivotIndexPlusOne := pivotIndex + 1

		if pivotIndexPlusOne < rightIndex { // push our right segment onto the stack for partitioning
			segmentStack.Push(segment{
				leftIndex:  pivotIndexPlusOne,
				rightIndex: rightIndex,
			})
		}
	}

	return list
}

// Our leftIndex and rightIndex represent the lower and upper bounds of a segment
// of list such that,
//
//	leftIndex >= 0 and <= rightIndex, AND
//	rightIndex <= len(list)-1 and >= leftIndex
//
// pivotValue will be a value that exists around the middle of the segment
//
// When we have completed our partition operation,
// - all values that are < pivotValue will be on the left of our pivotValue position, AND
// - all values that are > pivotValue will be on the right of our pivotValue position
//
// # We return the index of pivotValue - which might not be near the middle of segment anymore
//
// Is there value in partition returning (int, []int)? Our slice will not grow or shrink as
// Hoare's Quicksort will sort the list in-place. For now, let's rely on the knowledge of
// how quicksort works and only return the pivotValue index
//
// Initially, the nested for-loops were run in goroutines. I thought this would be an easy
// parallelism win. However, in this case, testing showed that utilising goroutines, **increased**
// execution time by around x100
//
// Executing partition() on a goroutine was an improvement on the for-loop goroutines above.
// However, testing showed that utilising goroutines here **increased** execution time
// by around x5.
//
// Goroutines have been removed as the quickest execution on any unsorted list up to
// 1000000 items is with no goroutines
//
// For another experiment in the use of goroutines, see AsyncSort
func partition2[T ordered](leftIndex int, rightIndex int, pivotValue T, list []T) (pivotIndex int, inOrder bool) {

	inOrder = true

	// for our nested for-loops to work in our do..while - in Hoare's Quicksort,
	// after a swap, the left index is incremented and the right index is
	// decremented - we need to move leftIndex left one index and move rightIndex
	// right one index - ie we are starting immediately outside our segment's lower
	// and upper bounds
	leftIndex--
	rightIndex++

	for {

		// going L->R, find the index of the first value that is >= our pivot value
		for leftIndex++; list[leftIndex] < pivotValue; leftIndex++ {
			if leftIndex < len(list) {
				if list[leftIndex] > list[leftIndex+1] {
					inOrder = false
				}
			}
		}

		// going R->L, find the index of the first value that is <= our pivot value
		for rightIndex--; list[rightIndex] > pivotValue; rightIndex-- {
			if rightIndex > 0 {
				if list[rightIndex] < list[rightIndex-1] {
					inOrder = false
				}
			}
		}

		if leftIndex >= rightIndex {
			pivotIndex = rightIndex
			return
		}

		if list[leftIndex] == list[rightIndex] {
			// if we got to here, leftIndex and rightIndex are pointing to our pivot value
			// and there is no value in swapping them
			continue
		}

		// If we got to here, swap the values our left and right indexes are pointing to
		// - our left and right indexes haven't met or crossed AND
		// - they are pointing to values that are positioned on the wrong side of our pivotValue
		swap(leftIndex, rightIndex, list)

		// compare our leftIndex item with the items either side of it for order
		if leftIndex > 0 {
			if list[leftIndex] < list[leftIndex-1] {
				inOrder = false
			}
		}

		if leftIndex < len(list) - 1 {
			if list[leftIndex] > list[leftIndex+1] {
				inOrder = false
			}
		}

		// compare our rightIndex item with the items either side of it for order
		if rightIndex > 0 {
			if list[rightIndex] > list[rightIndex-1] {
				inOrder = false
			}
		}

		if rightIndex < len(list) - 1 {
			if list[rightIndex] < list[rightIndex+1] {
				inOrder = false
			}
		}
	}
}
