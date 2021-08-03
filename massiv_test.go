package main

import (
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T) {
	type testCase struct {
		arr      []uint
		resStart uint
		resStop  uint
	}
	var start, stop uint
	testCases := []testCase{
		{
			arr:      []uint{8, 7, 6, 5, 4, 3, 2, 1},
			resStart: 0,
			resStop:  0,
		},
		{
			arr:      []uint{1, 2, 3, 4, 5, 6, 7, 8},
			resStart: 0,
			resStop:  0,
		},
		{
			arr:      []uint{100, 9, 8, 7, 6, 5, 20, 30, 40, 50, 60, 70, 80},
			resStart: 6,
			resStop:  12,
		},
		{
			arr:      []uint{10, 9, 8, 7, 6, 5, 20, 30, 40, 50, 60, 70, 80},
			resStart: 1,
			resStop:  5,
		},
		{
			arr:      []uint{2, 3, 4, 5, 6, 5, 4},
			resStart: 5,
			resStop:  6,
		},

		{
			arr:      []uint{},
			resStart: 0,
			resStop:  0,
		},
		{
			arr:      []uint{2, 3},
			resStart: 0,
			resStop:  0,
		},
		{
			arr:      []uint{4, 31, 2},
			resStart: 1,
			resStop:  1,
		},
		{
			arr:      []uint{2, 1, 4},
			resStart: 1,
			resStop:  1,
		},

		{
			arr:      []uint{2, 3, 2},
			resStart: 1,
			resStop:  1,
		},
	}

	for _, testCaseVal := range testCases {
		start, stop = findUnsortedSubarray(testCaseVal.arr)
		if start != testCaseVal.resStart && stop != testCaseVal.resStop {
			t.Errorf(
				"For the array:%v wanted-(%d,%d), got (%d,%d) ",
				testCaseVal.arr, testCaseVal.resStart, testCaseVal.resStop, start, stop)
		}

	}
}
