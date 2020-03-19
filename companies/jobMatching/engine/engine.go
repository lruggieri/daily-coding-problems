package engine

import (
	"github.com/lruggieri/daily-coding-problems/companies/jobMatching/invertedIndex"
	"github.com/lruggieri/daily-coding-problems/companies/jobMatching/job"
	"github.com/lruggieri/daily-coding-problems/companies/jobMatching/tokenizer"
	"sync"
)

type Engine struct{
	descriptionInvertedIndex *invertedIndex.InvertedIndex

	jLock sync.RWMutex
	jobs                    map[int64]*job.Job
}

func New()*Engine{
	return &Engine{
		descriptionInvertedIndex: invertedIndex.New(),
		jobs:                    make(map[int64]*job.Job),
	}
}

// job passed as a copy to avoid the user to be able to change it afterwords
func (e *Engine) Insert(job job.Job){
	e.jLock.Lock()
	e.jobs[job.Id] = &job
	e.jLock.Unlock()

	tokenizerToBeUsed := tokenizer.GetTokenizerSimple()
	e.descriptionInvertedIndex.Insert( job.Id, tokenizerToBeUsed(job.Description.String()) )
}

func (e *Engine) GetJobs() []int64{
	e.jLock.RLock()
	defer e.jLock.RUnlock()
	jobIDs := make([]int64, len(e.jobs))
	count := 0
	for jID := range e.jobs{
		jobIDs[count] = jID
		count ++
	}
	return jobIDs
}

func (e *Engine) GetSimilarJobs(jobID int64, threshold float64) []int64{
	e.jLock.RLock()
	e.jLock.RUnlock()
	inputJob,ok := e.jobs[jobID]

	if !ok{return []int64{}}
	tokenizerToBeUsed := tokenizer.GetTokenizerSimple()
	similarDescriptionJobs := e.descriptionInvertedIndex.GetSimilarDocuments(jobID, tokenizerToBeUsed(inputJob.Description.String()), threshold)
	similarJobs := make([]int64, 0)
	for _, simDescJobID := range similarDescriptionJobs{
		simDescJob,ok := e.jobs[simDescJobID]
		if ok{
			if inputJob.Location.IsEqual(&simDescJob.Location) && inputJob.Title.IsEqual(&simDescJob.Title) && inputJob.Company.IsEqual(&simDescJob.Company){
				similarJobs = append(similarJobs, simDescJobID)
			}
		}
	}
	return similarJobs
}