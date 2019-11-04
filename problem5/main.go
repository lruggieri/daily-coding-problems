package main


type choice func(a,b int) int
type operatingFunction func(choice) int //our "pair" type

func cons(a,b int) operatingFunction{

	//instance of a pair
	pair := func(c choice) int{
		return c(a,b)
	}

	return pair
}

func car(of operatingFunction) int{
	getRight := func(a,b int)int{
		return a
	}

	return of(getRight)
}

func cdr(of operatingFunction) int{
	getLeft := func(a,b int) int{
		return b
	}

	return of(getLeft)
}