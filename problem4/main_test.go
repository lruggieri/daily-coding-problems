package main

import "testing"

func TestFindLowestPositiveMissingInteger(t *testing.T) {
	type test struct{
		input []int
		expectedOutput int
	}

	tests:= []test{
		{
			input:[]int{3,4,-1,1},
			expectedOutput:2,
		},
		{
			input:[]int{3,4,-1,2,1},
			expectedOutput:5,
		},
		{
			input:[]int{3,4,-1,3},
			expectedOutput:1,
		},
		{//all negative
			input:[]int{-5,-3,-10,-1},
			expectedOutput:1,
		},
		{//negative and positive
			input:[]int{-5,-3,7,9},
			expectedOutput:1,
		},
		{
			input:[]int{1,2,4,3},
			expectedOutput:5,
		},
		{//lots of repetition
			input:[]int{-1,1,1,2,3,4,5,2,1,1,2,0,0},
			expectedOutput:6,
		},
	}

	for tstIdx, tst := range tests{
		res := FindLowestPositiveMissingInteger(tst.input)
		if res != tst.expectedOutput{
			t.Error("test #",tstIdx,"fail (expected",tst.expectedOutput,", got",res,")")
		}
	}
}
