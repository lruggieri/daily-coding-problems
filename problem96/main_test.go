package problem96

import (
	"github.com/lruggieri/daily-coding-problems/util"
	"testing"
)

func TestSolution1(t *testing.T) {
	res := Solution([]int{1,2,3})
	util.AssertContainsSliceInt(t,res,[]int{1,2,3})
	util.AssertContainsSliceInt(t,res,[]int{1,3,2})
	util.AssertContainsSliceInt(t,res,[]int{2,1,3})
	util.AssertContainsSliceInt(t,res,[]int{2,3,1})
	util.AssertContainsSliceInt(t,res,[]int{3,1,2})
	util.AssertContainsSliceInt(t,res,[]int{3,2,1})
}

func TestSolution2(t *testing.T) {
	res := Solution([]int{1,2,3,4})
	util.AssertContainsSliceInt(t, res, []int{1, 2, 3, 4})
	util.AssertContainsSliceInt(t, res, []int{1, 2, 4, 3})
	util.AssertContainsSliceInt(t, res, []int{1, 3, 2, 4})
	util.AssertContainsSliceInt(t, res, []int{1, 3, 4, 2})
	util.AssertContainsSliceInt(t, res, []int{1, 4, 3, 2})
	util.AssertContainsSliceInt(t, res, []int{1, 4, 2, 3})
	util.AssertContainsSliceInt(t, res, []int{2, 1, 3, 4})
	util.AssertContainsSliceInt(t, res, []int{2, 1, 4, 3})
	util.AssertContainsSliceInt(t, res, []int{2, 3, 1, 4})
	util.AssertContainsSliceInt(t, res, []int{2, 3, 4, 1})
	util.AssertContainsSliceInt(t, res, []int{2, 4, 1, 3})
	util.AssertContainsSliceInt(t, res, []int{2, 4, 3, 1})
	util.AssertContainsSliceInt(t, res, []int{3, 1, 2, 4})
	util.AssertContainsSliceInt(t, res, []int{3, 1, 4, 2})
	util.AssertContainsSliceInt(t, res, []int{3, 2, 1, 4})
	util.AssertContainsSliceInt(t, res, []int{3, 2, 4, 1})
	util.AssertContainsSliceInt(t, res, []int{3, 4, 1, 2})
	util.AssertContainsSliceInt(t, res, []int{3, 4, 2, 1})
	util.AssertContainsSliceInt(t, res, []int{4, 1, 2, 3})
	util.AssertContainsSliceInt(t, res, []int{4, 1, 3, 2})
	util.AssertContainsSliceInt(t, res, []int{4, 2, 1, 3})
	util.AssertContainsSliceInt(t, res, []int{4, 2, 3, 1})
	util.AssertContainsSliceInt(t, res, []int{4, 3, 1, 2})
	util.AssertContainsSliceInt(t, res, []int{4, 3, 2, 1})
}