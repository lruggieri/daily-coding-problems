package main

import "testing"

func TestFunctionalProgramming(t *testing.T){
	type test struct{
		functionToCall func(of operatingFunction) int
		pairMember1 int
		pairMember2 int
		expectedResult int
	}

	tests := []test{
		{
			functionToCall:car,
			pairMember1:3,
			pairMember2:4,
			expectedResult:3,
		},
		{
			functionToCall:cdr,
			pairMember1:3,
			pairMember2:4,
			expectedResult:4,
		},
	}

	for tstIdx,tst := range tests{
		if tst.functionToCall(cons(tst.pairMember1,tst.pairMember2)) != tst.expectedResult{
			t.Error("test #",tstIdx," failed")
		}
	}
}
