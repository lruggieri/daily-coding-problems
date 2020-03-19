package engine

import (
	"github.com/lruggieri/daily-coding-problems/companies/jobMatching/job"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEngine_Insert(t *testing.T) {
	newEngine := New()
	newEngine.Insert(
		job.Job{
			Id:          1,
			Company:     job.Company("C1"),
			Location:    job.Location("L1"),
			Title:       job.Title("T1"),
			Description: job.Description("D1"),
		},
	)
	newEngine.Insert(
		job.Job{
			Id:          2,
			Company:     job.Company("C2"),
			Location:    job.Location("L2"),
			Title:       job.Title("T2"),
			Description: job.Description("D2"),
		},
	)

	assert.Equal(t, true, equalSliceElements([]int64{1,2}, newEngine.GetJobs()))
}

func TestEngine_GetSimilarJobs(t *testing.T) {
	newEngine := New()
	newEngine.Insert(
		job.Job{
			Id:          1,
			Company:     job.Company("C1"),
			Location:    job.Location("L1"),
			Title:       job.Title("T1"),
			Description: job.Description("one two three four five six seven eight nine ten"),
		},
	)
	newEngine.Insert(
		job.Job{
			Id:          2,
			Company:     job.Company("C1"),
			Location:    job.Location("L1"),
			Title:       job.Title("T1"),
			Description: job.Description("ten nine eight seven six five four three"),
		},
	)
	newEngine.Insert(
		job.Job{
			Id:          3,
			Company:     job.Company("C2"), //company changed
			Location:    job.Location("L1"),
			Title:       job.Title("T1"),
			Description: job.Description("ten nine eight seven six five four three"),
		},
	)
	newEngine.Insert(
		job.Job{
			Id:          4,
			Company:     job.Company("C1"),
			Location:    job.Location("L1"),
			Title:       job.Title("T1"),
			Description: job.Description("one two three four five"),
		},
	)

	assert.Equal(t, []int64{2}, newEngine.GetSimilarJobs(1, 80))
	assert.Equal(t, []int64{1}, newEngine.GetSimilarJobs(2,80))
	assert.Equal(t, []int64{}, newEngine.GetSimilarJobs(3,80))
	assert.Equal(t, true,
		equalSliceElements(
			[]int64{2,4},
			newEngine.GetSimilarJobs(1,50),
		),
	)
}

// helps checking if two slices have the same elements (disregarding the position)
func equalSliceElements(slice1, slice2 []int64) bool{
	if len(slice1) != len(slice2){return false}

	mapSlice1 := make(map[int64]bool, len(slice1))
	for _,e1 := range slice1{
		mapSlice1[e1] = true
	}

	for _,e2 := range slice2{
		if _, ok := mapSlice1[e2] ;!ok{
			return false
		}
	}
	return true
}