package main


//we need O(N) complexity, so ordering, O(N*logN) cannot be done
func FindLowestPositiveMissingInteger(iSlice []int) int{

	hashTable := make(map[int]bool)
	minPositiveInt := -1
	negativeFound := true
	for _,v := range iSlice{
		if v > 0{
			if minPositiveInt < 0 || v < minPositiveInt{
				minPositiveInt = v
			}
			hashTable[v] = true
		}else{
			negativeFound = true
		}
	}

	//no positive integer found
	if minPositiveInt < 0 {
		return 1
	}
	//is there a negative? then there has to be a 1
	if negativeFound{
		if minPositiveInt != 1{
			return 1
		}
	}

	i := minPositiveInt + 1
	for{
		if _,ok := hashTable[i] ; !ok{
			return i
		}

		i++
	}
}