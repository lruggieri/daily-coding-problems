package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionalProgramming(t *testing.T){

	assert.Equal(t,3,car(cons(3,4)))
	assert.Equal(t,4,cdr(cons(3,4)))

}
