package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input a list of numbers, separated by comma")
	inputList, err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}

	inputNumbersString := strings.Split(inputList, ",")
	inputNumbers := make([]int,0,len(inputNumbersString))
	for nIdx, n := range inputNumbersString {
		cleanNumber := strings.TrimSpace(n)
		if nInt, err := strconv.Atoi(cleanNumber); err != nil{
			log.Fatal("integer #"+strconv.Itoa(nIdx)+" is not a valid int")
		}else{
			inputNumbers = append(inputNumbers, nInt)
		}
	}

	var k int
	_, err = fmt.Scanf("%d", &k)
	if err != nil{
		log.Fatal(err)
	}

	//we have out input, we can proceed
	fmt.Println("Does your input integers sum to "+strconv.Itoa(k)+"? ",AddToK(inputNumbers,k))
}

//complexity O(n)
func AddToK(iIntegers []int, iK int) bool{
	//reorder the input slice
	sort.Ints(iIntegers)

	leftPointer, rightPointer := 0,len(iIntegers)-1

	for{
		if rightPointer == leftPointer {return false}

		sum := iIntegers[leftPointer] + iIntegers[rightPointer]
		if sum == iK{
			return true
		}else if sum < iK {
			leftPointer++
		}else{
			rightPointer--
		}
	}
}
