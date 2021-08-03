package main

import (
	"fmt"
)

func main() {
	//var arr []uint
	//arr := []uint{1, 2, 3, 4, 15, 3, 7, 8}
	//arr := []uint{8, 7, 6, 5, 14, 3, 2, 1}
	//arr := []uint{8, 8, 3, 4, 15, 3, 7, 8}
	arr := []uint{4, 31, 2, 1}

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
	var startVal, endVal, i, j, arrLen uint
	var startFound, stopFound bool

	if len(array) < 3 {
		return 0, 0
	}
	arrLen = uint(len(array))
	startVal = array[0]
	endVal = array[arrLen-1]
	startDiff.startInd = 0
	startDiff.stopInd = 0
	endDiff.stopInd = arrLen - 1
	endDiff.startInd = arrLen - 1

	i = 0
	j = arrLen - 1

	if startVal > endVal {
		for {
			if (i > arrLen) || (j == 0) {
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
			}

			i += 1
			j -= 1
		}

	} else if startVal < endVal {
		for {
			if (i > arrLen) || (j == 0) {
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
			}

			i += 1
			j -= 1
		}
	} else {
		for {
			if (i > arrLen) || (j == 0) {
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
			}
			i += 1
			j -= 1
		}
	}
	return 0, 0
}
