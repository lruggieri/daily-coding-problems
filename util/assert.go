package util

import (
	"reflect"
	"testing"
)

func AssertContainsSliceInt(t *testing.T, output [][]int, element []int) {
	if !containsInt(output, element) {
		t.Fail()
	}
}
func containsInt(output [][]int, element []int) bool {
	for _, array := range output {
		if reflect.DeepEqual(array, element) {
			return true
		}
	}
	return false
}