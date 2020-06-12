// In memory inverted index.
// Associates tokens with IDs
package invertedIndex

import (
	"github.com/lruggieri/daily-coding-problems/companies/jobMatching/tokenizer"
	"sync"
)

type InvertedIndex struct {
	sync.RWMutex
	index map[tokenizer.Token]map[int64]bool
}

func New() *InvertedIndex {
	return &InvertedIndex{
		index: make(map[tokenizer.Token]map[int64]bool),
	}
}
func (ii *InvertedIndex) Insert(documentID int64, tokens tokenizer.Tokens) {
	ii.Lock()
	defer ii.Unlock()

	for _, t := range tokens {
		if tokenIDs, ok := ii.index[t]; ok {
			if _, alreadyInserted := tokenIDs[documentID]; !alreadyInserted {
				tokenIDs[documentID] = true
			}
		} else {
			ii.index[t] = map[int64]bool{documentID: true}
		}
	}
}
func (ii *InvertedIndex) GetSimilarDocuments(documentID int64, tokens tokenizer.Tokens, thresholdPercentage float64) []int64 {
	ii.RLock()
	defer ii.RUnlock()

	/*
		For each document with at least one common token, tracks how many common tokens it has
	*/
	similarDocumentTokens := make(map[int64]float64, 0)
	for _, t := range tokens {
		if tokenIDs, ok := ii.index[t]; ok {
			for docID := range tokenIDs {
				if docID != documentID {
					if _, ok := similarDocumentTokens[docID]; ok {
						similarDocumentTokens[docID]++
					} else {
						similarDocumentTokens[docID] = 1
					}
				}
			}
		}
	}
	similarDocuments := make([]int64, 0)
	for docID, similarities := range similarDocumentTokens {
		if similarities/float64(len(tokens))*100 >= thresholdPercentage {
			similarDocuments = append(similarDocuments, docID)
		}
	}
	return similarDocuments
}
