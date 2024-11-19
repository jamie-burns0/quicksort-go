package quicksort

import (
	//"log"
)

type pivotValueSegment[T ordered] struct {
	segment segment
	pivotValue T
}

type pivotIndexSegment struct {
	segment segment
	pivotIndex int
}

//var maxSegChanLen int

// AsyncSort is an investigation into the value of breaking up Sort to run in
// goroutines and pass data between the goroutines on channels. The implementation
// below works but needs to use buffered channels that need to be about 5% the
// size of our list.
//
// If we use unbuffered channels we will nearly always receive a deadlock error.
// This is caused by the func handling a pivot index when it is available. In two
// places it sends a segment to segChan. When segChan is unbuffered sending the
// second segment will block until the first segment is received by our func that
// is handling a segment when it is available.
//
// If we were only sending one segment to segChan, we might not get the deadlock
// error. However, Hoare's Quicksort works on a divide-and-conquer strategy that
// will divide a segment into two segments, divide each of those segments into
// two segments and so on.
//
// In terms of speed to sort a large list, Sort is still quicker than AsyncSort.
// AsyncSort takes about 3x longer than Sort.
//
// So while AsyncSort has been a good leap into learning about goroutines and
// channels, our synchronous Sort is still by far the fastest way to sort a list
// using Hoare's Quicksort algorithm with a midpoint pivot value strategy.
func AsyncSort[ T ordered](list []T) []T {

	// buffer := int(len(list)/20) + 1
	buffer := int(len(list)) + 1
	segChan := make(chan segment, buffer)
	pivotValueChan := make(chan pivotValueSegment[T], buffer)
	pivotIndexChan := make(chan pivotIndexSegment, buffer)
	doneChan := make(chan segment, buffer)

	go func() { // get our pivotValue when a segment is available
		
		for seg := range segChan {
			//log.Printf("<-segChan [%v]\n", seg)
			
			if seg.rightIndex <= seg.leftIndex { // drop bad segments
				continue
			}

			midIndex := seg.rightIndex - (seg.rightIndex-seg.leftIndex)/2
			pivotValue := list[midIndex]

			sendSegmentToPivotValueChan(seg, pivotValue, pivotValueChan)
		}
	}()

	go func() { // partition our segment when a pivot value is available
		for pvSeg := range pivotValueChan {
			//log.Printf("<-pivotValueChan [%v]\n", pvSeg)

			pivotIndex, inOrder := partition2(pvSeg.segment.leftIndex, pvSeg.segment.rightIndex, pvSeg.pivotValue, list)

			if !inOrder {
				sendSegmentToPivotIndexChan2(pvSeg.segment, pivotIndex, pivotIndexChan)
			} else {
				//log.Printf("Segment %v is in order\n", pvSeg.segment)
				sendSegmentToDoneChan2(pvSeg.segment.leftIndex, pvSeg.segment.rightIndex, doneChan)
			}
		}
	}()

	go func() { // ... when a pivot index is available
		for piSeg := range pivotIndexChan {
			//log.Printf("<-pivotIndexChan [%v]\n", piSeg)

			leftIndex := piSeg.segment.leftIndex
			rightIndex := piSeg.segment.rightIndex
			pivotIndex := piSeg.pivotIndex

			if (rightIndex - leftIndex) <= maxGuaranteedSortedSegmentSize {
				// if we got to here, this segment of our list is guaranteed to be sorted
				// We can add our segment to 
				sendSegmentToDoneChan(leftIndex, rightIndex, doneChan)
				continue
			}
	
			if pivotIndex > leftIndex { // push our left segment onto the channel for partitioning
				sendSegmentToSegChan(leftIndex, pivotIndex, segChan)
			} else if pivotIndex == leftIndex { // our left segment is one item long
				sendSegmentToDoneChan(leftIndex, leftIndex, doneChan)
			}
	
			pivotIndexPlusOne := pivotIndex + 1

			if pivotIndexPlusOne < rightIndex { // push our right segment onto the stack for partitioning
				sendSegmentToSegChan(pivotIndexPlusOne, rightIndex, segChan)
			} else if pivotIndexPlusOne == rightIndex { // our right segment is one item long
				sendSegmentToDoneChan(rightIndex, rightIndex, doneChan)
			}
		}
	}()

	seedSegment := segment{
		leftIndex: 0,
		rightIndex: len(list) - 1,
	}

	//log.Printf("-segChan<- [%v] (seed segment)\n", seedSegment)
	segChan<- seedSegment
	//log.Printf("+segChan<- [%v] (seed segment)\n", seedSegment)

	remaining := len(list)
	//sortedSegmentCount := 0

	for {
		doneSeg := <-doneChan
		//log.Printf("<-doneChan [%v]\n",doneSeg)
		remaining -= (doneSeg.rightIndex - doneSeg.leftIndex + 1)
		//log.Printf("[remaining=%v]\n", remaining)
		//sortedSegmentCount++
		if remaining == 0 {
			//log.Printf("[sortedSegmentCount=%v]\n",sortedSegmentCount)
			break
		}
	}

	//log.Printf("[maxSegChanLen=%v]\n",maxSegChanLen)

	return list
}

func sendSegmentToSegChan(leftIndex int, rightIndex int, segChan chan<- segment) {
	seg := segment{
		leftIndex: leftIndex,
		rightIndex: rightIndex,
	}
	//log.Printf("-segChan<- [%v]\n", seg)
	// segChanLen := len(segChan)
	// if segChanLen > maxSegChanLen {
	// 	maxSegChanLen = segChanLen
	// }
	segChan<- seg
	//log.Printf("+segChan<- [%v]\n", seg)
}

func sendSegmentToPivotValueChan[T ordered](seg segment, pivotValue T, pivotValueChan chan<- pivotValueSegment[T]) {
	pvSeg := pivotValueSegment[T]{
		segment: seg,
		pivotValue: pivotValue,
	}
	//log.Printf("pivotValueChan<- [%v]\n",pvSeg)
	pivotValueChan<- pvSeg
}

func sendSegmentToPivotIndexChan(seg segment, pivotIndex int, pivotIndexChan chan<- pivotIndexSegment) {
	piSeg := pivotIndexSegment{
		segment:    seg,
		pivotIndex: pivotIndex,
	}
	//log.Printf("-pivotIndexChan<- [%v]\n", piSeg)
	pivotIndexChan <- piSeg
	//log.Printf("+pivotIndexChan<- [%v]\n", piSeg)
}

func sendSegmentToDoneChan(leftIndex int, rightIndex int, doneChan chan<- segment) {
	seg := segment{
		leftIndex: leftIndex,
		rightIndex: rightIndex,
	}
	//log.Printf("-doneChan<- [%v]\n", seg)
	doneChan<- seg
	//log.Printf("+doneChan<- [%v]\n", seg)
}
