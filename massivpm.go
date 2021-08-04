package main

import (
	"fmt"
)

func main() {
	var arr []uint
	var start, stop uint

	start, stop = findUnsortedSubarray(arr)
	fmt.Printf("Start=%d, Stop=%d", start, stop)
}

func findUnsortedSubarray(array []uint) (left uint, right uint) {
	type diff struct {
		startInd uint
		stopInd  uint
	}

	var startDiff diff
	var endDiff diff
	var i, j, arrLen uint
	var startFound, stopFound bool

	if len(array) < 3 {
		return 0, 0
	}
	arrLen = uint(len(array))
	startDiff.startInd = 0
	startDiff.stopInd = 0
	endDiff.stopInd = arrLen - 1
	endDiff.startInd = arrLen - 1

	i = 0
	j = arrLen - 1

	if array[0] > array[arrLen-1] {
		for {
			if i > j {
				break
			}
			if !startFound {
				if array[i] > array[i+1] {
					startDiff.stopInd = i + 1
				} else {
					startFound = true
					if stopFound {
						return startDiff.stopInd + 1, endDiff.startInd
					}
				}
				i += 1
			}
			if !stopFound {
				if array[j] < array[j-1] {
					endDiff.startInd = j - 1
				} else {
					stopFound = true
					if startFound {
						return startDiff.stopInd + 1, endDiff.startInd
					}
				}
				j -= 1
			}
		}
	} else if array[0] < array[arrLen-1] {
		for {
			if i > j {
				break
			}
			if !startFound {
				if array[i] < array[i+1] {
					startDiff.stopInd = i + 1
				} else {
					startFound = true
					if stopFound {
						return startDiff.stopInd + 1, endDiff.startInd
					}
				}
				i += 1
			}
			if !stopFound {
				if array[j] > array[j-1] {
					endDiff.startInd = j - 1
				} else {
					stopFound = true
					if startFound {
						return startDiff.stopInd + 1, endDiff.startInd
					}
				}
				j -= 1
			}
		}
	} else {
		for {
			if i > j {
				break
			}

			if !startFound {
				if array[i] == array[i+1] {
					startDiff.stopInd = i + 1
				} else {
					startFound = true
					if stopFound {
						return startDiff.stopInd + 1, endDiff.startInd - 1
					}
				}
				i += 1
			}

			if !stopFound {
				if array[j] == array[j-1] {
					endDiff.startInd = j - 1
				} else {
					stopFound = true
					if startFound {
						return startDiff.stopInd + 1, endDiff.startInd - 1
					}
				}
				j -= 1
			}
		}
	}
	return 0, 0
}
