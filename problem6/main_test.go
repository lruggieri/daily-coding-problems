package main

import (
	"fmt"
	"github.com/lruggieri/daily-coding-problems/util"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"testing"
)

func TestMemoryAddressToBytesAndBack(t *testing.T) {

	type test struct{
		memoryAddress string
		expectedBytes []bool
	}

	tests := []test{
		{
			memoryAddress:"c000066340",
			expectedBytes:[]bool{false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,true,false,false,true,true,false,false,false,true,true,false,true,false,false,false,false,false,false},
		},
		/*{
			memoryAddress:"0xc000066340",
			expectedBytes:[]bool{true,true,true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},*/
	}

	for tstIdx, tst := range tests{
		res := MemoryAddressToBytes(tst.memoryAddress)
		if len(res) != len(tst.expectedBytes){
			t.Error("test #",tstIdx,"failed, len is different. Expected ",len(tst.expectedBytes),"but got",len(res))
			t.FailNow()
		}
		for bIdx,b := range tst.expectedBytes{
			if b != tst.expectedBytes[bIdx]{
				t.Error("test #",tstIdx,"failed")
				t.FailNow()
			}
		}

		//testing revert
		if BytesToMemoryAddress(res) != tst.memoryAddress{
			t.Error("test #",tstIdx,"error: revert from bytes to memory address string failed")
			t.FailNow()
		}
	}
}

func TestGeneral(t *testing.T){
	debug.SetGCPercent(-1) //disable GC

	util.Logger.Level = logrus.InfoLevel
	head := &XorList{
		Val:0,
	}

	//this way of adding and getting without keeping track of previous stuff if absurdly slow. Do not set 'elements' too high unless you have time to waste
	elements := 100
	for i := 1 ; i < elements ; i++{
		head.Add(i)
	}

	//now our list should hold 'elements' elements and each one should have the value of its index. Let's make sure of it
	for i := 0 ; i < elements ; i++{
		element := head.Get(i)
		if element != nil{
			if element.Val != i{
				t.Error("element #",i,"does not match (value is '",element.Val,"')")
				t.FailNow()
			}
		}else{
			t.Error("element #",i,"is nil")
			t.FailNow()
		}
	}

	//this is a faster way to fetch elements, keeping track of previous/next elements
	currentElement := head
	previousElement := BinaryMemoryAddress(make([]bool,64))
	for i := 0 ; i < elements ; i++{
		if currentElement != nil{
			if currentElement.Val != i{
				t.Error("element #",i,"does not match (value is '",currentElement.Val,"')")
				t.FailNow()
			}
		}else{
			t.Error("element #",i,"is nil")
			t.FailNow()
		}
		tempElement := currentElement
		currentElement = currentElement.next(previousElement)
		previousElement = MemoryAddressToBytes(fmt.Sprintf("%p",tempElement))
	}

}