package usecases

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

// SearchServiceImpl is an implementation of the SearchService
type SearchServiceImpl struct {
	wordRegexp    *regexp.Regexp
	trigramCounts map[string]float64
	trigramList   map[string]SearchObject
}

// NewSearchServiceImpl returns a new instance of SearchServiceImpl
func NewSearchServiceImpl() *SearchServiceImpl {
	return &SearchServiceImpl{
		wordRegexp:    regexp.MustCompile(`\w+`),
		trigramCounts: make(map[string]float64),
		trigramList:   make(map[string]SearchObject),
	}
}

// Index adds the given object to the search index
func (s *SearchServiceImpl) Index(obj SearchObject) {
	// get trigram counts and add them to trigram counts list
	trigrams := s.getTrigramCounts(obj.Content)
	for trigram := range trigrams {
		s.trigramCounts[trigram] += 1.0
	}

	// add search object to trigram list
	obj.Trigrams = trigrams
	s.trigramList[obj.ID] = obj
}

// Search returns the list of relevant results
func (s *SearchServiceImpl) Search(objtype string, query string) []SearchObject {
	queryTrigrams := s.getTrigramList(query)

	results := make([]SearchObject, 0)
	for _, doc := range s.trigramList {
		score := 0.0
		if doc.Type == objtype {
			for _, queryTrigram := range queryTrigrams {
				if docTrigramFreq, ok := doc.Trigrams[queryTrigram]; ok {
					score += docTrigramFreq * s.getIDF(queryTrigram)
				}
			}

			// add to result if score is above cutoff
			score /= float64(len(queryTrigrams))
			if score > ScoreCutoff {
				fmt.Println(score, doc.Content)
				doc.Score = score
				results = append(results, doc)
			}
		}
	}

	// sort search results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results
}

// getTrigramList extracts a list of trigrams from a string
func (s *SearchServiceImpl) getTrigramList(query string) []string {
	matches := s.wordRegexp.FindAllString(strings.ToLower(query), -1)
	trigrams := make([]string, 0)
	for _, match := range matches {
		word := "  " + match + " "
		for i := 0; i < len(word)-2; i++ {
			trigrams = append(trigrams, word[i:i+2])
		}
	}
	sort.Strings(trigrams)
	return trigrams
}

// getTrigramCounts extracts the trigrams from the query terms into a list
// and gives them normalized weights
func (s *SearchServiceImpl) getTrigramCounts(query string) map[string]float64 {
	// build list of trigrams and their counts
	trigrams := s.getTrigramList(query)
	trigramCounts := make(map[string]float64)
	for _, trigram := range trigrams {
		trigramCounts[trigram] += 1.0
	}

	// normalise trigram count list
	length := float64(len(trigrams))
	for key, val := range trigramCounts {
		trigramCounts[key] = val / length
	}

	return trigramCounts
}

// getIDF gets the inverse of the number of documents with the trigram in it
func (s *SearchServiceImpl) getIDF(trigram string) float64 {
	docCount := 0.0
	for _, doc := range s.trigramList {
		for docTrigram := range doc.Trigrams {
			if trigram == docTrigram {
				docCount += 1.0
			}
		}
	}
	return math.Log(float64(len(s.trigramList)) / (docCount + 1.0))
}
