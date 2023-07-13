package main

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/lruggieri/daily-coding-problems/util"
)

const BitSize = 64

type BinaryMemoryAddress []bool //representation of a memory address in bits

type XorList struct {
	Val interface{}
	Xor BinaryMemoryAddress
}

// xl has to be head
func (xl *XorList) Add(iVal interface{}) {
	//first, we have to get to the last list element
	listElement := xl
	prevNode := BinaryMemoryAddress(make([]bool, 64))
	for {
		util.Logger.Debug(fmt.Sprintf("listElement: %v. PrevNode: %v\n", listElement.Val, prevNode))
		if listElement.Xor == nil {
			break
		}
		tempPrev := MemoryAddressToBytes(fmt.Sprintf("%p", listElement))
		oList := listElement.next(prevNode)
		prevNode = tempPrev
		if oList == nil {
			break
		}
		listElement = oList
	}
	newNode := &XorList{Val: iVal}
	newNodeBytes := MemoryAddressToBytes(fmt.Sprintf("%p", newNode))
	util.Logger.Debug(fmt.Sprintf("newNode created.\n\tmemory: %v\n\tbytes: %v\n",
		fmt.Sprintf("%p", newNode), newNodeBytes))
	//changing listElement Xor address based on prevNode and newNode
	listElement.Xor = MakeXorAddress(prevNode, newNodeBytes)
	util.Logger.Debug(fmt.Sprintf("previously last node Xor: %v\n", listElement.Xor))
}

/*
xl has to be head
if iIndex is -1, return last element
*/
func (xl *XorList) Get(iIndex int) (oList *XorList) {
	listElement := xl
	prevNode := BinaryMemoryAddress(make([]bool, 64))
	i := 0
	oList = listElement
	for {
		i++
		if iIndex >= 0 && i > iIndex {
			break
		}
		if listElement.Xor == nil {
			break
		}
		tempPrev := MemoryAddressToBytes(fmt.Sprintf("%p", listElement))
		oList = listElement.next(prevNode)
		prevNode = tempPrev
		if oList == nil {
			break
		}
		listElement = oList
	}
	return
}
func (xl *XorList) next(iPrevAddr BinaryMemoryAddress) (oList *XorList) {
	otherAddress := MakeXorAddress(xl.Xor, iPrevAddr)
	if otherAddress != nil {
		addr, err := strconv.ParseInt("0x"+BytesToMemoryAddress(otherAddress), 0, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		oList = (*XorList)(unsafe.Pointer(uintptr(addr))) //is scary just thinking about it...
	}
	return
}
func (xl *XorList) prev(iNextAddr BinaryMemoryAddress) (oList *XorList) {
	return xl.next(iNextAddr) //the code is exactly the same
}

func MakeXorAddress(iAddr1, iAddr2 BinaryMemoryAddress) (oRes BinaryMemoryAddress) {
	if len(iAddr1) != len(iAddr2) {
		return nil
	}
	oRes = make(BinaryMemoryAddress, len(iAddr1))
	for i := range iAddr1 {
		oRes[i] = (iAddr1[i] && !iAddr2[i]) || (!iAddr1[i] && iAddr2[i])
	}
	return
}

// from a memory address string representation return a BinaryMemoryAddress. 0x is dropped during the process if present.
func MemoryAddressToBytes(iMemoryAddress string) BinaryMemoryAddress {
	if len(iMemoryAddress) > 2 {
		var rawHexAddress string
		if iMemoryAddress[:2] == "0x" {
			rawHexAddress = iMemoryAddress[2:]
		} else {
			rawHexAddress = iMemoryAddress
		}

		uIntRawAddress, err := strconv.ParseUint(rawHexAddress, 16, BitSize)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		bits := make([]bool, 0, BitSize)
		for i := 0; i < BitSize; i++ {
			bits = append([]bool{uIntRawAddress&1 == 1}, bits...)
			uIntRawAddress = uIntRawAddress >> 1
		}
		return bits
	} else {
		return nil
	}
}

// from a BinaryMemoryAddress return the memory address string representation. No 0x added.
func BytesToMemoryAddress(iMemoryBytes BinaryMemoryAddress) (oAddressMemory string) {
	var uIntRawAddress uint64
	for i, b := range iMemoryBytes {
		var bitSetVar uint64
		if b {
			bitSetVar = 1
		}
		uIntRawAddress = uIntRawAddress | (bitSetVar << uint64(len(iMemoryBytes)-1-i))
	}
	return strconv.FormatUint(uIntRawAddress, 16)
}
