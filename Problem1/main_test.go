package main

import (
	"strconv"
	"testing"
)

func TestAddToK(t *testing.T) {
	type test struct{
		integers []int
		k int
		expectedOutput bool
	}

	tests := []test{
		{
			integers:[]int{5,1,9,7,8,6},
			k: 13,
			expectedOutput:true,
		},
		{
			integers:[]int{5,1,9,7,8,6},
			k: 11,
			expectedOutput:true,
		},
		{
			integers:[]int{5,1,9,7,8,599},
			k: 600,
			expectedOutput:true,
		},
		{
			integers:[]int{5,7,8,9,8},
			k: 16,
			expectedOutput:true,
		},
		{
			integers:[]int{1,2,3,4,5,6},
			k: 8,
			expectedOutput:true,
		},
	}

	for tstIdx, tst := range tests{
		res := AddToK(tst.integers, tst.k)
		if res != tst.expectedOutput{
			t.Error("test #"+strconv.Itoa(tstIdx)+" failed")
			t.FailNow()
		}
	}
}
