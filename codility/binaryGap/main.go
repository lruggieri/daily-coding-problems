package main

import (
	"bytes"
	"strconv"
)

func main(){

}

func Solution(N int) int{
	binaryRep := GetBinaryRepresentation(N, 2)
	maxZero, currentZero := 0,0
	prevOne := false // if the last char was a '1'
	for _,c := range binaryRep{
		switch c {
		case '0':{
			//either we are inside a zero series or we just passed a 1
			if currentZero > 0 || prevOne{
				currentZero++
			}
			prevOne = false
		}
		case '1':{
			if currentZero > maxZero{
				maxZero = currentZero
			}
			prevOne = true
			currentZero = 0
		}
		}
	}
	return maxZero
}

func GetBinaryRepresentation(n int, base int) string{
	result := bytes.NewBufferString("")
	currentInt := n
	for currentInt > 0{
		result.WriteString(strconv.Itoa(currentInt % base))
		currentInt = currentInt / base
	}
	if result.Len() == 0{return "0"}
	runes := []rune(result.String())
	for i,j := 0,len(runes)-1 ; i < j; i,j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}