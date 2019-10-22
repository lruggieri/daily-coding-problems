package main

import "testing"

func TestProblem2(t *testing.T) {
	type test struct{
		input []int
		expectedOutput []int
	}

	tests := []test{
		{
			input:[]int{1,2,3,4},
			expectedOutput:[]int{24,12,8,6},
		},
		{
			input:[]int{1,3,2,4},
			expectedOutput:[]int{24,8,12,6},
		},
	}

	for tstIdx,tst := range tests{


		res := Trivial(tst.input)
		checkResult(t,tstIdx,res,tst.expectedOutput,"trivial")

		res = WithDivision(tst.input)
		checkResult(t,tstIdx,res,tst.expectedOutput,"WithDivision")

		res = WithoutDivision(tst.input)
		checkResult(t,tstIdx,res,tst.expectedOutput,"WithoutDivision")

	}

}

func checkResult(t *testing.T, tstIdx int, iResult, iExpectedOutput []int, funcName string){
	if len(iResult) != len(iExpectedOutput){
		t.Error(funcName+"test #",tstIdx,"has a different length than expected")
		t.FailNow()
	}
	for i,v := range iExpectedOutput{
		if v != iResult[i]{
			t.Error(funcName+"test #",tstIdx,", expected",v,"but got",iResult[i])
			t.FailNow()
		}
	}
}